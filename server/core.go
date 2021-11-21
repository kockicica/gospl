package server

import (
	"net/http"
	"net/url"
	"strconv"

	"gospl/nbs/core/GetBank"
	"gospl/nbs/core/GetBankStatus"
	"gospl/nbs/core/GetBankType"

	"github.com/edermanoel94/rest-go"
)

func (h *handler) bankHandler(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//err := rest.CheckPathVariables(params, "bank-code")
	//if err != nil {
	//	rest.Error(w, err, http.StatusInternalServerError)
	//	return
	//}
	vals, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		rest.Error(w, err, http.StatusBadRequest)
		return
	}

	bankCode := 0
	if cBankCode := vals.Get("bank-code"); cBankCode != "" {
		bankCode, err = strconv.Atoi(cBankCode)
		if err != nil {
			rest.Error(w, unableToConvertToNumberError(cBankCode), http.StatusBadRequest)
			return
		}
	}

	data, err := h.nbsClient.GetBank(&GetBank.Request{BankCode: bankCode})
	if err != nil {
		rest.Error(w, err, http.StatusInternalServerError)
		return
	}

	if data == nil {
		data = GetBank.BankDataSetBanks{}
	}
	rest.Marshalled(w, &data, http.StatusOK)
}

func (h *handler) bankTypeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := h.nbsClient.GetBankType(&GetBankType.Request{})
	if err != nil {
		rest.Error(w, err, http.StatusInternalServerError)
	} else {
		rest.Marshalled(w, &data, http.StatusOK)
	}
}

func (h *handler) bankStatusHandler(w http.ResponseWriter, r *http.Request) {
	data, err := h.nbsClient.GetBankStatus(&GetBankStatus.Request{})
	if err != nil {
		rest.Error(w, err, http.StatusInternalServerError)
	} else {
		rest.Marshalled(w, &data, http.StatusOK)
	}
}
