package GetBankStatus

import (
	"encoding/xml"
	"fmt"

	"github.com/gosuri/uitable"
)

type Request struct {
	XMLName      xml.Name `xml:"http://communicationoffice.nbs.rs GetBankStatus"`
	BankStatusID int      `xml:"bankStatusID,omitempty"`
}

type Response struct {
	XMLName xml.Name           `xml:"http://communicationoffice.nbs.rs GetBankStatusResponse"`
	Result  ResponseTypeResult `xml:"GetBankStatusResult"`
}

type ResponseTypeResult struct {
	Content struct {
		Data struct {
			Statuses BankDataSetBankStatuses `xml:"BankStatus"`
		} `xml:"BankDataSet"`
	} `xml:"diffgram"`
}

type BankDataSetBankStatuses []BankDataSetBankStatus

type BankDataSetBankStatus struct {
	XMLName      xml.Name `xml:"BankStatus" json:"-"`
	BankStatusID int      `xml:"BankStatusID" json:"bankStatusID,omitempty"`
	Name         string   `xml:"Name" json:"name,omitempty"`
	Description  string   `xml:"Description" json:"description,omitempty"`
}

func (s BankDataSetBankStatuses) WriteOut() {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("STATUS ID", "NAME", "DESCRIPTION")
	for _, r := range s {
		table.AddRow(r.BankStatusID, r.Name, r.Description)
	}
	fmt.Println(table)

}
