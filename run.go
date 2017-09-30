package storageImage

import (
	"flag"
)

var (
	certFile, keyFile, serverAddr, bucket, projectID string
	storageClient                                    IStorage
)

func init() {
	flag.StringVar(&bucket, "bucket", "", "bucket name")
	flag.Parse()

	storageClient = &StorageClient{}
}
