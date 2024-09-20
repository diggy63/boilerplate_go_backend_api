package todos

import (
	"database/sql"

	"github.com/diggy63/boilerplate_go_api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateToDo(newToDo types.NewToDo) error {
	_, err := s.db.Exec("Insert Into to_do (list_id, title, is_done) Values ($1, $2, $3)", newToDo.ListID, newToDo.Title, newToDo.Completed)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetToDosByListID(listID int) ([]types.ToDo, error) {
	return nil, nil
}

func (s *Store) DeleteToDoByID(toDoID int) error {
	return nil
}

func (s *Store) UpdateToDoByID(toDoID int, completed bool) error {
	return nil
}
