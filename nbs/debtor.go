package nbs

import (
	"github.com/kockicica/gospl/nbs/debtor/GetDebtor"
)

func (c *Client) GetDebtor(request *GetDebtor.Request) (GetDebtor.Debtors, error) {
	response := GetDebtor.Response{}
	if err := c.handle("CommunicationOfficeService1_0/DebtorService.asmx", "GetDebtor", request, &response); err != nil {
		return nil, err
	}
	return response.Result.Content.Data.Debtors, nil
}
