package books

type BookRequest struct {
	Id        int     `json:"id"`
	Name      string  `json:"name" binding:"required"`
	Author    int     `json:"author" binding:"required"`
	Coverpage string  `json:"coverpage"`
	Synopsis  string  `json:"synopsis"`
	Price     float64 `json:"price" binding:"required"`
}
