package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	domainErrors "github.com/kenethrrizzo/bookland-service/cmd/api/domain/errors"
)

type MessageResponse struct {
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
		errResponse := MessageResponse{
			Message: appErr.Err.Error(),
		}

		switch appErr.Type {
		case domainErrors.NotFound:
			JSON(w, http.StatusNotFound, errResponse)
		default:
			JSON(w, http.StatusInternalServerError, errResponse)
		}
	} else {
		errResponse := MessageResponse{
			Message: "Internal Server Error",
		}
		JSON(w, http.StatusInternalServerError, errResponse)
	}
}

// TODO: Implementar uso de s3 para carga de imagenes
func SaveFile(w http.ResponseWriter, r *http.Request, formFile string, pathToSave string) (*string, error) {
	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile(formFile)
	if err != nil {
		Error(w, err)
		return nil, err
	}
	defer file.Close()

	fileNameSplited := strings.Split(header.Filename, ".")
	fileExtension := fileNameSplited[len(fileNameSplited)-1]

	tempFile, err := os.CreateTemp(pathToSave, fmt.Sprintf("%s-*.%s", time.Now(), fileExtension))
	if err != nil {
		Error(w, err)
		return nil, err
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		Error(w, err)
		return nil, err
	}
	tempFile.Write(fileBytes)

	fileName := tempFile.Name()

	return &fileName, nil
}
