package storageImage

import (
	"net/http"

	"github.com/justinas/alice"

	secure "gopkg.in/unrolled/secure.v1"
)

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
