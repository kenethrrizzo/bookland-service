package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	domainErrors "github.com/kenethrrizzo/bookland-service/cmd/api/domain/errors"
	"github.com/spf13/viper"
)

// Poner en paquete para usar en comun
func ErrJSON(err error) *Response {
	return &Response{
		Status: "ERROR",
		Data:   gin.H{"error": err.Error()},
	}
}

func ValidateJWT(c *gin.Context) {
	unauthorizedErr := domainErrors.NewAppErrorWithType(domainErrors.UnauthorizedError)

	if c.GetHeader("Authorization") == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrJSON(unauthorizedErr))
		return
	}

	// TODO: Validar uso de Bearer

	tokenHeader := strings.Split(c.GetHeader("Authorization"), " ")[1]

	token, err := jwt.Parse(tokenHeader, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrJSON(unauthorizedErr))
			return nil, unauthorizedErr
		}
		return []byte(viper.GetString("secret")), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrJSON(unauthorizedErr))
		return
	}

	if token.Valid {
		c.Next()
	}
}
