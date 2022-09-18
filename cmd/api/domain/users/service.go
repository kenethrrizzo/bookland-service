package users

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
	return s.userRepo.Register(user)
}

func (s *Service) Login(user *User) (*Authentication, error) {
	return s.userRepo.Login(user)
}
