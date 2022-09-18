package users

import (
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

}

func (h *UserHandler) Login(c *gin.Context) {
	
}