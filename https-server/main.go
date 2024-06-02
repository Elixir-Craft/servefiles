package main

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
	Size  string
}

// DirectoryListing holds the data for the HTML template
type DirectoryListing struct {
	CurrentPath string
	ParentPath  string
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

	// Check if the path is a directory or a file
	fileInfo, err := os.Stat(path)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	if fileInfo.IsDir() {
		// If it's a directory, list its contents
		files, err := os.ReadDir(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// check if index.html exists in the directory
		for _, file := range files {
			if file.Name() == "index.html" {
				http.ServeFile(w, req, path+"/index.html")
				return
			}
		}

		var fileInfos []FileInfo
		for _, file := range files {
			fileInfo, _ := file.Info()
			fileInfos = append(fileInfos, FileInfo{Name: file.Name(), IsDir: file.IsDir(), Size: fileSize(fileInfo.Size())})
		}

		data := DirectoryListing{
			CurrentPath: req.URL.Path,
			ParentPath:  filepath.Dir(req.URL.Path),
			Files:       fileInfos,
		}

		err = templates.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		// If it's a file, serve it directly
		http.ServeFile(w, req, path)
	}
}

// isImage checks if a file is an image based on its extension
func isImage(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp":
		return true
	}
	return false
}

// file size conversion
func fileSize(size int64) string {
	const (
		_  = iota
		kb = 1 << (10 * iota)
		mb
		gb
		tb
		pb
		eb
		zb
		yb
	)

	switch {
	case size < kb:
		return fmt.Sprintf("%d bytes", size)
	case size < mb:
		return fmt.Sprintf("%.2f KB", float64(size)/float64(kb))
	case size < gb:
		return fmt.Sprintf("%.2f MB", float64(size)/float64(mb))
	case size < tb:
		return fmt.Sprintf("%.2f GB", float64(size)/float64(gb))
	case size < pb:
		return fmt.Sprintf("%.2f TB", float64(size)/float64(tb))
	default:
		return fmt.Sprintf("%.2f PB", float64(size)/float64(pb))

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
