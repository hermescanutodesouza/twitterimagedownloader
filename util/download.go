package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	downloadSize := int64(size)

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	if downloadSize == 55587 ||
		downloadSize == 52380 ||
		downloadSize == 55846 ||
		downloadSize == 38894 ||
		downloadSize == 44583 ||
		downloadSize == 44583 ||
		downloadSize == 40907 ||
		downloadSize == 34934 ||
		downloadSize == 32410 {
		os.Remove(filepath)
		return fmt.Errorf("File : %v removed", filepath)
	}

	return err
}
