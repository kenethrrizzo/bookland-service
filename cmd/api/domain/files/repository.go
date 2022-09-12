package files

type FileRepository interface {
	UploadFile(string, string) (*string, error)
}