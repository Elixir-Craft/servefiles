package certgen

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"gopkg.in/yaml.v2"

	"crypto/x509/pkix"

	"github.com/Elixir-Craft/servefiles/localip"
)

// file location by os
func getConfigFilePath() string {
	var configDir string
	var configFileName string

	switch OS := runtime.GOOS; OS {
	case "windows":
		// On Windows, use the APPDATA directory
		configDir = filepath.Join(os.Getenv("APPDATA"), "ServeIt")
		configFileName = "config.yaml"
	case "darwin":
		// On macOS, use the home directory
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		configDir = filepath.Join(home, "Library", "Application Support", "ServeIt")
		configFileName = "config.yaml"
	default:
		// On Unix-like systems, use the home directory
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		configDir = filepath.Join(home, ".ServeIt")
		configFileName = "config.yaml"
	}

	// Create configuration directory if it doesn't exist
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, 0755)
	}

	return filepath.Join(configDir, configFileName)
}

type CertConfig struct {
	Organization  string `yaml:"organization"`
	Country       string `yaml:"country"`
	Province      string `yaml:"province"`
	Locality      string `yaml:"locality"`
	StreetAddress string `yaml:"street_address"`
	PostalCode    string `yaml:"postal_code"`
}

type Config struct {
	Cert CertConfig `yaml:"cert"`
}

func loadConfig() (Config, error) {
	config := Config{
		Cert: CertConfig{
			Organization:  "ServeIt Inc",
			Country:       "LK",
			Province:      "",
			Locality:      "Colombo",
			StreetAddress: "Main Street",
			PostalCode:    "00000",
		},
	}
	if _, err := os.Stat(getConfigFilePath()); err == nil {
		file, err := os.ReadFile(getConfigFilePath())
		if err != nil {
			return config, err
		}
		err = yaml.Unmarshal(file, &config)
		if err != nil {
			return config, err
		}
	}
	fmt.Println(config)
	return config, nil
}

func Certsetup() (serverTLSConf *tls.Config, clientTLSConf *tls.Config, err error) {
	config, err := loadConfig()
	if err != nil {
		return nil, nil, err
	}

	// set up our CA certificate
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization:  []string{config.Cert.Organization},
			Country:       []string{config.Cert.Country},
			Province:      []string{config.Cert.Province},
			Locality:      []string{config.Cert.Locality},
			StreetAddress: []string{config.Cert.StreetAddress},
			PostalCode:    []string{config.Cert.PostalCode},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	// generate a private key for the CA
	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}

	// create CA
	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return nil, nil, err
	}

	// pem encode
	caPEM := new(bytes.Buffer)
	pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})

	// pem encode the private key
	caPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(caPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivKey),
	})

	localIPs, err := localip.Get()
	if err != nil {
		return nil, nil, err
	}

	cert := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization:  []string{config.Cert.Organization},
			Country:       []string{config.Cert.Country},
			Province:      []string{config.Cert.Province},
			Locality:      []string{config.Cert.Locality},
			StreetAddress: []string{config.Cert.StreetAddress},
			PostalCode:    []string{config.Cert.PostalCode},
		},
		IPAddresses:  append([]net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback}, localIPs...),
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	certPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, cert, ca, &certPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return nil, nil, err
	}

	certPEM := new(bytes.Buffer)
	pem.Encode(certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})

	certPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(certPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(certPrivKey),
	})

	serverCert, err := tls.X509KeyPair(certPEM.Bytes(), certPrivKeyPEM.Bytes())
	if err != nil {
		return nil, nil, err
	}

	serverTLSConf = &tls.Config{
		Certificates: []tls.Certificate{serverCert},
	}

	certpool := x509.NewCertPool()
	certpool.AppendCertsFromPEM(caPEM.Bytes())
	clientTLSConf = &tls.Config{
		RootCAs: certpool,
	}

	// create directory if it does not exist
	os.Mkdir("cert", 0755)

	os.WriteFile("cert/cert.pem", certPEM.Bytes(), 0644)
	os.WriteFile("cert/key.pem", certPrivKeyPEM.Bytes(), 0644)

	return
}

func CertFilesExist(certFilePath, keyFilePath string) bool {
	_, err := os.Stat(certFilePath)
	if os.IsNotExist(err) {
		return false
	}

	_, err = os.Stat(keyFilePath)

	return err == nil

}
