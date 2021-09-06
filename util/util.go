package util

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var filesForCheck []string

type Screename struct {
	Screenname string `json:"screenname"`
}

func getFileHash(path string) string {
	fullPath, _ := os.Getwd()
	file := filepath.Join(fullPath, "files", "checkfiles", path)
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetHash() []string {
	return filesForCheck
}

func Checkfiles() {
	log.Println("Creating hash files to be ignored")
	fullPath, _ := os.Getwd()
	checkfiles := filepath.Join(fullPath, "files", "checkfiles")
	files, err := ioutil.ReadDir(checkfiles)
	if err != nil {
		log.Println("Error Checking files to be avoid")
		log.Println(err.Error())
	}
	log.Println("Found", len(files), "files")
	for _, f := range files {
		filesForCheck = append(filesForCheck, getFileHash(f.Name()))
	}
}

func CheckFolder(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		log.Println("Creating folder", folder)
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func CreateJsonFile(file string, content string) {
	f, err := os.Create(file)
	if err != nil {
		log.Println(err)
		return
	}
	l, err := f.WriteString(content)
	if err != nil {
		log.Println(err)
		err := f.Close()
		if err != nil {
			log.Fatalln(err)
			return
		}
		return
	}
	log.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		log.Println(err)
		return
	}
}

func ResolveFilePath(filename string) string {
	fullPath, _ := os.Getwd()
	return filepath.Join(fullPath, filename)
}

func CatchGeneralError(err *error) {
	if *err != nil {
		log.Fatalln(*err)
	}
}
