package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"gospl/nbs"

	"github.com/edermanoel94/rest-go"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type Server struct {
	server   *http.Server
	listener net.Listener
	port     int
}

type handler struct {
	nbsClient *nbs.Client
}

func (h *handler) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	rest.Error(w, errors.New("Not found"), http.StatusNotFound)
}

func (s *Server) Start() error {

	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	log.Print("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		return err
	}
	log.Print("Shutting down server done.")
	return nil
}

func New(port int, baseUrl, username, password, licenceId string) *Server {

	nbsClient := nbs.NewClient(baseUrl, username, password, licenceId)
	handler := handler{nbsClient: nbsClient}
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(handler.notFoundHandler)
	accountRouter := r.PathPrefix("/account").Methods(http.MethodGet).Subrouter()
	accountRouter.HandleFunc("", handler.accountHandler)
	accountRouter.HandleFunc("/type", handler.accountTypeHandler)
	accountRouter.HandleFunc("/status", handler.accountStatusHandler)

	coreSubrouter := r.PathPrefix("/core").Methods(http.MethodGet).Subrouter()
	coreSubrouter.HandleFunc("/currency", handler.currencyHandler)

	bankSubrouter := coreSubrouter.PathPrefix("/bank").Subrouter()
	bankSubrouter.HandleFunc("", handler.bankHandler)
	bankSubrouter.HandleFunc("/status", handler.bankStatusHandler)
	bankSubrouter.HandleFunc("/type", handler.bankTypeHandler)

	exchangeSubrouter := r.PathPrefix("/exchange").Subrouter()
	exchangeSubrouter.HandleFunc("/list-type", handler.exchangeRateListType)
	exchangeSubrouter.HandleFunc("/current", handler.currentExchangeRate)
	exchangeSubrouter.HandleFunc("/rsd-eur/current", handler.currentExchangeRateRsdEur)
	exchangeSubrouter.HandleFunc("/currency", handler.exchangeRateByCurrency)

	srv := &Server{
		server: &http.Server{
			Addr: fmt.Sprintf(":%d", port),
		},
	}
	srv.server.Handler = r
	return srv
}
