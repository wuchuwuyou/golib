package main

import (
	"fmt"
	"testing"
)

func TestIsExist(t *testing.T) {
	path := "/Users/murphy/Desktop/test2.txt"
	f, err := IsExist(path)
	fmt.Println(f, err)
}
func TestCreateFile(t *testing.T) {
	var err LDError
	path := "/Users/murphy/Desktop/test1.txt"
	// file := IsExist("/Users/murphy/Desktop/test.txt",&err)
	err = CreateFile(path)
	fmt.Println(err.err)
	fmt.Println(err.code)
}
