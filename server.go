package storageImage

import (
	"crypto/tls"
	"net/http"
	"time"

	"log"

	"golang.org/x/net/http2"
)

var (
	server *http.Server
)

func init() {
	handler := http.NewServeMux()
	handler.Handle("/favicon.ico", http.NotFoundHandler())
	handler.Handle("/image/", imageHandle())

	server = &http.Server{
		Addr:         serverAddr,
		Handler:      handler,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
		},
	}

	err := http2.ConfigureServer(server, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
