package users

import "time"

type User struct {
	Id          int       `db:"Id"`
	Name        string    `db:"Name"`
	Surname     string    `db:"Surname"`
	Email       string    `db:"Email"`
	Username    string    `db:"Username"`
	Password    string    `db:"Password"`
	DateOfBirth time.Time `db:"DateOfBirth"`
}
