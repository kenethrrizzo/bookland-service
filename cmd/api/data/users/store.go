package users

import (
	"database/sql"

	domainErrors "github.com/kenethrrizzo/bookland-service/cmd/api/domain/errors"
	userDomain "github.com/kenethrrizzo/bookland-service/cmd/api/domain/users"
	"github.com/pkg/errors"
)

const (
	readError = "error buscando un libro en la base de datos"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s Store) Register(user *userDomain.User) error {
	sqlInsert := `INSERT INTO User (Name, Surname, Email, Username, Password, DateOfBirth) VALUES (?, ?, ?, ?, ?, ?)`

	userDB := userDomainToUserSchema(user)

	_, err := s.db.Exec(sqlInsert, userDB.Name, userDB.Surname, userDB.Email, userDB.Username, userDB.Password, userDB.DateOfBirth)

	return err
}

func (s Store) Login(user *userDomain.User) (*userDomain.User, error) {
	sqlSelect := `SELECT Email, Username, Password FROM User WHERE Email = ?`

	userDB := userDomainToUserSchema(user)
	err := s.db.QueryRow(sqlSelect, userDB.Email).Scan(&userDB.Email, &userDB.Username, &userDB.Password)
	switch err {
	default:
		appErr := domainErrors.NewAppError(errors.Wrap(err, readError), domainErrors.RepositoryError)
		return nil, appErr
	case sql.ErrNoRows:
		appErr := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
		return nil, appErr
	case nil:
		return userSchemaToUserDomain(userDB), nil
	}
}
