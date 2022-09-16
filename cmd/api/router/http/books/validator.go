package books

import (
	"mime/multipart"
)

type BookRequest struct {
	Id        *int                  `form:"id"`
	Name      string                `form:"name"`
	Author    int                   `form:"author"`
	Coverpage *multipart.FileHeader `form:"coverpage"`
	Synopsis  string                `form:"synopsis"`
	Price     float64               `form:"price"`
}
