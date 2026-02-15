package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vishalpaliwal79/ecom/services/user"
)

type APIserver struct {
	addr string
	db   *sql.DB
}

func newAPIServer(addr string, db *sql.DB) *APIserver {
	return &APIserver{
		addr: addr,
		db:   db,
	}
}

func (s *APIserver) run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("api/v1").Subrouter()
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
