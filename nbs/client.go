package nbs

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/pkg/errors"
)

type Writer interface {
	WriteOut()
}

type AuthenticationHeader struct {
	XMLName   xml.Name `xml:"AuthenticationHeader"`
	XMLNs     string   `xml:"xmlns,attr"`
	Username  string   `xml:"UserName"`
	Password  string   `xml:"Password"`
	LicenceId string   `xml:"LicenceID"`
}

type Client struct {
	httpClient http.Client
	Username   string
	Password   string
	LicenceId  string
	baseUrl    string
	ctx        context.Context
	verbose    bool
}

func NewClient(baseUrl, username, password, licenceId string) *Client {
	client := &Client{
		httpClient: http.Client{
			Timeout: time.Second * 30,
		},
		Username:  username,
		Password:  password,
		LicenceId: licenceId,
		baseUrl:   baseUrl,
		ctx:       context.Background(),
	}
	return client
}

func (c *Client) WithContext(ctx context.Context) *Client {
	c.ctx = ctx
	return c
}

func (c *Client) createAuthHeader() *AuthenticationHeader {
	return &AuthenticationHeader{
		XMLNs:     XmlNs,
		Username:  c.Username,
		Password:  c.Password,
		LicenceId: c.LicenceId,
	}
}

func (c *Client) soapCall(serviceUrl, action string, data interface{}) ([]byte, error) {
	v := soapRQ{
		XMLNsSoap: "http://www.w3.org/2003/05/soap-envelope",
		XMLNsXSI:  "http://www.w3.org/2001/XMLSchema-instance",
		XMLNsXSD:  "http://www.w3.org/2001/XMLSchema",
		Headers:   soapRQHeader{Headers: []interface{}{c.createAuthHeader()}},
		Body: soapRQBody{
			Payload: data,
		},
	}

	payload, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, err
	}

	wsUrl := fmt.Sprintf("%s/%s", c.baseUrl, serviceUrl)

	req, err := http.NewRequest("POST", wsUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(c.ctx)

	req.Header.Set("Accept", "text/xml, multipart/related")
	req.Header.Set("SOAPAction", action)
	req.Header.Set("Content-Type", "application/soap+xml; charset=utf8")

	if c.verbose {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Request:\n%s\n", dump)
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		soapFault, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read SOAP response body")
		}
		msg := fmt.Sprintf("HTTP status code: %d, SOAP fault: \n%#v", response.StatusCode, string(soapFault))
		return nil, errors.New(msg)
	}
	defer response.Body.Close()

	if c.verbose {

		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Response:\n%s\n", dump)
	}

	bodyBts, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return bodyBts, nil

}

func (c *Client) handle(serviceUrl, action string, request, response interface{}) error {

	data, err := c.soapCall(serviceUrl, action, request)
	if err != nil {
		return err
	}

	rspEnvelope := Envelope{}
	rspEnvelope.Body = Body{Content: &response}
	if err = xml.Unmarshal(data, &rspEnvelope); err != nil {
		return err
	}

	return nil

}

func (c *Client) WriteJson(data interface{}, name string) error {

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(name, jsonData, 0666); err != nil {
		return err
	}

	return nil
}

func (c *Client) SetVerbose(verbose bool) {
	c.verbose = verbose
}
