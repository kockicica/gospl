package GetExchangeRateByCurrency

import (
	"encoding/xml"
	"fmt"

	"github.com/gosuri/uitable"
)

type Request struct {
	XMLName                xml.Name `xml:"http://communicationoffice.nbs.rs GetExchangeRateByCurrency"`
	CurrencyCode           int      `xml:"currencyCode"`
	DateFrom               string   `xml:"dateFrom,omitempty"`
	DateTo                 string   `xml:"dateTo,omitempty"`
	ExchangeRateListTypeID int      `xml:"exchangeRateListTypeID"`
}

type Response struct {
	XMLName xml.Name           `xml:"http://communicationoffice.nbs.rs GetExchangeRateByCurrencyResponse"`
	Result  ResponseTypeResult `xml:"GetExchangeRateByCurrencyResult"`
}

type ResponseTypeResult struct {
	Schema  string `xml:"schema"`
	Content struct {
		Data struct {
			ExchangeRates ExchangeRateDataSetExchangeRates `xml:"ExchangeRate"`
		} `xml:"ExchangeRateDataSet"`
	} `xml:"diffgram"`
}

type ExchangeRateDataSetExchangeRates []ExchangeRateDataSetExchangeRate

type ExchangeRateDataSet struct {
	XMLName xml.Name `xml:"ExchangeRateDataSet" json:"-"`

	ExchangeRate []ExchangeRateDataSetExchangeRate `xml:"ExchangeRate" json:"-"`
}

type ExchangeRateDataSetExchangeRate struct {
	XMLName                xml.Name `xml:"ExchangeRate" json:"-"`
	ExchangeRateListNumber int      `xml:"ExchangeRateListNumber" json:"exchangeRateListNumber"`
	Date                   string   `xml:"Date" json:"date"`
	CreateDate             string   `xml:"CreateDate" json:"createDate"`
	DateTo                 string   `xml:"DateTo" json:"dateTo"`
	ExchangeRateListTypeID int      `xml:"ExchangeRateListTypeID" json:"exchangeRateListTypeID"`
	CurrencyGroupID        *int     `xml:"CurrencyGroupID" json:"currencyGroupID"`
	CurrencyCode           int      `xml:"CurrencyCode" json:"currencyCode"`
	CurrencyCodeNumChar    string   `xml:"CurrencyCodeNumChar" json:"currencyCodeNumChar"`
	CurrencyCodeAlfaChar   string   `xml:"CurrencyCodeAlfaChar" json:"currencyCodeAlfaChar"`
	CurrencyNameSerCyrl    string   `xml:"CurrencyNameSerCyrl" json:"currencyNameSerCyrl"`
	CurrencyNameSerLat     string   `xml:"CurrencyNameSerLat" json:"currencyNameSerLat"`
	CurrencyNameEng        string   `xml:"CurrencyNameEng" json:"currencyNameEng"`
	CountryNameSerCyrl     string   `xml:"CountryNameSerCyrl" json:"countryNameSerCyrl"`
	CountryNameSerLat      string   `xml:"CountryNameSerLat" json:"countryNameSerLat"`
	CountryNameEng         string   `xml:"CountryNameEng" json:"countryNameEng"`
	Unit                   int      `xml:"Unit" json:"unit"`
	BuyingRate             *float64 `xml:"BuyingRate" json:"BuyingRate"`
	MiddleRate             *float64 `xml:"MiddleRate" json:"MiddleRate"`
	SellingRate            *float64 `xml:"SellingRate" json:"sellingRate"`
	FixingRate             *float64 `xml:"FixingRate" json:"fixingRate"`
}

func (rs ExchangeRateDataSetExchangeRates) WriteOut() {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("COUNTRY", "CODE", "CODE ALFA", "DATE", "BUYING RATE", "SELLING RATE", "FIXING RATE", "MIDDLE RATE")
	for _, r := range rs {
		table.AddRow(r.CountryNameEng, r.CurrencyCode, r.CurrencyCodeAlfaChar, r.Date, *r.BuyingRate, *r.SellingRate, *r.FixingRate, *r.MiddleRate)
	}
	fmt.Println(table)

}
