package GetExchangeRateListType

import (
	"encoding/xml"
	"fmt"

	"github.com/gosuri/uitable"
)

type Request struct {
	XMLName                xml.Name `xml:"http://communicationoffice.nbs.rs GetExchangeRateListType"`
	ExchangeRateListTypeID int      `xml:"exchangeRateListTypeID,omitempty"`
}

type Response struct {
	XMLName xml.Name           `xml:"http://communicationoffice.nbs.rs GetExchangeRateListTypeResponse"`
	Result  ResponseTypeResult `xml:"GetExchangeRateListTypeResult"`
}

type ResponseTypeResult struct {
	Schema  string `xml:"schema"`
	Content struct {
		Data struct {
			ListTypes ExchangeRateListTypes `xml:"ExchangeRateListType"`
		} `xml:"ExchangeRateDataSet"`
	} `xml:"diffgram"`
}

type ExchangeRateListTypes []ExchangeRateListType

type ExchangeRateListType struct {
	ID                     string `xml:"id,attr" json:"-"`
	RowOrder               string `xml:"rowOrder,attr" json:"-"`
	HasChanges             string `xml:"hasChanges,attr" json:"-"`
	Name                   string `xml:"Name" json:"name"`
	Description            string `xml:"Description" json:"description"`
	ExchangeRateListTypeID int    `xml:"ExchangeRateListTypeID" json:"exchangeRateListTypeID"`
}

func (e ExchangeRateListTypes) WriteOut() {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("NAME", "DESCRIPTION", "LIST TYPE ID")
	for _, r := range e {
		table.AddRow(r.Name, r.Description, r.ExchangeRateListTypeID)
	}
	fmt.Println(table)

}
