package nbs

import (
	"github.com/kockicica/gospl/nbs/account/GetCompanyAccount"
	"github.com/kockicica/gospl/nbs/account/GetCompanyAccountStatus"
	"github.com/kockicica/gospl/nbs/account/GetCompanyAccountType"
)

func (c *Client) GetCompanyAccount(request *GetCompanyAccount.Request) (GetCompanyAccount.CompanyAccounts, error) {

	response := GetCompanyAccount.Response{}
	if err := c.handle("CommunicationOfficeService1_0/CompanyAccountService.asmx", "GetCompanyAccount", request, &response); err != nil {
		return nil, err
	}
	return response.Result.Content.Data.Accounts, nil

}

func (c *Client) GetCompanyAccountStatus(request *GetCompanyAccountStatus.Request) (GetCompanyAccountStatus.CompanyAccountStatuses, error) {
	response := GetCompanyAccountStatus.Response{}
	if err := c.handle("CommunicationOfficeService1_0/CompanyAccountService.asmx", "GetCompanyAccountStatus", request, &response); err != nil {
		return nil, err
	}
	return response.Result.Content.Data.Statuses, nil
}

func (c *Client) GetCompanyAccountType(request *GetCompanyAccountType.Request) (GetCompanyAccountType.CompanyAccountTypes, error) {
	response := GetCompanyAccountType.Response{}
	if err := c.handle("CommunicationOfficeService1_0/CompanyAccountService.asmx", "GetCompanyAccountType", request, &response); err != nil {
		return nil, err
	}
	return response.Result.Content.Data.Types, nil
}
