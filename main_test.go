package main

import (
	"fmt"
	"testing"
)

func TestIsExist(t *testing.T) {
	path := "/Users/murphy/Desktop/test1.txt"
	f, err := IsExist(path)
	fmt.Println(f, err)

	path1 := "/Users/murphy/Desktop/test3.txt"
	f1, err1 := IsExist(path1)
	fmt.Println(f1, err1)
}
