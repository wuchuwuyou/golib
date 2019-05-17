package api

import (
	"log"
	"testing"
)

func TestLogin(t *testing.T) {
	err := Login("Xxxx", "xxx")
	if err != nil {
		log.Println(err)
	}
}

//params filePath, path, pathType, neid, from string, overwrite bool
func TestUploadFile(t *testing.T) {
	err := UploadFile("", "", "", "", "", true)
	if err != nil {
		log.Println(err)
	}
}
