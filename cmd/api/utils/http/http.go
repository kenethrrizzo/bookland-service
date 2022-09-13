package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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

// TODO: Arreglar metodo para manejo de formularios
func SaveTempFile(w http.ResponseWriter, r *http.Request, formFile string) (*string, error) {
	pathToSave := "./tmp"

	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile(formFile)
	if err != nil {
		Error(w, err)
		return nil, err
	}
	defer file.Close()

	fileNameSplited := strings.Split(header.Filename, ".")
	fileExtension := fileNameSplited[len(fileNameSplited)-1]

	tempFile, err := os.CreateTemp(pathToSave, fmt.Sprintf("%s-*.%s", "tmp", fileExtension))
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

	filePath := fmt.Sprintf("./%s/%s", pathToSave, fileName)

	return &filePath, nil
}
