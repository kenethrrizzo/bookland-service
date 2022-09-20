package users

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type UserResponse struct {
	AccessToken string `json:"accessToken"`
}
