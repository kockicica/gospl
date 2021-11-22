package nbs

import (
	"encoding/xml"

	"github.com/gosuri/uitable"
)

type soapRQ struct {
	XMLName   xml.Name `xml:"soap12:Envelope"`
	XMLNsSoap string   `xml:"xmlns:soap12,attr"`
	XMLNsXSI  string   `xml:"xmlns:xsi,attr"`
	XMLNsXSD  string   `xml:"xmlns:xsd,attr"`
	Headers   soapRQHeader
	Body      soapRQBody
}

type soapRQHeader struct {
	XMLName xml.Name `xml:"soap12:Header"`
	Headers []interface{}
}

type soapRQBody struct {
	XMLName xml.Name `xml:"soap12:Body"`
	Payload interface{}
}

type Envelope struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Header  *Header  `xml:",omitempty"`
	Body    Body
}

type Header struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Content interface{} `xml:",omitempty"`
}

type Body struct {
	//XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	XMLName xml.Name    `xml:"Body"`
	Fault   *Fault      `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type Fault struct {
	//XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	XMLName xml.Name `xml:"Fault"`
	Code    string   `xml:"Code>Value,omitempty"`
	Reason  string   `xml:"Reason>Text,omitempty"`
	Node    string   `xml:"Node,omitempty"`
	// Detail  string   `xml:"detail,omitempty"`
	ErrorInfo ErrorInfo `xml:"detail>ErrorInfo"`
}

type ErrorInfo struct {
	ErrorType    string `xml:"ErrorType"`
	ErrorCode    int    `xml:"ErrorCode"`
	ErrorMessage string `xml:"ErrorMessage"`
}

func (f Fault) Error() string {
	table := uitable.New()
	table.Wrap = true
	table.AddRow("CODE:", f.Code)
	table.AddRow("REASON:", f.Reason)
	table.AddRow("NODE:", f.Node)
	table.AddRow("ERROR TYPE:", f.ErrorInfo.ErrorType)
	table.AddRow("ERROR CODE:", f.ErrorInfo.ErrorCode)
	table.AddRow("ERROR MESSAGE:", f.ErrorInfo.ErrorMessage)
	return table.String()
}

func (h *Header) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var (
		token xml.Token
		err   error
	)
loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}
		if token == nil {
			break
		}
		switch se := token.(type) {
		case xml.StartElement:
			if err := d.DecodeElement(h.Content, &se); err != nil {
				return err
			}
		case xml.EndElement:
			break loop
		}
	}
	return nil
}

func (b *Body) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("content must be a pointer to a struct")
	}
	var (
		token    xml.Token
		err      error
		consumed bool
	)
loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}
		if token == nil {
			break
		}
		envelopeNamespace := "http://schemas.xmlsoap.org/soap/envelope/"
		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside soap body")
			} else if se.Name.Space == envelopeNamespace && se.Name.Local == "Fault" {
				b.Fault = &Fault{}
				b.Content = nil
				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}
				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}
				consumed = true
			}
		case xml.EndElement:
			break loop
		}
	}
	return nil
}
