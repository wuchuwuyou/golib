package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/valyala/fasthttp"
)

var (
	sessionId string
)

//export Entry
type Entry interface {
	path() string
	pathType() string
	rev() string
	neid() string
	preNeid() string
}

//export MetaData
/**
 * @note 获取文件信息
 * @params path,pathType string
 * @param path 文件路径
 * @param pathType 文件空间
 * @return Entry,error
 */
func MetaData(path, pathType string) (string, error) {

	log.Printf(path, pathType)
	return "", nil
}

//export Login
//params username string,password string
func Login(username, password string) error {
	fmt.Println(username)
	fmt.Println(password)
	url := "hurl"
	args := &fasthttp.Args{}
	args.Add("loginkey", username)
	args.Add("password", password)
	args.Add("loginType", "0")
	args.Add("uuid", "ss")
	args.Add("device_name", "xx")

	client := &fasthttp.Client{}
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.SetConnectionClose()
	request.SetRequestURI(url)
	args.WriteTo(request.BodyWriter())
	request.Header.SetMethod("POST")
	request.Header.SetContentType("xxxx")
	request.Header.SetUserAgent("xxxx")
	request.Header.SetCookie("xxx", "xx")
	if err := client.Do(request, response); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(response.Header.Header()))
		fmt.Println(string(response.Body()))
		b := response.Body()
		var v map[string]string
		json.Unmarshal(b, &v)
		sessionId = v["X-LENOVO-SESS-ID"]
		fmt.Println(v)
		fmt.Println(sessionId)
	}
	return nil
}

//export UploadFile
//params filePath, path, pathType, neid, from string, overwrite bool
func UploadFile(filePath, path, pathType, neid, from string, overwrite bool) error {
	if !(len(sessionId) > 0) {
		var err error = errors.New("请先登录")
		return err
	}
	fmt.Printf(filePath, path, pathType, neid, from, overwrite, "\n")

	doUpload(url, filePath)
	return nil
}
func doUpload(url, filepath string) {
	fmt.Printf("url:%s\npath:%s", url, filepath)
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res, err := http.Post(url, "application/x-www-form-urlencoded", file) // 第二个参数用来指定 "Content-Type"
	// resp, err := http.Post(Url, "image/jpeg", file)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("\nBody:%s", res.Body)
	fmt.Printf("\nHeader:%s", res.Header)
	defer res.Body.Close()
}
