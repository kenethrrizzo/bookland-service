package users

type UserRepository interface {
	Register(*User) error
	Login(*User) (*User, error)
}
