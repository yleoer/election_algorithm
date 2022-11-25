package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"election_algorithm/pkg/global"
	"election_algorithm/pkg/router"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	cfg := global.GetConfig()
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.SbiPort()),
		Handler:      h2c.NewHandler(router.NewRouter(), new(http2.Server)),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	log.Fatal(server.ListenAndServe())
}
