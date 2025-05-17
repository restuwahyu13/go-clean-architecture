package pkg

import (
	"crypto/rand"
	"crypto/tls"
	"net/http"
	"os"

	"github.com/ory/graceful"
	cons "github.com/restuwahyu13/go-clean-architecture/shared/constants"
	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
	helper "github.com/restuwahyu13/go-clean-architecture/shared/helpers"
)

func Graceful(Handler func() *dto.GracefulConfig) error {
	parser := helper.NewParser()
	inboundSize, _ := parser.ToInt(os.Getenv("INBOUND_SIZE"))

	h := Handler()
	secure := true

	if _, ok := os.LookupEnv("GO_ENV"); ok && os.Getenv("GO_ENV") != "development" {
		secure = false
	}

	server := http.Server{
		Handler:        h.HANDLER,
		Addr:           ":" + h.ENV.APP.PORT,
		MaxHeaderBytes: inboundSize,
		TLSConfig: &tls.Config{
			Rand:               rand.Reader,
			InsecureSkipVerify: secure,
		},
	}

	Logrus(cons.INFO, "Server listening on port %s", h.ENV.APP.PORT)
	return graceful.Graceful(server.ListenAndServe, server.Shutdown)
}
