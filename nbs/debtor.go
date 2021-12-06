package nbs

import (
	"gospl/nbs/debtor/GetDebtor"
	"gospl/nbs/debtor/GetEnforcedCollectionDebtorBlockadeStatus"
)

func (c *Client) GetDebtor(request *GetDebtor.Request) (GetDebtor.Debtors, error) {
	response := GetDebtor.Response{}
	if err := c.handle("CommunicationOfficeService1_0/DebtorService.asmx", "GetDebtor", request, &response); err != nil {
		return nil, err
	}
	return response.Result.Content.Data.Debtors, nil
}

func (c *Client) GetEnforcedCollectionDebtorBlockadeStatus(request *GetEnforcedCollectionDebtorBlockadeStatus.Request) (GetEnforcedCollectionDebtorBlockadeStatus.Statuses, error) {
	response := GetEnforcedCollectionDebtorBlockadeStatus.Response{}
	if err := c.handle("CommunicationOfficeService1_0/DebtorService.asmx", "GetEnforcedCollectionDebtorBlockadeStatus", request, &response); err != nil {
		return nil, err
	}
	return response.Result.Content.Data.Statuses, nil
}
