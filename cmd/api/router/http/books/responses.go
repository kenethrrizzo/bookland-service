package books

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type BookResponse struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Author    int      `json:"author"`
	Genres    []string `json:"genres"`
	CoverPage string   `json:"coverpage"`
	Synopsis  string   `json:"synopsis"`
	Price     float64  `json:"price"`
}
