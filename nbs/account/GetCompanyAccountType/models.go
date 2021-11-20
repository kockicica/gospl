package GetCompanyAccountType

import (
	"encoding/xml"
	"fmt"

	"github.com/gosuri/uitable"
)

type Request struct {
	XMLName              xml.Name `xml:"http://communicationoffice.nbs.rs GetCompanyAccountType"`
	CompanyAccountTypeID int      `xml:"companyAccountTypeID"`
}

type Response struct {
	XMLName xml.Name           `xml:"http://communicationoffice.nbs.rs GetCompanyAccountTypeResponse"`
	Result  ResponseTypeResult `xml:"GetCompanyAccountTypeResult"`
}

type ResponseTypeResult struct {
	Content struct {
		Data struct {
			Types CompanyAccountTypes `xml:"CompanyAccountType"`
		} `xml:"CompanyAccountDataSet"`
	} `xml:"diffgram"`
}

type CompanyAccountTypes []CompanyAccountType

type CompanyAccountType struct {
	XMLName              xml.Name `xml:"CompanyAccountType" json:"-"`
	CompanyAccountTypeID int      `xml:"CompanyAccountTypeID" json:"companyAccountTypeID"`
	Name                 string   `xml:"Name" json:"name"`
	Description          string   `xml:"Description" json:"description"`
}

func (t CompanyAccountTypes) WriteOut() {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("TYPE ID", "NAME", "DESCRIPTION")
	for _, r := range t {
		table.AddRow(r.CompanyAccountTypeID, r.Name, r.Description)
	}
	fmt.Println(table)
}
