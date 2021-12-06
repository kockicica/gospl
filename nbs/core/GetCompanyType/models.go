package GetCompanyType

import (
	"encoding/xml"
	"fmt"

	"github.com/gosuri/uitable"
)

type Request struct {
	XMLName       xml.Name `xml:"http://communicationoffice.nbs.rs GetCompanyType"`
	CompanyTypeID int      `xml:"companyTypeID,omitempty"`
}

type Response struct {
	XMLName xml.Name           `xml:"http://communicationoffice.nbs.rs GetCompanyTypeResponse"`
	Result  ResponseTypeResult `xml:"GetCompanyTypeResult"`
}

type ResponseTypeResult struct {
	Content struct {
		Data struct {
			Types Types `xml:"CompanyType"`
		} `xml:"CompanyTypeDataSet"`
	} `xml:"diffgram"`
}

type Types []Type

type Type struct {
	XMLName       xml.Name `xml:"CompanyType" json:"-"`
	CompanyTypeID int      `xml:"CompanyTypeID" json:"companyTypeID"`
	Name          string   `xml:"Name" json:"name"`
	Description   string   `xml:"Description" json:"description"`
}

func (t Types) WriteOut() {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("STATUS ID", "NAME", "DESCRIPTION")
	for _, r := range t {
		table.AddRow(r.CompanyTypeID, r.Name, r.Description)
	}
	fmt.Println(table)

}
