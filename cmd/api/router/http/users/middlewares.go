package users

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	domainErrors "github.com/kenethrrizzo/bookland-service/cmd/api/domain/errors"
)

var SECRET = []byte("texto-super-secreto") // Debe estar en variables de entorno

func ValidateJWT(c *gin.Context) {
	unauthorizedErr := domainErrors.NewAppErrorWithType(domainErrors.UnauthorizedError)

	if c.GetHeader("Authorization") == "" {
		c.AbortWithError(http.StatusUnauthorized, unauthorizedErr)
		return
	}

	tokenHeader := strings.Split(c.GetHeader("Authorization"), " ")[1]

	token, err := jwt.Parse(tokenHeader, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			c.AbortWithError(http.StatusUnauthorized, unauthorizedErr)
			return nil, unauthorizedErr
		}
		return SECRET, nil
	})
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, unauthorizedErr)
		return
	}

	if token.Valid {
		c.Next()
	}
}
