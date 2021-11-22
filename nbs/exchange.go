package nbs

import (
	"github.com/kockicica/gospl/nbs/exchange/GetCurrentExchangeRate"
	"github.com/kockicica/gospl/nbs/exchange/GetCurrentExchangeRateRsdEur"
	"github.com/kockicica/gospl/nbs/exchange/GetExchangeRateByCurrency"
	"github.com/kockicica/gospl/nbs/exchange/GetExchangeRateListType"
)

func (c *Client) GetCurrentExchangeRate(request *GetCurrentExchangeRate.Request) (GetCurrentExchangeRate.ExchangeRateDataSetExchangeRateList, error) {

	response := GetCurrentExchangeRate.Response{}

	if err := c.handle("CommunicationOfficeService1_0/ExchangeRateService.asmx", "GetCurrentExchangeRate", request, &response); err != nil {
		return nil, err
	}

	return response.Result.Content.Data.ExchangeRates, nil
}

func (c *Client) GetCurrentExchangeRateRsdEur() (GetCurrentExchangeRateRsdEur.ExchangeRateRsdEur, error) {

	request := &GetCurrentExchangeRateRsdEur.Request{}
	response := GetCurrentExchangeRateRsdEur.Response{}

	if err := c.handle("CommunicationOfficeService1_0/ExchangeRateService.asmx", "GetCurrentExchangeRateRsdEur", request, &response); err != nil {
		return GetCurrentExchangeRateRsdEur.ExchangeRateRsdEur{}, err
	}
	return response.Result.Content.Data.ExchangeRate, nil
}

func (c *Client) GetExchangeRateListType(request *GetExchangeRateListType.Request) (GetExchangeRateListType.ExchangeRateListTypes, error) {
	response := GetExchangeRateListType.Response{}
	if err := c.handle("CommunicationOfficeService1_0/ExchangeRateService.asmx", "GetExchangeRateListType", request, &response); err != nil {
		return GetExchangeRateListType.ExchangeRateListTypes{}, err
	}
	return response.Result.Content.Data.ListTypes, nil
}

func (c *Client) GetExchangeRateByCurrency(request *GetExchangeRateByCurrency.Request) (GetExchangeRateByCurrency.ExchangeRateDataSetExchangeRates, error) {
	response := GetExchangeRateByCurrency.Response{}
	if err := c.handle("CommunicationOfficeService1_0/ExchangeRateService.asmx", "GetExchangeRateByCurrency", request, &response); err != nil {
		return GetExchangeRateByCurrency.ExchangeRateDataSetExchangeRates{}, err
	}
	return response.Result.Content.Data.ExchangeRates, nil
}
