package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"gospl/nbs"

	"github.com/gorilla/mux"
)

type Server struct {
	server   *http.Server
	listener net.Listener
	port     int
}

type handler struct {
	nbsClient *nbs.Client
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
	r.HandleFunc("/account", handler.accountHandler)
	r.HandleFunc("/account/type", handler.accountTypeHandler)
	r.HandleFunc("/account/status", handler.accountStatusHandler)
	r.HandleFunc("/core/bank", handler.bankHandler)
	r.HandleFunc("/core/bank/status", handler.bankStatusHandler)
	r.HandleFunc("/core/bank/type", handler.bankTypeHandler)
	r.HandleFunc("/exchange/list-type", handler.exchangeRateListType)
	r.HandleFunc("/exchange/current", handler.currentExchangeRate)
	r.HandleFunc("/exchange/rsd-eur/current", handler.currentExchangeRateRsdEur)
	r.HandleFunc("/exchange/currency", handler.exchangeRateByCurrency)

	srv := &Server{
		server: &http.Server{
			Addr: fmt.Sprintf(":%d", port),
		},
	}
	srv.server.Handler = r
	return srv
}
