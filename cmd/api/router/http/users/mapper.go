package users

import userDomain "github.com/kenethrrizzo/bookland-service/cmd/api/domain/users"

func userRequestToUserDomain(userRequest *UserRequest) *userDomain.User {
	return &userDomain.User{
		Name:        userRequest.Name,
		Surname:     userRequest.Surname,
		Email:       userRequest.Email,
		Username:    userRequest.Username,
		Password:    userRequest.Password,
		DateOfBirth: userRequest.DateOfBirth,
	}
}

func authDomainToUserResponse(authDomain *userDomain.Authentication) *UserResponse {
	return &UserResponse{
		AccessToken: authDomain.AccessToken,
	}
}
