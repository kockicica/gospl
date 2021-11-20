package GetBankType

import (
	"encoding/xml"
	"fmt"

	"github.com/gosuri/uitable"
)

type Request struct {
	XMLName    xml.Name `xml:"http://communicationoffice.nbs.rs GetBankType"`
	BankTypeID int      `xml:"bankTypeID,omitempty"`
}

type Response struct {
	XMLName xml.Name           `xml:"http://communicationoffice.nbs.rs GetBankTypeResponse"`
	Result  ResponseTypeResult `xml:"GetBankTypeResult"`
}

type ResponseTypeResult struct {
	Content struct {
		Data struct {
			Types BankTypes `xml:"BankType"`
		} `xml:"BankDataSet"`
	} `xml:"diffgram"`
}

type BankTypes []BankType

type BankType struct {
	XMLName     xml.Name `xml:"BankType" json:"-"`
	BankTypeID  int      `xml:"BankTypeID" json:"bankTypeID,omitempty"`
	Name        string   `xml:"Name" json:"name,omitempty"`
	Description string   `xml:"Description" json:"description,omitempty"`
}

func (t BankTypes) WriteOut() {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("BANK TYPE ID", "NAME", "DESCRIPTION")
	for _, r := range t {
		table.AddRow(r.BankTypeID, r.Name, r.Description)
	}
	fmt.Println(table)
}
