package GetBank

import (
	"encoding/xml"
	"fmt"

	"github.com/gosuri/uitable"
)

type Request struct {
	XMLName                      xml.Name `xml:"http://communicationoffice.nbs.rs GetBank"`
	BankID                       string   `xml:"bankID,omitempty"`
	BankCode                     int      `xml:"bankCode,omitempty"`
	NationalIdentificationNumber int      `xml:"nationalIdentificationNumber,omitempty"`
	TaxIdentificationNumber      string   `xml:"taxIdentificationNumber,omitempty"`
	Date                         string   `xml:"date,omitempty"`
}

type Response struct {
	XMLName xml.Name           `xml:"http://communicationoffice.nbs.rs GetBankResponse"`
	Result  ResponseTypeResult `xml:"GetBankResult"`
}

type ResponseTypeResult struct {
	Content struct {
		Data struct {
			Banks BankDataSetBanks `xml:"Bank"`
		} `xml:"BankDataSet"`
	} `xml:"diffgram"`
}

type BankDataSetBanks []BankDataSetBank

type BankDataSetBank struct {
	//XMLName                      xml.Name `xml:"Bank"`
	BankID                       string `xml:"BankID" json:"bankID"`
	BankHistoryID                string `xml:"BankHistoryID" json:"bankHistoryID"`
	StartDate                    string `xml:"StartDate" json:"startDate"`
	EndDate                      string `xml:"EndDate" json:"endDate"`
	BankCode                     int64  `xml:"BankCode" json:"bankCode"`
	NationalIdentificationNumber int64  `xml:"NationalIdentificationNumber" json:"nationalIdentificationNumber"`
	TaxIdentificationNumber      string `xml:"TaxIdentificationNumber" json:"taxIdentificationNumber"`
	Name                         string `xml:"Name" json:"name"`
	LogoSmall                    string `xml:"LogoSmall" json:"logoSmall"`
	LogoBig                      string `xml:"LogoBig" json:"logoBig"`
	Address                      string `xml:"Address" json:"address"`
	City                         string `xml:"City" json:"city"`
	Region                       string `xml:"Region" json:"region"`
	State                        string `xml:"State" json:"state"`
	Country                      string `xml:"Country" json:"country"`
	PostalCode                   string `xml:"PostalCode" json:"postalCode"`
	Phone                        string `xml:"Phone" json:"phone"`
	Fax                          string `xml:"Fax" json:"fax"`
	Email                        string `xml:"Email" json:"email"`
	WebAddress                   string `xml:"WebAddress" json:"webAddress"`
	Director                     string `xml:"Director" json:"director"`
	Description                  string `xml:"Description" json:"description"`
	BankTypeID                   int    `xml:"BankTypeID" json:"bankTypeID"`
	BankStatusID                 int    `xml:"BankStatusID" json:"bankStatusID"`
}

func (b BankDataSetBanks) WriteOut() {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("ID", "NAME", "CODE", "DESCRIPTION", "NATIONAL ID", "TAX ID")
	for _, r := range b {
		table.AddRow(r.BankID, r.Name, r.BankCode, r.Description, r.NationalIdentificationNumber, r.TaxIdentificationNumber)
	}
	fmt.Println(table)

}
