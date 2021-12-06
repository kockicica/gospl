package GetEnforcedCollectionDebtorBlockadeStatus

import (
	"encoding/xml"
	"fmt"

	"github.com/gosuri/uitable"
)

type Request struct {
	XMLName                                  xml.Name `xml:"http://communicationoffice.nbs.rs GetEnforcedCollectionDebtorBlockadeStatus"`
	EnforcedCollectionDebtorBlockadeStatusID int      `xml:"enforcedCollectionDebtorBlockadeStatusID,omitempty"`
}

type Response struct {
	XMLName xml.Name           `xml:"http://communicationoffice.nbs.rs GetEnforcedCollectionDebtorBlockadeStatusResponse"`
	Result  ResponseTypeResult `xml:"GetCompanyTypeResult"`
}

type ResponseTypeResult struct {
	Content struct {
		Data struct {
			Statuses Statuses `xml:"EnforcedCollectionDebtorBlockadeStatus"`
		} `xml:"DebtorDataSet"`
	} `xml:"diffgram"`
}

type Statuses []Status

type Status struct {
	XMLName                                  xml.Name `xml:"EnforcedCollectionDebtorBlockadeStatus" json:"-"`
	EnforcedCollectionDebtorBlockadeStatusID int      `xml:"EnforcedCollectionDebtorBlockadeStatusID" json:"enforcedCollectionDebtorBlockadeStatusID"`
	Name                                     string   `xml:"Name" json:"name"`
	Description                              string   `xml:"Description" json:"description"`
}

func (s Statuses) WriteOut() {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("BLOCKADE STATUS ID", "NAME", "DESCRIPTION")
	for _, r := range s {
		table.AddRow(r.EnforcedCollectionDebtorBlockadeStatusID, r.Name, r.Description)
	}
	fmt.Println(table)

}
