package storageImage

import (
	"context"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"

	log "github.com/sirupsen/logrus"
)

const googleCloudProject = "GOOGLE_CLOUD_PROJECT"

// IStorage gcloud storage interface
type IStorage interface {
	CreateClient(ctx context.Context) error
	SaveImg(ctx context.Context, prefix, bucket string, bufMap bufMedia) error
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

// SaveImg save one image into gcloud storage
func (storageCli *StorageClient) SaveImg(ctx context.Context, prefix, bucket string, bufMap bufMedia) error {
	if storageCli.client == nil {
		if err := storageCli.CreateClient(ctx); err != nil {
			log.Warningf("%s Error creating client", prefix)
			return err
		}
	}
	defer storageCli.client.Close()

	object := storageCli.client.Bucket(bucket).Object(fmt.Sprintf("%s/%s", prefix, bufMap.Path))
	wc := object.NewWriter(ctx)
	wc.ContentType = bufMap.MimeTyp
	if _, err := io.Copy(wc, bufMap.Buf); err != nil {
		log.Warningf("{StorageClient}{SaveImg} Error while coping file", prefix)
		return err
	}
	if err := wc.Close(); err != nil {
		log.Warningf("{StorageClient}{SaveImg} Error closing file", prefix)
		return err
	}

	return nil
}
