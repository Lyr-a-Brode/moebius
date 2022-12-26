package services

import (
	"context"
	"fmt"
	"github.com/mholt/archiver"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

type BookService struct{}

func NewBookService() BookService {
	return BookService{}
}

func (s BookService) ProcessBook(ctx context.Context, src io.Reader, size int64) error {
	r, err := createArchiveReader(src, size)
	if err != nil {
		return err
	}
	defer func(r archiver.Reader) {
		err := r.Close()
		if err != nil {
			log.WithContext(ctx).WithError(err).Warn("Unable to close archive reader")
		}
	}(r)

	i := 0
	for {
		f, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		joined := filepath.Join("/tmp", strconv.Itoa(i))
		out, err := os.Create(joined)
		if err != nil {
			return fmt.Errorf("%s: creating new file: %v", joined, err)
		}
		defer out.Close()

		_, err = io.Copy(out, f)
		if err != nil {
			return fmt.Errorf("%s: writing file: %v", joined, err)
		}

		i++
	}

	return nil
}

func createArchiveReader(src io.Reader, size int64) (r archiver.Reader, err error) {
	rar := &archiver.Rar{}

	err = rar.Open(src, size)
	if err == nil {
		return rar, nil
	}

	zip := &archiver.Zip{}

	err = zip.Open(src, size)
	if err != nil {
		return nil, err
	}

	return zip, nil
}
