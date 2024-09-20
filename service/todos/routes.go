package todos

import (
	"net/http"
	"strconv"

	"github.com/diggy63/boilerplate_go_api/service/auth"
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
	vars := mux.Vars(r)
	listID, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)

	}
	authHeader := r.Header.Get("Authorization")
	//get token from auth package
	token, err := auth.GetToken(authHeader)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	var payload types.NewToDoPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	_, err = auth.DecodeUserInfo(token)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	err = h.store.CreateToDo(types.NewToDo{ListID: listID, Title: payload.Title, Completed: false})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
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
