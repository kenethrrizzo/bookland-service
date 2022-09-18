package users

import "time"

type User struct {
	Id          int
	Name        string
	Surname     string
	Email       string
	Username    string
	Password    string
	DateOfBirth time.Time
}

type Authentication struct {
	AccessToken string
	// TODO: RefreshToken
}
