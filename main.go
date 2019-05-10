package main

import (
	"C"
	"fmt"
	"goLib/action"
	"goLib/file"
)

func main() {
	SayHello("123")
	b := []byte("def")
	SayHelloByte(b)
	SayBye()
	// file.IsExist("../test.txt")
	file.IsExist("/Users/murphy/Desktop/test.txt")
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
