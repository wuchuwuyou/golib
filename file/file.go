package file

import (
	"io/ioutil"
	"log"
	"os"
)

//export FileExist
func FileExist(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return err
	}
	return nil
}

//export IsExist
func IsExist(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

//export ReadDir
func ReadDir(filePath string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(filePath)
	return files, err
}

//export CreateFile
func CreateFile(filePath string) error {
	newFile, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println(newFile)
	newFile.Close()
	return nil
}
