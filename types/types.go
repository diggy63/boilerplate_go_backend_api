package types

import "time"

// User Types
type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=100"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

//ToDoList Types

type NewToDoPayload struct {
	Title string `json:"title"`
}

type NewToDoList struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
}

type ToDoList struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type ToDoListStore interface {
	GetToDoListsByUserID(userID int) ([]ToDoList, error)
	DeleteToDoListByID(listID int) error
	CreateToDoList(NewToDoList) error
}

//ToDo Types

type ToDoStore interface {
	CreateToDo(toDoID int) error
	GetToDosByListID(listID int) ([]ToDo, error)
	DeleteToDoByID(todDoID int) error
	UpdateToDoByID(toDoID int, completed bool) error
}

type ToDo struct {
	ID        int       `json:"id"`
	ListID    int       `json:"list_id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
