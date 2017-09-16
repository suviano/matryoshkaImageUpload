package storageImage

import (
	"flag"
	"log"
)

var (
	certFile, keyFile, serverAddr, bucket, projectID string
	localStorage                                     IStorage
)

func init() {
	flag.StringVar(&certFile, "certFile", "server.crt", "certificate file to https")
	flag.StringVar(&keyFile, "keyFile", "server.key", "key file to https")
	flag.StringVar(&serverAddr, "serverAddr", ":8080", "server address")
	flag.StringVar(&projectID, "projectID", "you-project", "define default gcloud project")
	flag.StringVar(&bucket, "bucket", "you-bucket", "bucket name")
	flag.Parse()
}

func main() {
	localStorage = &StorageClient{}
	err := localStorage.CreateClient(projectID)
	if err != nil {
		log.Fatalf("{creating client returned {%v}}", err)
	}
	serve()
}
