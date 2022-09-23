package auth

type Response struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}
