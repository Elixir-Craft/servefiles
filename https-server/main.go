package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/Elixir-Craft/https-server/certgen"
)

var (
	CertFilePath = "cert/cert.pem"
	KeyFilePath  = "cert/key.pem"
)

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {

	// show terminal output
	fmt.Println("Received request from", req.RemoteAddr)

	w.Write([]byte("Hello,World!\n"))
}
func main() {

	certgen.Certsetup()
	// load tls certificates
	serverTLSCert, err := tls.LoadX509KeyPair(CertFilePath, KeyFilePath)
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverTLSCert},
	}
	server := http.Server{
		Addr:      ":4443",
		Handler:   http.HandlerFunc(httpRequestHandler),
		TLSConfig: tlsConfig,
	}

	// start the server show terminal output
	fmt.Println("Server is running on https://localhost:4443")

	defer server.Close()
	log.Fatal(server.ListenAndServeTLS("", ""))
}
