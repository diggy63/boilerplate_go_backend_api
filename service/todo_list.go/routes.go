package todo_list

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/diggy63/boilerplate_go_api/service/auth"
	"github.com/diggy63/boilerplate_go_api/types"
	"github.com/diggy63/boilerplate_go_api/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ToDoListStore
}

func NewHandler(store types.ToDoListStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/todo_lists", h.handleGetToDoLists).Methods("GET")
	router.HandleFunc("/todo_lists", h.handleCreateToDoList).Methods("POST")
	router.HandleFunc("/todo_lists/{id}", h.handleDeleteToDoList).Methods("DELETE")
}

func (h *Handler) handleCreateToDoList(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	//get token from auth package
	token, err := auth.GetToken(authHeader)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	//get json payload
	var payload types.NewToDoListPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	//decode jwt using auth package
	user_id, err := auth.DecodeUserInfo(token)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	//create todo list with foriegn key of user id
	err = h.store.CreateToDoList(types.NewToDoList{UserID: user_id, Title: payload.Title})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	//return status ok with message
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "todo list created"})
}

func (h *Handler) handleGetToDoLists(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	//get token from auth package
	token, err := auth.GetToken(authHeader)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	user_id, err := auth.DecodeUserInfo(token)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	toDoLists, err := h.store.GetToDoListsByUserID(user_id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error getting todo list: %v", err))
	}
	utils.WriteJSON(w, http.StatusOK, toDoLists)
}

func (h *Handler) handleDeleteToDoList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	list_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error converting id to int: %v", err))
	}
	err = h.store.DeleteToDoListByID(list_id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error deleting todo list: %v", err))
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "delete todo list"})
}
