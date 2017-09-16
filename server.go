package storageImage

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/justinas/alice"
	secure "gopkg.in/unrolled/secure.v1"

	"log"

	"golang.org/x/net/http2"
)

var (
	server    *http.Server
)

func init() {
	handler := http.NewServeMux()
	handler.Handle("/favicon.ico", http.NotFoundHandler())
	handler.Handle("/image", imageHandle())

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

	if err := http2.ConfigureServer(server, nil); err != nil {
		log.Fatalf("%v", err)
	}
}

func serve() {
	log.Printf("{ initializing server on port %s }\n", serverAddr)
	log.Fatal(server.ListenAndServeTLS(certFile, keyFile))
}

func applyMiddlewares(fn http.Handler) http.Handler {
	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:            []string{},
		HostsProxyHeaders:       []string{},
		SSLRedirect:             false,
		SSLTemporaryRedirect:    false,
		SSLHost:                 "",
		SSLProxyHeaders:         map[string]string{},
		STSSeconds:              0,
		STSIncludeSubdomains:    false,
		STSPreload:              false,
		ForceSTSHeader:          false,
		FrameDeny:               false,
		CustomFrameOptionsValue: "",
		ContentTypeNosniff:      false,
		BrowserXssFilter:        false,
		ContentSecurityPolicy:   "",
		PublicKey:               "",
		ReferrerPolicy:          "",
		IsDevelopment:           true,
	})
	return alice.New(secureMiddleware.Handler).Then(fn)
}
