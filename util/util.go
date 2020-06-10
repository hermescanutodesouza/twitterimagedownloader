package util

import (
	"log"
	"os"
	"path/filepath"
)

type Screename struct {
	Screenname string `json:"screenname"`
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


