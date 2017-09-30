package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"github.com/suviano/matryoshkaImageUpload"
)

var (
	prefix, bucket string
)

func main() {
	flag.StringVar(&prefix, "prefix", "test", "prefix name")
	flag.StringVar(&bucket, "bucket", "", "bucket name")
	flag.Parse()

	filePath := "cassie-boca-296277.jpg"

	b, err := ioutil.ReadFile(fmt.Sprintf("sample_image/%s", filePath))
	if err != nil {
		log.Warning("Error reading file")
		panic(err)
	}
	err = matryoshka.WriteImage(prefix, filePath, bucket, bytes.NewBuffer(b))
	if err != nil {
		log.Fatalf("Error uploading image -> %v", err)
	}
}
