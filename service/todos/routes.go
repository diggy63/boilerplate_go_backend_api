package todos

import (
	"net/http"

	"github.com/diggy63/boilerplate_go_api/types"
	"github.com/diggy63/boilerplate_go_api/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ToDoStore
}

func NewHandler(store types.ToDoStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/todo_lists/{id}/todos", h.handleGetToDos).Methods("GET")
	router.HandleFunc("/todo_lists/{id}/todos", h.handleCreateToDo).Methods("POST")
	router.HandleFunc("/todo_lists/{id}/todos/{todo_id}", h.handleDeleteToDo).Methods("DELETE")
	router.HandleFunc("/todo_lists/{id}/todos/{todo_id}", h.handleUpdateToDo).Methods("PUT")
}

func (h *Handler) handleCreateToDo(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "todo created"})
}

func (h *Handler) handleGetToDos(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "getting todos"})
}

func (h *Handler) handleDeleteToDo(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "delete todo"})
}

func (h *Handler) handleUpdateToDo(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "updated todo"})
}
