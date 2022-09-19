package users

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	domainErrors "github.com/kenethrrizzo/bookland-service/cmd/api/domain/errors"
	"golang.org/x/crypto/bcrypt"
)

var SECRET = []byte("texto-super-secreto") // Debe estar en variables de entorno

type UserService interface {
	Register(*User) (*Authentication, error)
	Login(*User) (*Authentication, error)
}

type Service struct {
	userRepo UserRepository
}

func NewService(userRepo UserRepository) *Service {
	return &Service{userRepo}
}

func (s *Service) Register(user *User) (*Authentication, error) {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	err = s.userRepo.Register(user)
	if err != nil {
		return nil, err
	}

	return createJWT(user)
}

func (s *Service) Login(user *User) (*Authentication, error) {
	userInDb, err := s.userRepo.Login(user)
	if err != nil {
		return nil, err
	}
	if !checkPasswordHash(user.Password, userInDb.Password) {
		return nil, domainErrors.NewAppErrorWithType(domainErrors.UnauthorizedError)
	}

	return &Authentication{
		AccessToken: "token",
	}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil
}

func createJWT(user *User) (*Authentication, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["userName"] = user.Username
	claims["userEmail"] = user.Email

	tokenStr, err := token.SignedString(SECRET)
	if err != nil {
		return nil, domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
	}

	return &Authentication{
		AccessToken: tokenStr,
	}, nil
}
