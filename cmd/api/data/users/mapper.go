package users

import userDomain "github.com/kenethrrizzo/bookland-service/cmd/api/domain/users"

func userDomainToUserSchema(user *userDomain.User) *User {
	return &User{
		Id:          user.Id,
		Name:        user.Name,
		Surname:     user.Surname,
		Email:       user.Email,
		Username:    user.Username,
		Password:    user.Password,
		DateOfBirth: user.DateOfBirth,
	}
}

// update
func userSchemaToUserDomain(user *User) *userDomain.User {
	return &userDomain.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}
