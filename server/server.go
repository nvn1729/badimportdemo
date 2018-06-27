//Adapted from https://github.com/jcbsmpsn/golang-https-example
package main

import (
	"crypto/tls"
	"net/http"
	"os"
)

func main() {
	srv := &http.Server{
		Addr: "localhost:8443",
		Handler: &handler{},
		TLSConfig: &tls.Config{},
	}
	srv.ListenAndServeTLS(os.Args[1], os.Args[2])
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hi, I'm using " + os.Args[1] + "\n"))
}
