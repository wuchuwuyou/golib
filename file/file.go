package file

// type FileError struct {
// }

// func (fe *FileError) Error() string {
// 	return "文件错误"
// }

// func IsExist(filePath string) (os.FileInfo, error) {
// 	fileInfo, err := os.Stat(filePath)
// 	log.Println(fileInfo.Name(), fileInfo.Size(), fileInfo.ModTime(), fileInfo.IsDir(), fileInfo.Mode(), fileInfo.Sys())
// 	log.Println(err)
// 	return fileInfo, err
// }
// func ReadDir(filePath string) ([]os.FileInfo, error) {
// 	files, err := ioutil.ReadDir(filePath)
// 	return files, err
// }
// func CreateFile(filePath string) (*os.File, error) {
// 	newFile, err := os.Create("test.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(newFile)
// 	newFile.Close()
// 	return newFile, err
// }
