package main

import (
	"C"
	"fmt"
	"goLib/action"
	"os"
	_"io/ioutil"
	"log"
)

func main() {
	SayHello("123")
	b := []byte("def")
	SayHelloByte(b)
	SayBye()
	// file.IsExist("../test.txt")
	// file.IsExist("/Users/murphy/Desktop/test.txt")
}

type LDError struct {
	err error
	code int
}


//export SayHello
func SayHello(name string) {
	action.Start("https://www.google.com")
	fmt.Printf("func in Golang SayHello says: Hello, %s!\n", name)
}

//export SayHelloByte
func SayHelloByte(name []byte) {
	fmt.Printf("func in Golang SayHelloByte says: Hello, %s!\n", string(name))
}

//export SayBye
func SayBye() {
	fmt.Println("func in Golang SayBye says: Bye!")
}

//export UserName
func UserName(name string)(string,error) {
	var n = name
	fmt.Println("func in Golang SayBye says: Bye!")
	err := fmt.Errorf("%s Not The King", n)
	return n,err
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
// // read dir
// func LDReadDir(filePath string) ([]os.FileInfo, error) {
// 	files, err := ioutil.ReadDir(filePath)
// 	return files, err
// }

/**
	create file
*/
func CreateFile(filePath string) LDError {
	newFile, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
	var e LDError
	if err != nil {
		e.err = err
		e.code = -101
	}
	return e
}