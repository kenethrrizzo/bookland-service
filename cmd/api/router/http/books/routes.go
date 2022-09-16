package books

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kenethrrizzo/bookland-service/cmd/api/domain/books"
	httpUtil "github.com/kenethrrizzo/bookland-service/cmd/api/utils/http"
)

type BookHandler struct {
	service books.BookService
}

func NewHandler(svc books.BookService) *BookHandler {
	return &BookHandler{svc}
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.service.GetAllBooks()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var booksResponse []BookResponse

	for _, r := range books {
		booksResponse = append(booksResponse, *bookDomaintoBookResponse(&r))
	}

	c.JSON(http.StatusOK, booksResponse)
}

func (h *BookHandler) GetBooksByGenre(c *gin.Context) {
	genre := c.Param("genre")
	
	books, err := h.service.GetBooksByGenre(genre)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var booksResponse []BookResponse

	for _, r := range books {
		booksResponse = append(booksResponse, *bookDomaintoBookResponse(&r))
	}

	c.JSON(http.StatusOK, booksResponse)
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("bookID"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book, err := h.service.GetBookByID(bookID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, bookDomaintoBookResponse(book))
}

func (h *BookHandler) RegisterNewBook(c *gin.Context) {
	var bookRequest BookRequest

	if err := c.Bind(&bookRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := bookRequestToBookDomain(&bookRequest)

	if bookRequest.Coverpage != nil {
		coverImgTmpRoute := fmt.Sprintf("./tmp/%s",
			httpUtil.GenerateUniqueFileName(bookRequest.Coverpage))

		err := c.SaveUploadedFile(bookRequest.Coverpage, coverImgTmpRoute)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		book.CoverPage = coverImgTmpRoute
	} else {
		book.CoverPage = ""
	}

	book, err := h.service.RegisterNewBook(book)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, bookDomaintoBookResponse(book))
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("bookID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	var bookRequest BookRequest

	if err := c.Bind(&bookRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := bookRequestToBookDomain(&bookRequest)
	if bookRequest.Coverpage != nil {
		coverImgTmpRoute := fmt.Sprintf("./tmp/%s",
			httpUtil.GenerateUniqueFileName(bookRequest.Coverpage))

		err := c.SaveUploadedFile(bookRequest.Coverpage, coverImgTmpRoute)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		book.CoverPage = coverImgTmpRoute
	} else {
		book.CoverPage = ""
	}

	book, err = h.service.UpdateBook(book, bookID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, bookDomaintoBookResponse(book))
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("bookID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = h.service.DeleteBook(bookID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &gin.H{
		"message": "deleted",
	})
}
