package storage

import (
	"context"

	"github.com/gobuffalo/buffalo/binding"
)

type manager interface {
	Upload(context.Context, binding.File) (string, string, error)
	UploadByte([]byte, string) (string, error)
	AccessURLFor(path string) (string, error)
	CreateVideo(path, filename string) error
}

func currentManager() manager {
	// if authKey != "{}" {
	// 	return remoteManager{}
	// }

	return remoteManager{}
}

func Upload(ctx context.Context, file binding.File) (string, string, error) {
	return currentManager().Upload(ctx, file)
}
