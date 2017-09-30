package storageImage

import (
	"flag"
)

var (
	bucket        string
	storageClient IStorage
)

func init() {
	flag.StringVar(&bucket, "bucket", "", "bucket name")
	flag.Parse()

	storageClient = &StorageClient{}
}
