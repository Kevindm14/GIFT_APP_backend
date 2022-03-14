package storage

import (
	"context"
	"mime/multipart"
)

type localManager struct{}

// UploadByte implements manager
func (l localManager) UploadByte(file []byte, name string) (string, error) {
	panic("unimplemented")
}

func (l localManager) Upload(ctx context.Context, file *multipart.FileHeader) (string, error) {
	return "", nil
}

func (l localManager) AccessURLFor(path string) (string, error) {
	return "", nil
}

func (l localManager) CreateVideo(name, filename string) error {
	return nil
}
