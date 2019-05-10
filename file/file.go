package file

import (
	"log"
	"os"
)

func IsExist(filePath string) error {
	fileInfo, err := os.Stat(filePath)
	log.Println(fileInfo.Name(), fileInfo.Size(), fileInfo.ModTime(), fileInfo.IsDir(), fileInfo.Mode(), fileInfo.Sys())
	log.Println(err)
	return err
}
