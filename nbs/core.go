package nbs

import (
	"gospl/nbs/core/GetBank"
	"gospl/nbs/core/GetBankStatus"
	"gospl/nbs/core/GetBankType"
)

func (c *Client) GetBank(request *GetBank.Request) (GetBank.BankDataSetBanks, error) {
	response := GetBank.Response{}
	if err := c.handle("CommunicationOfficeService1_0/CoreService.asmx", "GetBank", request, &response); err != nil {
		return nil, err
	}
	return response.Result.Content.Data.Banks, nil
}

func (c *Client) GetBankStatus(request *GetBankStatus.Request) (GetBankStatus.BankDataSetBankStatuses, error) {
	response := GetBankStatus.Response{}
	if err := c.handle("CommunicationOfficeService1_0/CoreService.asmx", "GetBankStatus", request, &response); err != nil {
		return nil, err
	}
	return response.Result.Content.Data.Statuses, nil
}

func (c *Client) GetBankType(request *GetBankType.Request) (GetBankType.BankTypes, error) {
	response := GetBankType.Response{}
	if err := c.handle("CommunicationOfficeService1_0/CoreService.asmx", "GetBankType", request, &response); err != nil {
		return nil, err
	}
	return response.Result.Content.Data.Types, nil
}
