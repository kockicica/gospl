package GetCompanyAccount

import (
	"encoding/xml"
	"fmt"

	"github.com/gosuri/uitable"
)

type Request struct {
	XMLName                      xml.Name `xml:"http://communicationoffice.nbs.rs GetCompanyAccount"`
	NationalIdentificationNumber int      `xml:"nationalIdentificationNumber,omitempty"`
	TaxIdentificationNumber      string   `xml:"taxIdentificationNumber,omitempty"`
	BankCode                     int      `xml:"bankCode,omitempty"`
	AccountNumber                int      `xml:"accountNumber,omitempty"`
	ControlNumber                int      `xml:"controlNumber,omitempty"`
	CompanyName                  string   `xml:"companyName,omitempty"`
	City                         string   `xml:"city,omitempty"`
	StartItemNumber              int      `xml:"startItemNumber,omitempty"`
	EndItemNumber                int      `xml:"endItemNumber,omitempty"`
}

type Response struct {
	XMLName xml.Name           `xml:"http://communicationoffice.nbs.rs GetCompanyAccountResponse"`
	Result  ResponseTypeResult `xml:"GetCompanyAccountResult"`
}

type ResponseTypeResult struct {
	Content struct {
		Data struct {
			Accounts CompanyAccounts `xml:"CompanyAccount"`
		} `xml:"CompanyAccountDataSet"`
	} `xml:"diffgram"`
}

type CompanyAccounts []CompanyAccount

type CompanyAccount struct {
	XMLName                        xml.Name `xml:"CompanyAccount" json:"-"`
	ID                             string   `xml:"ID,attr" json:"-"`
	RowOrder                       string   `xml:"RowOrder,attr" json:"-"`
	Account                        string   `xml:"Account" json:"account"`
	BankCode                       int64    `xml:"BankCode" json:"bankCode"`
	AccountNumber                  int64    `xml:"AccountNumber" json:"accountNumber"`
	ControlNumber                  int64    `xml:"ControlNumber" json:"controlNumber"`
	CompanyName                    string   `xml:"CompanyName" json:"companyName"`
	NationalIdentificationNumber   int64    `xml:"NationalIdentificationNumber" json:"nationalIdentificationNumber"`
	TaxIdentificationNumber        string   `xml:"TaxIdentificationNumber" json:"taxIdentificationNumber"`
	Address                        string   `xml:"Address" json:"address"`
	City                           string   `xml:"City" json:"city"`
	MunicipalityCode               int64    `xml:"MunicipalityCode" json:"municipalityCode"`
	ActivityCode                   int64    `xml:"ActivityCode" json:"activityCode"`
	MunicipalityName               string   `xml:"MunicipalityName" json:"municipalityName"`
	ActivityName                   string   `xml:"ActivityName" json:"activityName"`
	BankName                       string   `xml:"BankName" json:"bankName"`
	CompanyAccountStatusID         int64    `xml:"CompanyAccountStatusID" json:"companyAccountStatusID"`
	CompanyAccountBlockadeStatusID int64    `xml:"CompanyAccountBlockadeStatusID" json:"companyAccountBlockadeStatusID"`
	CompanyAccountTypeID           int64    `xml:"CompanyAccountTypeID" json:"companyAccountTypeID"`
	LegalUserTypeID                int64    `xml:"LegalUserTypeID" json:"legalUserTypeID"`
	InitializationDate             string   `xml:"InitializationDate" json:"initializationDate"`
	ChangeDate                     string   `xml:"ChangeDate" json:"changeDate"`
	UpdateDate                     string   `xml:"UpdateDate" json:"updateDate"`
	BankID                         string   `xml:"BankID" json:"bankID"`
}

func (a CompanyAccounts) WriteOut() {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("NAME", "CITY", "TAX ID", "NATIONAL ID", "BANK", "ACCOUNT", "STATUS ID", "TYPE ID")
	for _, row := range a {
		table.AddRow(row.CompanyName, row.City, row.TaxIdentificationNumber, row.NationalIdentificationNumber, row.BankName, row.Account, row.CompanyAccountStatusID, row.CompanyAccountTypeID)
	}
	fmt.Println(table)
}
