package files

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	domainErrors "github.com/kenethrrizzo/bookland-service/cmd/api/domain/errors"
)

type Store struct {
	s3client *s3.Client
}

func NewStore(s3client *s3.Client) *Store {
	return &Store{s3client}
}

func (s *Store) UploadFile(bucketName string, filePath string) (*string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.UnknownError)
		return nil, appErr
	}
	defer file.Close()

	fileName := filepath.Base(filePath)

	_, err = s.s3client.PutObject(context.TODO(), &s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		ACL:    types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		appErr := domainErrors.NewAppError(err, domainErrors.UnknownError)
		return nil, appErr
	}

	fileURL := s.GetFileURL(bucketName, fileName)

	log.Println(fileURL)

	return &fileURL, nil
}

func (s *Store) GetFileURL(bucketName string, filePath string) string {
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, filePath)

	return url
}
