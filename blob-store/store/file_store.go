package store

import (
	"context"
	"github.com/Lyr-a-Brode/moebius/blob-store/trace"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
)

type FileStore struct {
	path string
}

func NewFileStore(path string) FileStore {
	return FileStore{path: path}
}

func (f FileStore) Put(ctx context.Context, blobID string, blob io.Reader) error {
	dst, err := os.Create(path.Join(f.path, blobID))
	if err != nil {
		return err
	}

	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			log.WithError(err).WithField(trace.LogFieldName, trace.FromContext(ctx)).
				Error("Unable to close file")
		}
	}(dst)

	if _, err = io.Copy(dst, blob); err != nil {
		return err
	}

	return nil
}
