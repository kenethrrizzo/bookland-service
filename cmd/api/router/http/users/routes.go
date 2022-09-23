package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userDomain "github.com/kenethrrizzo/bookland-service/cmd/api/domain/users"
)

type UserHandler struct {
	service userDomain.UserService
}

func NewHandler(service userDomain.UserService) *UserHandler {
	return &UserHandler{service}
}

// Poner en paquete para usar en comun
func ErrJSON(err error) *Response {
	return &Response{
		Status: "ERROR",
		Result: gin.H{"error": err.Error()},
	}
}

func OkJSON(data interface{}) *Response {
	return &Response{
		Status: "OK",
		Result: data,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var userRequest UserRequest

	if err := c.Bind(&userRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrJSON(err))
		return
	}

	userDomain, err := userRequestToUserDomain(&userRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrJSON(err))
		return
	}

	auth, err := h.service.Register(userDomain)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrJSON(err))
		return
	}

	userResponse := authDomainToUserResponse(auth)

	c.JSON(http.StatusCreated, OkJSON(userResponse))
}

func (h *UserHandler) Login(c *gin.Context) {
	var userRequest UserRequest

	if err := c.Bind(&userRequest); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrJSON(err))
		return
	}

	userDomain, err := userRequestToUserDomain(&userRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrJSON(err))
		return
	}

	auth, err := h.service.Login(userDomain)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrJSON(err))
		return
	}

	userResponse := authDomainToUserResponse(auth)

	c.JSON(http.StatusCreated, OkJSON(userResponse))
}
