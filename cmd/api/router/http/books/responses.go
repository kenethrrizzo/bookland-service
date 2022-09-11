package books

type BookResponse struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Author    int     `json:"author"`
	CoverPage string  `json:"coverpage"`
	Synopsis  string  `json:"synopsis"`
	Price     float64 `json:"price"`
}

type MessageResponse struct {
	Message string `json:"message"`
}