package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	domainErrors "github.com/kenethrrizzo/bookland-service/cmd/api/domain/errors"
	"github.com/sirupsen/logrus"
)

type MessageResponse struct {
	Message string `json:"message"`
}

func JSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		ERROR(w, err)
	}
}

func ERROR(w http.ResponseWriter, err error) {
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

func SaveFormFileToTempFolder(w http.ResponseWriter, file multipart.File, fHeader *multipart.FileHeader) (*string, error) {
	tmpFolder := "./tmp/"

	if _, err := os.Stat(tmpFolder); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(tmpFolder, os.ModePerm); err != nil {
			logrus.Error(err)
			return nil, err
		}
	}

	fileName := strings.TrimSpace(fHeader.Filename)
	fileType := strings.ToLower(fileName[len(fileName)-3:])

	// TODO: Guardar nombre de archivo como un hash generado
	fileName = fmt.Sprintf("%d%d%d%d%d%d%d%d%d.%s",
		time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(),
		time.Now().Minute(), time.Now().Second(), time.Now().UnixMilli(),
		time.Now().UnixMicro(), time.Now().Nanosecond(), fileType)

	if fileType != "png" && fileType != "jpg" {
		return nil, domainErrors.NewAppError(errors.New("invalid file type"), domainErrors.UnknownError)
	}

	fileRoute := tmpFolder + fileName

	newFile, err := os.Create(fileRoute)
	if err != nil {
		return nil, err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, file)
	if err != nil {
		return nil, err
	}

	return &fileRoute, nil
}
