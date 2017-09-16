package storageImage

import (
	"context"

	"cloud.google.com/go/storage"
	log "github.com/sirupsen/logrus"
)

// IStorage gcloud storage interface
type IStorage interface {
	CreateClient(projectID string) error
	SaveMultiple(bucket, objects []string) ([]string, error)
	Save(bucket, object string) (string, error)
}

// StorageClient bearer of cassandra driver
type StorageClient struct {
	client    *storage.Client
	projectID string
}

// CreateClient connect to Bucket
func (storageCli *StorageClient) CreateClient(projectID string) error {
	log.Warningln("{New client being created}")
	var err error
	ctx := context.Background()
	storageCli.projectID = projectID
	storageCli.client, err = storage.NewClient(ctx)
	return err
}

// SaveMultiple images into gcloudstorage
func (storageCli *StorageClient) SaveMultiple(bucket, objects []string) ([]string, error) {
	return nil, nil
}

// Save one image into gcloudstorage
func (storageCli *StorageClient) Save(bucket, object string) (string, error) {
	return "", nil
}
