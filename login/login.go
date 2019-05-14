package login

import (
	"C"
	"log"
)

func Login(username, password string) error {
	log.Println(username)
	log.Println(password)
	return nil
}
