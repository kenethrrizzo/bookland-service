package books

type Response struct {
	Status string      `json:"status"`
	Result interface{} `json:"result"`
}

type BookResponse struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Author    string   `json:"author"`
	Genres    []string `json:"genres"`
	CoverPage string   `json:"coverpage"`
	Synopsis  string   `json:"synopsis"`
	Price     float64  `json:"price"`
}
