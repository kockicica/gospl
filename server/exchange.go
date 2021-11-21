package server

import (
	"net/http"
	"net/url"
	"strconv"

	"gospl/nbs/exchange/GetCurrentExchangeRate"
	"gospl/nbs/exchange/GetExchangeRateByCurrency"
	"gospl/nbs/exchange/GetExchangeRateListType"

	"github.com/edermanoel94/rest-go"
)

func (h *handler) exchangeRateListType(w http.ResponseWriter, r *http.Request) {
	data, err := h.nbsClient.GetExchangeRateListType(&GetExchangeRateListType.Request{})
	if err != nil {
		rest.Error(w, err, http.StatusInternalServerError)
	} else {
		rest.Marshalled(w, &data, http.StatusOK)
	}
}

func (h *handler) currentExchangeRate(w http.ResponseWriter, r *http.Request) {

	query := &GetCurrentExchangeRate.Request{}

	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		rest.Error(w, err, http.StatusBadRequest)
		return
	}
	if v := params.Get("list-type"); v != "" {
		listType, err := strconv.Atoi(v)
		if err != nil {
			rest.Error(w, unableToConvertToNumberError(v), http.StatusBadRequest)
			return
		}
		query.ExchangeRateListTypeID = listType
	}

	data, err := h.nbsClient.GetCurrentExchangeRate(query)
	if err != nil {
		rest.Error(w, err, http.StatusInternalServerError)
	} else {
		if data == nil {
			data = GetCurrentExchangeRate.ExchangeRateDataSetExchangeRateList{}
		}
		rest.Marshalled(w, &data, http.StatusOK)
	}
}

func (h *handler) currentExchangeRateRsdEur(w http.ResponseWriter, r *http.Request) {
	data, err := h.nbsClient.GetCurrentExchangeRateRsdEur()
	if err != nil {
		rest.Error(w, err, http.StatusInternalServerError)
		return
	}

	rest.Marshalled(w, &data, http.StatusOK)

}

func (h *handler) exchangeRateByCurrency(w http.ResponseWriter, r *http.Request) {

	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		rest.Error(w, err, http.StatusBadRequest)
		return
	}
	query := &GetExchangeRateByCurrency.Request{}

	if v := params.Get("list-type"); v != "" {
		listType, err := strconv.Atoi(v)
		if err != nil {
			rest.Error(w, unableToConvertToNumberError(v), http.StatusBadRequest)
			return
		}
		query.ExchangeRateListTypeID = listType
	}
	if v := params.Get("currency-code"); v != "" {
		listType, err := strconv.Atoi(v)
		if err != nil {
			rest.Error(w, unableToConvertToNumberError(v), http.StatusBadRequest)
			return
		}
		query.CurrencyCode = listType
	}

	if v := params.Get("date-from"); v != "" {
		query.DateFrom = v
	}
	if v := params.Get("date-to"); v != "" {
		query.DateTo = v
	}

	data, err := h.nbsClient.GetExchangeRateByCurrency(query)
	if err != nil {
		rest.Error(w, err, http.StatusInternalServerError)
		return
	}

	rest.Marshalled(w, &data, http.StatusOK)

}
