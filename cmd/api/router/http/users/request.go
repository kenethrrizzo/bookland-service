package users

type UserRequest struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	DateOfBirth string `json:"dateOfBirth"`
}
