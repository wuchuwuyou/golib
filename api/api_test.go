package api

import (
	"testing"
	"log"
)
func TestLogin(t *testing.T) {
	err := Login("wangteng@lenovocloud.com","123456")
	if err != nil {
		log.Println(err)
	}
}
//params filePath, path, pathType, neid, from string, overwrite bool
func TestUploadFile(t *testing.T) {
	err := UploadFile("/Users/murphy/Desktop/test1.txt","123/test1.txt","self","","",true)
	if err != nil {
		log.Println(err)
	}
}
