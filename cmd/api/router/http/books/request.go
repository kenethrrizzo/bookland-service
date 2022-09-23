package books

import (
	"mime/multipart"
)

type BookRequest struct {
	ID        *int                  `form:"id"` // TODO: eliminar
	Name      string                `form:"name"`
	Author    string                `form:"author"`
	Genres    string                `form:"genres"`
	Coverpage *multipart.FileHeader `form:"coverpage"`
	Synopsis  string                `form:"synopsis"`
	Price     float64               `form:"price"`
}
