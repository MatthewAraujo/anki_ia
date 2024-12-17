package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MatthewAraujo/anki_ia/repository"
	"github.com/MatthewAraujo/anki_ia/service/anki"
	"github.com/MatthewAraujo/anki_ia/service/users"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

type APIServer struct {
	addr  string
	db    *repository.Queries
	dbTx  *sql.DB
	redis *redis.Client
}

func NewAPIServer(addr string, db *repository.Queries, dbTx *sql.DB, redis *redis.Client) *APIServer {
	return &APIServer{
		addr:  addr,
		db:    db,
		dbTx:  dbTx,
		redis: redis,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	// if the api changes in the future we can just change the version here, and the old version will still be available
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	costumersService := users.NewService(s.db, s.dbTx)
	customersHandler := users.NewHandler(costumersService)
	customersHandler.RegisterRoutes(subrouter.PathPrefix("/customers").Subrouter())

	ankiService := anki.NewService(s.db, s.dbTx)
	ankiHandler := anki.NewHandler(ankiService)
	ankiHandler.RegisterRoutes(subrouter.PathPrefix("/anki").Subrouter())

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
