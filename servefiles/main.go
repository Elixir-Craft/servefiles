package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/fatih/color"

	"github.com/Elixir-Craft/servefiles/localip"

	"github.com/Elixir-Craft/servefiles/certgen"

	"github.com/Elixir-Craft/servefiles/webtemplates"
)

var (
	CertFilePath  = certgen.GetConfigDir() + "/cert/cert.pem"
	KeyFilePath   = certgen.GetConfigDir() + "/cert/key.pem"
	HomeTemplate  = template.Must(template.New("home.html").Parse(webtemplates.Home))
	IndexTemplate = template.Must(template.New("index.html").Parse(webtemplates.Index))
	AuthTemplate  = template.Must(template.New("").Parse(webtemplates.Auth))
)

type FileInfo struct {
	Name  string
	IsDir bool
	Size  string
}

// flags
var (
	port            = flag.String("p", "4443", "port to listen on")
	regenerateCerts = flag.Bool("r", false, "regenerate certificates")
	password        = flag.String("P", "", "password to access the files")
)

// DirectoryListing holds the data for the HTML template
type DirectoryListing struct {
	CurrentPath string
	ParentPath  string
	Files       []FileInfo
}

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {

	// show terminal output
	fmt.Println("Received request from", req.RemoteAddr)

	err := HomeTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func fileHandler(w http.ResponseWriter, req *http.Request) {
	// Strip the "/files" prefix from the request URL
	path := "." + req.URL.Path[len("/files"):]
	// Decode percent-encoded characters
	path, err := url.PathUnescape(path)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

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

		err = IndexTemplate.ExecuteTemplate(w, "index.html", data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		// If it's a file, serve it directly
		http.ServeFile(w, req, path)
	}
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

func passwordProtected(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check for a session or a simple cookie to verify if the password was entered correctly
		if cookie, err := r.Cookie("password"); err != nil || !checkPassword(cookie.Value) {
			if r.Method == "POST" && checkPassword(r.FormValue("password")) {
				// Set a simple cookie for demonstration purposes (not secure for production)
				http.SetCookie(w, &http.Cookie{
					Name:   "password",
					Value:  r.FormValue("password"),
					Path:   "/",
					MaxAge: 3600, // Expires after 3600 seconds
				})
				http.Redirect(w, r, r.URL.Path, http.StatusFound)
				return
			}
			servePasswordPrompt(w, r)
			return
		}
		next(w, r)

	}
}

func servePasswordPrompt(w http.ResponseWriter, r *http.Request) {

	// read the html file

	err := AuthTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func checkPassword(enteredPassword string) bool {
	fmt.Println("Password check")
	return enteredPassword == *password
}

func main() {

	flag.Parse()

	// check if the certificates need to be regenerated
	// already exists or not provided flag
	if *regenerateCerts || !certgen.CertFilesExist(CertFilePath, KeyFilePath) {
		certgen.Certsetup()
	}
	// certgen.Certsetup()
	// load tls certificates
	serverTLSCert, err := tls.LoadX509KeyPair(CertFilePath, KeyFilePath)
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverTLSCert},
	}

	mux := http.NewServeMux()
	// mux.HandleFunc("/files/", fileHandler)
	// mux.HandleFunc("/files/", passwordProtected(fileHandler))

	// if password is provided then protect the files
	if *password != "" {
		mux.HandleFunc("/files/", passwordProtected(fileHandler))
	} else {
		mux.HandleFunc("/files/", fileHandler)
	}

	mux.HandleFunc("/", httpRequestHandler)

	server := http.Server{
		Addr:      ":" + *port,
		Handler:   mux,
		TLSConfig: tlsConfig,
	}

	// start the server show terminal output
	// fmt.Println("Server is running on https://localhost:4443")
	// all ip addresses on the local system
	ips, err := localip.Get()
	if err != nil {
		log.Fatalf("Error getting local IP addresses: %v", err)
	}

	fmt.Println("Server is running on:")
	for _, ip := range ips {
		// fmt.Printf(" https://%s:%s\n", ip, *port)
		color.Green("https://%s:%s\n", ip, *port)
	}
	fmt.Printf("Press Ctrl+C to stop the server\n\n")

	defer server.Close()
	log.Fatal(server.ListenAndServeTLS("", ""))
}
