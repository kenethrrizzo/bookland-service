package http

import (
	"encoding/json"
	"net/http"

	domainErrors "github.com/kenethrrizzo/bookland-service/cmd/api/domain/errors"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func JSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		Error(w, err)
	}
}

func Error(w http.ResponseWriter, err error) {
	appErr, ok := err.(*domainErrors.AppError)

	if ok {
		errResponse := ErrorResponse{
			Message: appErr.Err.Error(),
		}

		switch appErr.Type {
		case domainErrors.NotFound:
			JSON(w, http.StatusNotFound, errResponse)
		default:
			JSON(w, http.StatusInternalServerError, errResponse)
		}
	} else {
		errResponse := ErrorResponse{
			Message: "Internal Server Error",
		}
		JSON(w, http.StatusInternalServerError, errResponse)
	}
}
