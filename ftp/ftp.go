package ftp

import (
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

type FTPConnectionDetails struct {
	server   string
	port     string
	username string
	password string
}

func (f *FTPConnectionDetails) UploadExportToFTP(filename string) {
	c, err := ftp.Dial(f.server+f.port, ftp.DialWithTimeout(15*time.Second))
	if err != nil {
		log.Fatalf("Failed to dial FTP server %s", err)
	}
	if err := c.Login(f.username, f.password); err != nil {
		log.Fatalf("Failed to login %s", err)
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file %s %s", filename, err)
	}
	if err := c.Stor(filename, file); err != nil {
		log.Fatalf("FTP file creation error %s %s", filename, err)
	}
	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
