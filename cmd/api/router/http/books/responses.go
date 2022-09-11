package books

type BookResponse struct {
	Name      string  `json:"name"`
	Author    int     `json:"author"`
	CoverPage string  `json:"coverpage"`
	Synopsis  string  `json:"synopsis"`
	Price     float64 `json:"price"`
}
