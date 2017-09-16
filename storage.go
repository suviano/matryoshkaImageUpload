package storageImage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
)

const googleCloudProject = "GOOGLE_CLOUD_PROJECT"

// IStorage gcloud storage interface
type IStorage interface {
	CreateClient(ctx context.Context) error
	SaveImg(ctx context.Context, reader io.Reader, bucket, name string, overwrite bool) (string, error)
}

// StorageClient bearer of cassandra driver
type StorageClient struct {
	client *storage.Client
}

// CreateClient connect to Bucket
func (storageCli *StorageClient) CreateClient(ctx context.Context) error {
	var err error
	projectID := os.Getenv(googleCloudProject)
	if projectID == "" {
		return fmt.Errorf("Environment variable %s must be set", googleCloudProject)
	}
	if ctx == nil {
		ctx = context.Background()
	}
	storageCli.client, err = storage.NewClient(ctx)
	return err
}

// SaveImg one image into gcloudstorage
func (storageCli *StorageClient) SaveImg(ctx context.Context, reader io.Reader, bucket, name string, overwrite bool) (string, error) {
	if storageCli.client == nil {
		return "", errors.New("instantiate storageClient")
	}
	defer storageCli.client.Close()

	bh := storageCli.client.Bucket(bucket)
	wc := bh.Object(name).NewWriter(ctx)
	if _, err := io.Copy(wc, reader); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	return "", nil
}
