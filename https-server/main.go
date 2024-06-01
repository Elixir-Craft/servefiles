package main

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/Elixir-Craft/https-server/certgen"
)

var (
	CertFilePath = "cert/cert.pem"
	KeyFilePath  = "cert/key.pem"
	templates    = template.Must(template.ParseFiles("templates/index.html"))
)

type FileInfo struct {
	Name  string
	IsDir bool
}

// DirectoryListing holds the data for the HTML template
type DirectoryListing struct {
	CurrentPath string
	Files       []FileInfo
}

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {

	// show terminal output
	fmt.Println("Received request from", req.RemoteAddr)

	w.Write([]byte("Hello,World!\n"))
}

func fileHandler(w http.ResponseWriter, req *http.Request) {
	// Strip the "/files" prefix from the request URL
	path := "." + req.URL.Path[len("/files"):]

	files, err := os.ReadDir(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var fileInfos []FileInfo
	for _, file := range files {
		fileInfos = append(fileInfos, FileInfo{Name: file.Name(), IsDir: file.IsDir()})
	}

	data := DirectoryListing{
		CurrentPath: req.URL.Path,

		Files: fileInfos,
	}

	err = templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	mux := http.NewServeMux()
	mux.HandleFunc("/files/", fileHandler)
	mux.HandleFunc("/", httpRequestHandler)

	server := http.Server{
		Addr:      ":4443",
		Handler:   mux,
		TLSConfig: tlsConfig,
	}

	// start the server show terminal output
	fmt.Println("Server is running on https://localhost:4443")

	defer server.Close()
	log.Fatal(server.ListenAndServeTLS("", ""))
}
