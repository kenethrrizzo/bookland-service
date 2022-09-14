package files

type FileRepository interface {
	UploadFile(string) (*string, error)
}