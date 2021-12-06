package GetCompanyStatus

import (
	"encoding/xml"
	"fmt"

	"github.com/gosuri/uitable"
)

type Request struct {
	XMLName         xml.Name `xml:"http://communicationoffice.nbs.rs GetCompanyStatus"`
	CompanyStatusID int      `xml:"companyStatusID,omitempty"`
}

type Response struct {
	XMLName xml.Name           `xml:"http://communicationoffice.nbs.rs GetCompanyStatusResponse"`
	Result  ResponseTypeResult `xml:"GetCompanyStatusResult"`
}

type ResponseTypeResult struct {
	Content struct {
		Data struct {
			Statuses Statuses `xml:"CompanyStatus"`
		} `xml:"CompanyStatusDataSet"`
	} `xml:"diffgram"`
}

type Statuses []CompanyStatus

type CompanyStatus struct {
	XMLName         xml.Name `xml:"CompanyStatus" json:"-"`
	CompanyStatusID int      `xml:"CompanyStatusID"`
	Name            string   `xml:"Name"`
	Description     string   `xml:"Description"`
}

func (s Statuses) WriteOut() {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("STATUS ID", "NAME", "DESCRIPTION")
	for _, r := range s {
		table.AddRow(r.CompanyStatusID, r.Name, r.Description)
	}
	fmt.Println(table)

}
