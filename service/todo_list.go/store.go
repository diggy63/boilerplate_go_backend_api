package todo_list

import (
	"database/sql"

	"github.com/diggy63/boilerplate_go_api/types"
)

type Store struct {
	db *sql.DB
}

type ToDoListCollection struct {
	Lists []types.ToDoList
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetToDoListsByUserID(userID int) ([]types.ToDoList, error) {
	rows, err := s.db.Query("Select * From to_do_list Where user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var toDoLists []types.ToDoList
	for rows.Next() {
		var toDoList types.ToDoList
		err := rows.Scan(&toDoList.ID, &toDoList.UserID, &toDoList.Title, &toDoList.CreatedAt)
		if err != nil {
			return nil, err
		}
		toDoLists = append(toDoLists, toDoList)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return toDoLists, nil
}

func (s *Store) DeleteToDoListByID(listID int) error {
	_, err := s.db.Exec("Delete From to_do_list Where id = $1", listID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateToDoList(newTodo types.NewToDoList) error {
	_, err := s.db.Exec("Insert Into to_do_list (user_id, title) Values ($1, $2)", newTodo.UserID, newTodo.Title)
	if err != nil {
		return err
	}
	return nil
}
