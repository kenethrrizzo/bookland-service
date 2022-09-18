package users

type UserRepository interface {
	Register(*User) (*Authentication, error)
	Login(*User) (*Authentication, error)
}
