package framework

import "os"

//export IsExist
func IsExist(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}
