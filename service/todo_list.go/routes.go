package todo_list

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/diggy63/boilerplate_go_api/service/auth"
	"github.com/diggy63/boilerplate_go_api/types"
	"github.com/diggy63/boilerplate_go_api/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Handler struct {
	store types.ToDoListStore
}

func NewHandler(store types.ToDoListStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/todo_lists/{api_key}", h.handleGetToDoLists).Methods("GET")
	router.HandleFunc("/todo_lists", h.handleCreateToDoList).Methods("POST")
	router.HandleFunc("/todo_lists/{id}", h.handleDeleteToDoList).Methods("DELETE")
}

func (h *Handler) handleCreateToDoList(w http.ResponseWriter, r *http.Request) {
	secret, err := getSecret(w)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	//get json payload
	var payload types.NewToDoPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	//decode jwt using auth package
	user_id, err := auth.DecodeUserInfo([]byte(secret), payload.Apikey)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	//turn string id into int
	user_id_int, err := strconv.Atoi(user_id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	//create todo list with foriegn key of user id
	err = h.store.CreateToDoList(types.NewToDoList{UserID: user_id_int, Title: payload.Title})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	//return status ok with message
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "todo list created"})
}

func (h *Handler) handleGetToDoLists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apikey := vars["api_key"]
	secret, err := getSecret(w)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
	user_id, err := auth.DecodeUserInfo([]byte(secret), apikey)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	//turn string id into int
	user_id_int, err := strconv.Atoi(user_id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("err"))
	}
	toDoLists, err := h.store.GetToDoListsByUserID(user_id_int)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error getting todo list: %v", err))
	}
	fmt.Println(toDoLists)
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

// handles getting our secret
func getSecret(w http.ResponseWriter) (string, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error loading .env file: %v", err))
		return "", err
	}
	secret := os.Getenv("SECRET_JWT")
	if secret == "" {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("SECRET_JWT environment variable not set"))
		return "", err
	}
	return secret, nil
}
