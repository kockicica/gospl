package GetCurrency

import (
	"encoding/xml"
	"fmt"

	"github.com/gosuri/uitable"
)

type Request struct {
	XMLName              xml.Name `xml:"http://communicationoffice.nbs.rs GetCurrency"`
	CurrencyID           string   `xml:"currencyID,omitempty"`
	CurrencyCode         int      `xml:"currencyCode,omitempty"`
	CurrencyCodeNumChar  string   `xml:"currencyCodeNumChar,omitempty"`
	CurrencyCodeAlfaChar string   `xml:"currencyCodeAlfaChar,omitempty"`
}

type Response struct {
	XMLName xml.Name           `xml:"http://communicationoffice.nbs.rs GetCurrencyResponse"`
	Result  ResponseTypeResult `xml:"GetCurrencyResult"`
}

type ResponseTypeResult struct {
	Content struct {
		Data struct {
			Currencies Currencies `xml:"Currency"`
		} `xml:"CurrencyDataSet"`
	} `xml:"diffgram"`
}

type Currencies []Currency

type Currency struct {
	XMLName              xml.Name `xml:"Currency" json:"-"`
	CurrencyID           string   `xml:"CurrencyID" json:"currencyID,omitempty"`
	CurrencyCode         int      `xml:"CurrencyCode" json:"currencyCode,omitempty"`
	CurrencyCodeNumChar  string   `xml:"CurrencyCodeNumChar" json:"currencyCodeNumChar,omitempty"`
	CurrencyCodeAlfaChar string   `xml:"CurrencyCodeAlfaChar" json:"currencyCodeAlfaChar,omitempty"`
	Unit                 int      `xml:"Unit" json:"unit,omitempty"`
	Convertible          *int     `xml:"Convertible" json:"convertible,omitempty"`
	CurrencyNameSerCyrl  string   `xml:"CurrencyNameSerCyrl" json:"currencyNameSerCyrl,omitempty"`
	CurrencyNameSerLat   string   `xml:"CurrencyNameSerLat" json:"currencyNameSerLat,omitempty"`
	CurrencyNameEng      string   `xml:"CurrencyNameEng" json:"currencyNameEng,omitempty"`
	Indicator            *int     `xml:"Indicator" json:"indicator,omitempty"`
	CountryID            string   `xml:"CountryID" json:"countryID,omitempty"`
	CountryCode          *int     `xml:"CountryCode" json:"countryCode,omitempty"`
	CountryCodeNumChar   string   `xml:"CountryCodeNumChar" json:"countryCodeNumChar,omitempty"`
	CountryCodeAlfaChar3 string   `xml:"CountryCodeAlfaChar3" json:"countryCodeAlfaChar3,omitempty"`
	CountryCodeAlfaChar2 string   `xml:"CountryCodeAlfaChar2" json:"countryCodeAlfaChar2,omitempty"`
	CountryNameSerCyrl   string   `xml:"CountryNameSerCyrl" json:"countryNameSerCyrl,omitempty"`
	CountryNameSerLat    string   `xml:"CountryNameSerLat" json:"countryNameSerLat,omitempty"`
	CountryNameEng       string   `xml:"CountryNameEng" json:"countryNameEng,omitempty"`
}

func (c Currencies) WriteOut() {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("NAME", "CURRENCY CODE", "CCN", "CCA", "COUNTRY CODE NUM", "COUNTRY CODE ALPHA", "COUNTRY NAME")
	for _, r := range c {
		table.AddRow(r.CurrencyNameEng, r.CurrencyCode, r.CurrencyCodeNumChar, r.CurrencyCodeAlfaChar, r.CountryCodeNumChar, r.CurrencyCodeAlfaChar, r.CountryNameEng)
	}
	fmt.Println(table)

}
