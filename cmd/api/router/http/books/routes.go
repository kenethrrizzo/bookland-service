package books

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/kenethrrizzo/bookland-service/cmd/api/domain/books"
)

type BookHandler struct {
	service books.BookService
}

func NewHandler(svc books.BookService) *BookHandler {
	return &BookHandler{svc}
}

func (handl *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	bookIDstr := r.URL.Query().Get("bookID")

	log.Println("bookID: " + bookIDstr)
	bookID, err := strconv.Atoi(bookIDstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ha ocurrido un error al convertir parametro a entero")
		// TODO: Implementar manejo de errores como respuesta
		return
	}

	result, err := handl.service.GetBookByID(bookID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ha ocurrido un error al obtener libro(s)")
		// TODO: Implementar manejo de errores como respuesta
		return
	}

	response := &BookResponse{
		Name:      result.Name,
		Author:    result.Author,
		CoverPage: result.CoverPage,
		Synopsis:  result.Synopsis,
		Price:     result.Price,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("ha ocurrido un error al enviar respuesta: ", err)
		// TODO: Implementar manejo de errores como respuesta
	}
}
