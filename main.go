package main

import (
	"C"
	_"fmt"
	"os"
	"golib/file"
	"golib/api"
)
//export FileExist
func FileExist(filePath string) error {
	return file.FileExist(filePath)
}

//export IsExist
func IsExist(filePath string) (bool, error) {
	return file.IsExist(filePath)
}

type FileInfo = os.FileInfo

//export ReadDir
func ReadDir(filePath string) ([]FileInfo, error) {
	return file.ReadDir(filePath)
}

//export CreateFile
func CreateFile(filePath string) error {
	return file.CreateFile(filePath)
}

//export MetaData
/**
 * @note 获取文件信息
 * @params path,pathType string
 * @param path 文件路径
 * @param pathType 文件空间
 * @return Entry,error
 */
 func MetaData(path, pathType string) (string, error) {
	return api.MetaData(path,pathType)
}

//export Login
//params username string,password string
func Login(username, password string) error {
	return api.Login(username,password)
}

//export UploadFile
//params filePath, path, pathType, neid, from string, overwrite bool
func UploadFile(filePath, path, pathType, neid, from string, overwrite bool) error {
	return api.UploadFile(filePath,path,pathType,neid,from,overwrite)
}

func main() {
	// file.IsExist("../test.txt")
	// file.IsExist("/Users/murphy/Desktop/test.txt")
}