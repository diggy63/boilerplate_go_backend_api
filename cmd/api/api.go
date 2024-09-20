package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/diggy63/boilerplate_go_api/service/health"
	"github.com/diggy63/boilerplate_go_api/service/todo_list.go"
	"github.com/diggy63/boilerplate_go_api/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	//Register User Handlers
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)
	//Register ToDoList Handlers
	todoListStore := todo_list.NewStore(s.db)
	todoListHandler := todo_list.NewHandler(todoListStore)
	todoListHandler.RegisterRoutes(subrouter)
	//register health handler
	healthHandler := health.NewHandler()
	healthHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
