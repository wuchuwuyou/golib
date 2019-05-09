package action

import "fmt"

func Start(host string) string {
	fmt.Println("host:", host)
	return "Hello,world!"
}
