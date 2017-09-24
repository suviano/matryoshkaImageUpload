package storageImage

import (
	"flag"
	"log"
)

var (
	certFile, keyFile, serverAddr, bucket, projectID string
	storageClient                                    IStorage
)

func init() {
	flag.StringVar(&serverAddr, "serverAddr", ":8080", "server address")
	flag.StringVar(&bucket, "bucket", "", "bucket name")
	flag.Parse()
}

// Run start storage client and server
func Run() {
	storageClient = &StorageClient{}
	log.Printf("{ initializing server on port %s }\n", serverAddr)
	log.Fatal(server.ListenAndServe())
}
