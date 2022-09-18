package users

import (
	"database/sql"

	userDomain "github.com/kenethrrizzo/bookland-service/cmd/api/domain/users"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s Store) Register(user *userDomain.User) (*userDomain.Authentication, error) {
	// TODO: Implement method
	return nil, nil
}

func (s Store) Login(user *userDomain.User) (*userDomain.Authentication, error) {
	// TODO: Implement method
	return nil, nil
}
