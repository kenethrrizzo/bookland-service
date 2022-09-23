package users

type Response struct {
	Status string      `json:"status"`
	Result   interface{} `json:"result"`
}

type UserResponse struct {
	AccessToken string `json:"accessToken"`
}
