package users

import (
	"time"

	userDomain "github.com/kenethrrizzo/bookland-service/cmd/api/domain/users"
)

// TODO: Mejorar método
func userRequestToUserDomain(userRequest *UserRequest) (*userDomain.User, error) {
	if userRequest.DateOfBirth != "" {
		dateOfBirth, err := time.Parse("2006-01-02", userRequest.DateOfBirth)
		if err != nil {
			return nil, err
		}

		return &userDomain.User{
			Name:        userRequest.Name,
			Surname:     userRequest.Surname,
			Email:       userRequest.Email,
			Username:    userRequest.Username,
			Password:    userRequest.Password,
			DateOfBirth: dateOfBirth,
		}, nil
	}

	return &userDomain.User{
		Name:     userRequest.Name,
		Surname:  userRequest.Surname,
		Email:    userRequest.Email,
		Username: userRequest.Username,
		Password: userRequest.Password,
	}, nil
}

func authDomainToUserResponse(authDomain *userDomain.Authentication) *UserResponse {
	return &UserResponse{
		AccessToken: authDomain.AccessToken,
	}
}
