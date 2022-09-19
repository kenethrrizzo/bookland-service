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

func (h *UserHandler) Register(c *gin.Context) {
	var userRequest UserRequest

	if err := c.Bind(&userRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userDomain, err := userRequestToUserDomain(&userRequest)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	auth, err := h.service.Register(userDomain)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, authDomainToUserResponse(auth))
}

func (h *UserHandler) Login(c *gin.Context) {
	var userRequest UserRequest

	if err := c.Bind(&userRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userDomain, err := userRequestToUserDomain(&userRequest)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	auth, err := h.service.Login(userDomain)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, authDomainToUserResponse(auth))
}
