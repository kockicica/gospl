package GetDebtor

import (
	"encoding/xml"
	"fmt"

	"github.com/gosuri/uitable"
)

type Request struct {
	XMLName                      xml.Name `xml:"http://communicationoffice.nbs.rs GetDebtor"`
	DebtorID                     string   `xml:"debtorID,omitempty"`
	BankCode                     int      `xml:"bankCode,omitempty"`
	AccountNumber                int      `xml:"accountNumber,omitempty"`
	ControlNumber                int      `xml:"controlNumber,omitempty"`
	NationalIdentificationNumber int      `xml:"nationalIdentificationNumber,omitempty"`
	TaxIdentificationNumber      string   `xml:"taxIdentificationNumber"`
	StartItemNumber              int      `xml:"startItemNumber,omitempty"`
	EndItemNumber                int      `xml:"endItemNumber,omitempty"`
}

type Response struct {
	XMLName xml.Name           `xml:"http://communicationoffice.nbs.rs GetDebtorResponse"`
	Result  ResponseTypeResult `xml:"GetDebtorResult"`
}

type ResponseTypeResult struct {
	Content struct {
		Data struct {
			Debtors Debtors `xml:"Debtor"`
		} `xml:"DebtorDataSet"`
	} `xml:"diffgram"`
}

type Debtors []Debtor

type Debtor struct {
	XMLName                      xml.Name `xml:"Debtor" json:"-"`
	DebtorID                     string   `xml:"DebtorID" json:"debtorID,omitempty"`
	Name                         string   `xml:"Name" json:"name,omitempty"`
	Address                      string   `xml:"Address" json:"address,omitempty"`
	City                         string   `xml:"City" json:"city,omitempty"`
	Account                      *int     `xml:"Account" json:"account,omitempty"`
	NationalIdentificationNumber *int     `xml:"NationalIdentificationNumber" json:"nationalIdentificationNumber,omitempty"`
	TaxIdentificationNumber      string   `xml:"TaxIdentificationNumber" json:"taxIdentificationNumber,omitempty"`
	BlockadeTotalAmount          float64  `xml:"BlockadeTotalAmount" json:"blockadeTotalAmount,omitempty"`
	BlockadeCondition            int      `xml:"BlockadeCondition" json:"blockadeCondition,omitempty"`
	BanDate                      string   `xml:"BanDate" json:"banDate,omitempty"`
	BlockadeStatus               *int     `xml:"BlockadeStatus" json:"blockadeStatus,omitempty"`
}

func (d Debtors) WriteOut() {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("ID", "NAME", "CITY", "TAX ID", "ACCOUNT", "BAN DATE", "STATUS", "CONDITION", "TOTAL AMOUNT")
	for _, r := range d {
		table.AddRow(r.DebtorID, r.Name, r.City, r.TaxIdentificationNumber, *r.Account, r.BanDate, *r.BlockadeStatus, r.BlockadeCondition, r.BlockadeTotalAmount)
	}
	fmt.Println(table)

}
