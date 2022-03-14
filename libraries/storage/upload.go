package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gobuffalo/buffalo/binding"
	"github.com/gofrs/uuid"
)

type remoteManager struct{}

// UploadByte implements manager
func (r remoteManager) UploadByte(file []byte, name string) (string, error) {
	panic("unimplemented")
}

func (r remoteManager) Upload(ctx context.Context, file binding.File) (string, error) {
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(file.Filename)
	path := fmt.Sprintf("%v%v", uuid.Must(uuid.NewV4()), ext)
	path = strings.ToLower(path)

	pathFolder := fmt.Sprintf("./public/%v", path)
	new, err := os.Create(pathFolder)
	if err != nil {
		log.Fatal(err)
	}
	defer new.Close()

	if _, err := io.Copy(new, file.File); err != nil {
		return "", err
	}

	resp, err := cld.Upload.Upload(ctx, pathFolder, uploader.UploadParams{})
	if err != nil {
		return "", err
	}

	if err := os.Remove(pathFolder); err != nil {
		return "", err
	}

	log.Println(resp.SecureURL)

	return path, nil
}

func (r remoteManager) AccessURLFor(path string) (string, error) {
	return "", nil
}

func (r remoteManager) CreateVideo(name, filename string) error {
	return nil
}
