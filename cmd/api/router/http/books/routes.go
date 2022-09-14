package books

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kenethrrizzo/bookland-service/cmd/api/domain/books"
	httpUtil "github.com/kenethrrizzo/bookland-service/cmd/api/utils/http"
	"github.com/sirupsen/logrus"
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
		httpUtil.ERROR(w, err)
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
		httpUtil.ERROR(w, err)
		return
	}

	result, err := handl.service.GetBookByID(bookID)
	if err != nil {
		httpUtil.ERROR(w, err)
		return
	}

	response := bookDomaintoBookResponse(result)

	httpUtil.JSON(w, http.StatusOK, response)
}

func (handl *BookHandler) RegisterNewBook(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, fHeader, err := r.FormFile("coverpage")
	if err != nil {
		logrus.Error(err)
		httpUtil.ERROR(w, err)
		return
	}
	file.Close()

	coverImgRoute, err := httpUtil.SaveFormFileToTempFolder(w, file, fHeader)
	if err != nil {
		logrus.Error(err)
		httpUtil.ERROR(w, err)
		return
	}

	book, err := bookFormToBookDomain(w, r)
	if err != nil {
		logrus.Error(err)
		httpUtil.ERROR(w, err)
		return
	}

	bookDomain, err := handl.service.RegisterNewBook(book, *coverImgRoute)
	if err != nil {
		logrus.Error(err)
		httpUtil.ERROR(w, err)
		return
	}

	response := bookDomaintoBookResponse(bookDomain)

	httpUtil.JSON(w, http.StatusCreated, response)
}

func (handl *BookHandler) UpdateBookCoverImage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, fHeader, err := r.FormFile("coverpage")
	if err != nil {
		logrus.Error(err)
		httpUtil.ERROR(w, err)
		return
	}
	file.Close()

	coverImgRoute, err := httpUtil.SaveFormFileToTempFolder(w, file, fHeader)
	if err != nil {
		logrus.Error(err)
		httpUtil.ERROR(w, err)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		logrus.Error(err)
		httpUtil.ERROR(w, err)
		return
	}

	bookDomain, err := handl.service.UpdateBookCoverImage(id, *coverImgRoute)
	if err != nil {
		httpUtil.ERROR(w, err)
		return
	}

	response := bookDomaintoBookResponse(bookDomain)

	httpUtil.JSON(w, http.StatusOK, response)
}

func (handl *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookIDstr := chi.URLParam(r, "bookID")

	bookID, err := strconv.Atoi(bookIDstr)
	if err != nil {
		httpUtil.ERROR(w, err)
		return
	}

	err = handl.service.DeleteBook(bookID)
	if err != nil {
		httpUtil.ERROR(w, err)
		return
	}

	response := httpUtil.MessageResponse{
		Message: "deleted!",
	}

	httpUtil.JSON(w, http.StatusOK, response)
}
