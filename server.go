package storageImage

import (
	"net/http"
	"time"
)

var (
	server *http.Server
)

func init() {
	handler := http.NewServeMux()
	handler.Handle("/image/", imageHandle())

	server = &http.Server{
		Addr:         serverAddr,
		Handler:      handler,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}
}
