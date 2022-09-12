package files

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	domainErrors "github.com/kenethrrizzo/bookland-service/cmd/api/domain/errors"
)

type Store struct {
	s3client *s3.Client
}

func NewStore(s3client *s3.Client) *Store {
	return &Store{s3client}
}

func (s *Store) UploadFile(bucketName string, filePath string) (*string, error) {
	fileName := filepath.Base(filePath)

	upFile, err := os.Open(filePath)
	if err != nil {
		appErr := domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		return nil, appErr
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	fileSize := upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	result, err := s.s3client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   upFile,
	})
	if err != nil {
		appErr := domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		return nil, appErr
	}

	log.Println(result)

	return nil, nil
}
