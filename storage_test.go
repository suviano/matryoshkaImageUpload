package matryoshka

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupCreateClient(t *testing.T) *StorageClient {
	t.Helper()
	return &StorageClient{}
}

func TestStorageClient_CreateClient(t *testing.T) {
	storageClient := setupCreateClient(t)

	t.Run("MissingEnviromentVariable", func(t *testing.T) {
		err := os.Setenv(googleCloudProject, "")
		assert.Nil(t, err)

		err = storageClient.CreateClient(nil)
		assert.NotNil(t, err)
		assert.EqualError(t, fmt.Errorf("Environment variable %s must be set", googleCloudProject), err.Error())
	})

	t.Run("UsualRunWithUndefinedContext", func(t *testing.T) {
		err := os.Setenv(googleCloudProject, googleCloudProject)
		assert.Nil(t, err)

		err = storageClient.CreateClient(nil)
		assert.Nil(t, err)
	})

	t.Run("UsualRun", func(t *testing.T) {
		err := os.Setenv(googleCloudProject, googleCloudProject)
		assert.Nil(t, err)

		ctx := context.Background()
		err = storageClient.CreateClient(ctx)
		assert.Nil(t, err)
	})
}
