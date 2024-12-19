package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MatthewAraujo/anki_ia/repository"
	"github.com/MatthewAraujo/anki_ia/service/anki"
	"github.com/MatthewAraujo/anki_ia/service/users"
	"github.com/gorilla/mux"
	"github.com/openai/openai-go"
	"github.com/redis/go-redis/v9"
	"github.com/rs/cors"
)

type APIServer struct {
	addr         string
	db           *repository.Queries
	dbTx         *sql.DB
	redis        *redis.Client
	openAiClient *openai.Client
}

func NewAPIServer(addr string, db *repository.Queries, dbTx *sql.DB, redis *redis.Client, openAiClient *openai.Client) *APIServer {
	return &APIServer{
		addr:         addr,
		db:           db,
		dbTx:         dbTx,
		redis:        redis,
		openAiClient: openAiClient,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	router.Use(c.Handler)
	// if the api changes in the future we can just change the version here, and the old version will still be available
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	costumersService := users.NewService(s.db, s.dbTx)
	usersHandler := users.NewHandler(costumersService)
	usersHandler.RegisterRoutes(subrouter.PathPrefix("/users").Subrouter())

	ankiService := anki.NewService(s.db, s.dbTx, s.openAiClient)
	ankiHandler := anki.NewHandler(ankiService, *s.db)
	ankiHandler.RegisterRoutes(subrouter.PathPrefix("/anki").Subrouter())

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
