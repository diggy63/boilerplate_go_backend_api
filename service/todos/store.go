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

func (s *Store) CreateToDo(toDoID int) error {
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
