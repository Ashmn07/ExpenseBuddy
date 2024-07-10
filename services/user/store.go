package user

import (
	"database/sql"

	"github.com/Ashmn07/ExpenseBuddy/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
func (s *Store) GetUserByEmail(email string) (*types.User, error) {

	return nil, nil
}
func (s *Store) GetUserByID(id int) (*types.User, error) {

	return nil, nil
}
