package files

import (
	"context"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	domainErrors "github.com/kenethrrizzo/bookland-service/cmd/api/domain/errors"
)

const (
	BUCKETNAME = "bookland-bucket"
)

type Store struct {
	s3client *s3.Client
}

func NewStore(s3client *s3.Client) *Store {
	return &Store{s3client}
}

func (s *Store) UploadFile(filePath string) (*string, error) {
	uploader := manager.NewUploader(s.s3client)

	file, openErr := os.Open(filePath)
	if openErr != nil {
		appErr := domainErrors.NewAppError(openErr, domainErrors.UnknownError)
		return nil, appErr
	}
	defer file.Close()

	fileName := filepath.Base(filePath)

	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(BUCKETNAME),
		Key:    aws.String(fileName),
		ACL:    types.ObjectCannedACLPublicRead,
	})
	if uploadErr != nil {
		appErr := domainErrors.NewAppError(uploadErr, domainErrors.UnknownError)
		return nil, appErr
	}

	return &result.Location, nil
}
