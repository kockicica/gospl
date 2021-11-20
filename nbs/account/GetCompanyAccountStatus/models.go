package GetCompanyAccountStatus

import (
	"encoding/xml"
	"fmt"

	"github.com/gosuri/uitable"
)

type Request struct {
	XMLName                xml.Name `xml:"http://communicationoffice.nbs.rs GetCompanyAccountStatus"`
	CompanyAccountStatusID int      `xml:"companyAccountStatusID,omitempty"`
}

type Response struct {
	XMLName xml.Name           `xml:"http://communicationoffice.nbs.rs GetCompanyAccountStatusResponse"`
	Result  ResponseTypeResult `xml:"GetCompanyAccountStatusResult"`
}

type ResponseTypeResult struct {
	Content struct {
		Data struct {
			Statuses CompanyAccountStatuses `xml:"CompanyAccountStatus"`
		} `xml:"CompanyAccountDataSet"`
	} `xml:"diffgram"`
}

type CompanyAccountStatuses []CompanyAccountStatus

type CompanyAccountStatus struct {
	//XMLName                xml.Name `xml:"CompanyAccountStatus" json:"-"`
	CompanyAccountStatusID int    `xml:"CompanyAccountStatusID" json:"companyAccountStatusID"`
	Name                   string `xml:"Name" json:"name"`
	Description            string `xml:"Description" json:"description"`
}

func (s CompanyAccountStatuses) WriteOut() {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("STATUS ID", "NAME", "DESCRIPTION")
	for _, r := range s {
		table.AddRow(r.CompanyAccountStatusID, r.Name, r.Description)
	}
	fmt.Println(table)

}
