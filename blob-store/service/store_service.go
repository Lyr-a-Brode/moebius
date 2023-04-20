package service

import (
	"context"
	"github.com/google/uuid"
	"io"
)

type Store interface {
	Put(ctx context.Context, blobID string, blob io.Reader, format string) error
}

type StoreService struct {
	store Store
}

func NewStoreService(store Store) StoreService {
	return StoreService{
		store: store,
	}
}

func (s StoreService) StoreBlob(ctx context.Context, blob io.Reader, format string) (blobID string, err error) {
	blobID = uuid.New().String()

	if err := s.store.Put(ctx, blobID, blob, format); err != nil {
		return "", err
	}

	return blobID, nil
}
