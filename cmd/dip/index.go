package dip

import (
	"dip/internal"
	"net/http"
)

func StartDipHttpServer(tlsPort string) error {

	return http.ListenAndServe(":"+tlsPort, internal.DipHttpServer{})
}
