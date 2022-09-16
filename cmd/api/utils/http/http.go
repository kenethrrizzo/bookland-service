package http

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"
)

func GenerateUniqueFileName(f *multipart.FileHeader) string {
	ext := filepath.Ext(f.Filename)

	return fmt.Sprintf("file-%d%s", time.Now().UnixNano(), ext)
}
