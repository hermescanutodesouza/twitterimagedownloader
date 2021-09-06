package util

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func filesToSave(hash string) bool {
	nofile := GetHash()
	for _, v := range nofile {
		if v == hash {
			return false
		}
	}
	return true
}

func createHash(r []byte) string {
	h := sha256.New()
	if _, err := io.Copy(h, bytes.NewReader(r)); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func DownloadFile(filepath string, url, screenname string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	byteData, _ := ioutil.ReadAll(resp.Body)

	hash := createHash(byteData)

	if !filesToSave(hash) {
		return errors.New(fmt.Sprintf("Ignored %s %s %s", url, screenname, hash))
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	return err
}
