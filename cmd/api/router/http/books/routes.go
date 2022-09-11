package books

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kenethrrizzo/bookland-service/cmd/api/domain/books"
)

type BookHandler struct {
	service books.BookService
}

func NewHandler(svc books.BookService) *BookHandler {
	return &BookHandler{svc}
}

func (handl *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	results, err := handl.service.GetAllBooks()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ha ocurrido un error al obtener los libros")
		// TODO: Implementar manejo de errores como respuesta
		return
	}

	var response []BookResponse

	for _, result := range results {
		response = append(response, BookResponse{
			Name:      result.Name,
			Author:    result.Author,
			CoverPage: result.CoverPage,
			Synopsis:  result.Synopsis,
			Price:     result.Price,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("ha ocurrido un error al enviar respuesta: ", err)
		// TODO: Implementar manejo de errores como respuesta
	}
}

func (handl *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	bookIDstr := chi.URLParam(r, "bookID")

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
		log.Println("ha ocurrido un error al obtener el libro con id: " + bookIDstr)
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
