package server

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/kockicica/gospl/nbs/account/GetCompanyAccount"
	"github.com/kockicica/gospl/nbs/account/GetCompanyAccountStatus"
	"github.com/kockicica/gospl/nbs/account/GetCompanyAccountType"

	"github.com/edermanoel94/rest-go"
	"github.com/pkg/errors"
)

func (h *handler) accountTypeHandler(w http.ResponseWriter, r *http.Request) {

	data, err := h.nbsClient.GetCompanyAccountType(&GetCompanyAccountType.Request{})
	if err != nil {
		rest.Error(w, err, http.StatusInternalServerError)
	} else {
		rest.Marshalled(w, &data, http.StatusOK)
	}
}

func (h *handler) accountStatusHandler(w http.ResponseWriter, r *http.Request) {
	data, err := h.nbsClient.GetCompanyAccountStatus(&GetCompanyAccountStatus.Request{})
	if err != nil {
		rest.Error(w, err, http.StatusInternalServerError)
	} else {
		rest.Marshalled(w, &data, http.StatusOK)
	}
}

func (h *handler) accountHandler(w http.ResponseWriter, r *http.Request) {

	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		rest.Error(w, err, http.StatusInternalServerError)
		return
	}

	query := &GetCompanyAccount.Request{}

	if v := params.Get("name"); v != "" {
		query.CompanyName = v
	}
	if v := params.Get("national-id"); v != "" {
		nationalId, err := strconv.Atoi(v)
		if err != nil {
			rest.Error(w, unableToConvertToNumberError(v), http.StatusInternalServerError)
			return
		}
		query.NationalIdentificationNumber = nationalId
	}
	if v := params.Get("tax-id"); v != "" {
		query.TaxIdentificationNumber = v
	}
	if v := params.Get("bank-code"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil {
			rest.Error(w, unableToConvertToNumberError(v), http.StatusInternalServerError)
			return
		}
		query.BankCode = n
	}
	if v := params.Get("account-number"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil {
			rest.Error(w, unableToConvertToNumberError(v), http.StatusInternalServerError)
			return
		}
		query.AccountNumber = n
	}
	if v := params.Get("control-number"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil {
			rest.Error(w, unableToConvertToNumberError(v), http.StatusInternalServerError)
			return
		}
		query.ControlNumber = n
	}
	if v := params.Get("city"); v != "" {
		query.City = v
	}

	data, err := h.nbsClient.GetCompanyAccount(query)
	if err != nil {
		rest.Error(w, err, http.StatusInternalServerError)
		return
	}

	if data == nil {
		data = GetCompanyAccount.CompanyAccounts{}
	}

	rest.Marshalled(w, &data, http.StatusOK)

}

func unableToConvertToNumberError(number string) error {
	return errors.Errorf("Unable to convert %s to number", number)
}
