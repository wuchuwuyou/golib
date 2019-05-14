package main

import (
	"C"
	_"fmt"
	"os"
	"io/ioutil"
	"log"
)


func main() {
	// file.IsExist("../test.txt")
	// file.IsExist("/Users/murphy/Desktop/test.txt")
}
//export Login
func Login(username, password string) error {
	log.Println(username)
	log.Println(password)
	return nil
}

//export Entry
type Entry interface {
	path() string
	pathType() string
	rev() string
	neid() string
	preNeid() string
}
//export MetaData
func MetaData(path,pathType string) (Entry,error) {
	var entry Entry
	log.Println(entry)
	return entry,nil
}
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
		return false,err
	}
	return true,nil
}
type FileInfo = os.FileInfo

//export ReadDir
func ReadDir(filePath string) ([]FileInfo, error) {
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