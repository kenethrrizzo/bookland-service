package books

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kenethrrizzo/bookland-service/cmd/api/domain/books"
	httpUtil "github.com/kenethrrizzo/bookland-service/cmd/api/utils/http"
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
		httpUtil.Error(w, err)
		return
	}

	var response []BookResponse

	for _, result := range results {
		response = append(response, *bookDomaintoBookResponse(&result))
	}

	httpUtil.JSON(w, http.StatusOK, response)
}

func (handl *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	bookIDstr := chi.URLParam(r, "bookID")

	bookID, err := strconv.Atoi(bookIDstr)
	if err != nil {
		httpUtil.Error(w, err)
		return
	}

	result, err := handl.service.GetBookByID(bookID)
	if err != nil {
		httpUtil.Error(w, err)
		return
	}

	response := bookDomaintoBookResponse(result)

	httpUtil.JSON(w, http.StatusOK, response)
}

func (handl *BookHandler) RegisterNewBook(w http.ResponseWriter, r *http.Request) {
	var request BookRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		httpUtil.Error(w, err)
		return
	}

	book := bookRequestToBookDomain(&request)

	bookDomain, err := handl.service.RegisterNewBook(book)
	if err != nil {
		httpUtil.Error(w, err)
		return
	}

	response := bookDomaintoBookResponse(bookDomain)

	httpUtil.JSON(w, http.StatusCreated, response)
}

// TODO: Aplicar manejo de formularios
func (handl *BookHandler) UpdateBookCoverImage(w http.ResponseWriter, r *http.Request) {
	var request BookRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		httpUtil.Error(w, err)
		return
	}

	bookDomain, err := handl.service.UpdateBookCoverImage(request.Id, request.Coverpage)
	if err != nil {
		httpUtil.Error(w, err)
		return
	}

	response := bookDomaintoBookResponse(bookDomain)

	httpUtil.JSON(w, http.StatusOK, response)
}

func (handl *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookIDstr := chi.URLParam(r, "bookID")

	bookID, err := strconv.Atoi(bookIDstr)
	if err != nil {
		httpUtil.Error(w, err)
		return
	}

	err = handl.service.DeleteBook(bookID)
	if err != nil {
		httpUtil.Error(w, err)
		return
	}

	response := httpUtil.MessageResponse{
		Message: "deleted!",
	}

	httpUtil.JSON(w, http.StatusOK, response)
}
