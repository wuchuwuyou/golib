package file

import (
	"testing"
	"fmt"
)

func TestIsExist(t *testing.T) {
	path := "/Users/murphy/Desktop/test2.txt"
	f, err := IsExist(path)
	fmt.Println(f, err)
}
func TestCreateFile(t *testing.T) {
	var err error
	path := "/Users/murphy/Desktop/test1.txt"
	// file := IsExist("/Users/murphy/Desktop/test.txt",&err)
	err = CreateFile(path)
	fmt.Println(err)
}