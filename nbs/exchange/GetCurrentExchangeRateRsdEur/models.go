package GetCurrentExchangeRateRsdEur

import (
	"encoding/xml"
	"fmt"

	"github.com/gosuri/uitable"
)

type Request struct {
	XMLName xml.Name `xml:"http://communicationoffice.nbs.rs GetCurrentExchangeRateRsdEur"`
}

type Response struct {
	XMLName xml.Name           `xml:"http://communicationoffice.nbs.rs GetCurrentExchangeRateRsdEurResponse"`
	Result  ResponseTypeResult `xml:"GetCurrentExchangeRateRsdEurResult"`
}

type ResponseTypeResult struct {
	Schema  string `xml:"http://www.w3.org/2001/XMLSchema schema"`
	Content struct {
		Data struct {
			ExchangeRate ExchangeRateRsdEur `xml:"ExchangeRateRsdEur" json:"-"`
		} `xml:"ExchangeRateDataSet"`
	} `xml:"diffgram"`
}

type ExchangeRateRsdEur struct {
	ID           string  `xml:"id,attr" json:"-"`
	RowOrder     string  `xml:"rowOrder,attr" json:"-"`
	Date         string  `xml:"Date" json:"date"`
	TypeID       int     `xml:"TypeID" json:"typeID"`
	Amount       float32 `xml:"Amount" json:"amount"`
	ValidityDate string  `xml:"ValidityDate" json:"validityDate"`
}

func (e ExchangeRateRsdEur) WriteOut() {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("DATE", "TYPE_ID", "AMOUNT", "VALIDITY_DATE")
	table.AddRow(e.Date, e.TypeID, e.Amount, e.ValidityDate)
	fmt.Println(table)

}
