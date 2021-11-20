package GetCurrentExchangeRate

import (
	"encoding/xml"
	"fmt"

	"github.com/gosuri/uitable"
)

type Request struct {
	XMLName                xml.Name `xml:"http://communicationoffice.nbs.rs GetCurrentExchangeRate"`
	ExchangeRateListTypeID int      `xml:"exchangeRateListTypeID"`
}

type Response struct {
	XMLName xml.Name           `xml:"http://communicationoffice.nbs.rs GetCurrentExchangeRateResponse"`
	Result  ResponseTypeResult `xml:"GetCurrentExchangeRateResult"`
}

type ResponseTypeResult struct {
	Content struct {
		Data struct {
			ExchangeRates ExchangeRateDataSetExchangeRateList `xml:"ExchangeRate"`
		} `xml:"ExchangeRateDataSet"`
	} `xml:"diffgram"`
}

type ExchangeRateDataSetExchangeRateList []ExchangeRateDataSetExchangeRate

type ExchangeRateDataSet struct {
	XMLName      xml.Name                          `xml:"ExchangeRateDataSet" json:"-"`
	ExchangeRate []ExchangeRateDataSetExchangeRate `xml:"ExchangeRate" json:"exchangeRate"`
}

type ExchangeRateDataSetExchangeRate struct {
	XMLName                xml.Name `xml:"ExchangeRate" json:"-"`
	ExchangeRateListNumber int      `xml:"ExchangeRateListNumber" json:"exchangeRateListNumber"`
	Date                   string   `xml:"Date" json:"date"`
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
	BuyingRate             *float64 `xml:"BuyingRate" json:"buyingRate"`
	MiddleRate             *float64 `xml:"MiddleRate" json:"middleRate"`
	SellingRate            *float64 `xml:"SellingRate" json:"sellingRate"`
	FixingRate             *float64 `xml:"FixingRate" json:"fixingRate"`
}

func (l ExchangeRateDataSetExchangeRateList) WriteOut() {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("COUNTRY", "CODE", "CODE ALFA", "DATE", "BUYING RATE", "SELLING RATE", "FIXING RATE", "MIDDLE RATE")
	for _, r := range l {
		table.AddRow(r.CountryNameEng, r.CurrencyCode, r.CurrencyCodeAlfaChar, r.Date, *r.BuyingRate, *r.SellingRate, *r.FixingRate, *r.MiddleRate)
	}
	fmt.Println(table)

}
