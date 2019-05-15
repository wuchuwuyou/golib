package api

import (
	"C"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)
import (
	"errors"
	"net/http"
	"os"
	"strconv"
)

var (
	sessionId string
)

//export Login
//params username string,password string
func Login(username, password string) error {
	fmt.Println(username)
	fmt.Println(password)
	url := "https://console.box.lenovo.com/v2/user/login_new"
	args := &fasthttp.Args{}
	args.Add("loginkey", username)
	args.Add("password", password)
	args.Add("loginType", "0")
	args.Add("uuid", "1AC4890C-505A-5567-9BDD-73AEB25B00E7")
	args.Add("device_name", "Murphy's MacBook Pro (2)")

	/*
		body
			"device_name" = "Murphy\U7684MacBook Pro (2)";
			loginType = 0;
			loginkey = "wangteng@lenovocloud.com";
			password = 123456;
			uuid = "1AC4890C-505A-5567-9BDD-73AEB25B00E7";
		header
			"Content-Type" = "application/x-www-form-urlencoded";
			Cookie = "language=en;X-LENOVO-SESS-ID=xxxx";
			"User-Agent" = "LDClientMac_5.0.10.2_MacOSX_10.14.4";

	*/

	client := &fasthttp.Client{}
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.SetConnectionClose()
	request.SetRequestURI(url)
	args.WriteTo(request.BodyWriter())
	request.Header.SetMethod("POST")
	request.Header.SetContentType("application/x-www-form-urlencoded")
	request.Header.SetUserAgent("LDClientMac_5.0.10.2_MacOSX_10.14.4")
	request.Header.SetCookie("language", "zh-cn")
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
	fmt.Printf(filePath, path, pathType, neid, from, overwrite)
	// http: //disk.yutong.com:10081/v2/files/databox/公交新能源产品部/01 电机控制模块/01 模块内部共享文档/15 安全质量-冯勇敢邵彭真/2 FEMA数据库（QM单、生产、市场质量问题）/QM单问题汇总/电机控制模块7月10日-19日QM单质量问题汇总.xlsx?X-LENOVO-SESS-ID=7bo5mj66dame2fcbaer9ju8ju4&path_type=ent&from=&neid=244881708&rev=1b94a030624cae83&action=&op_type="
	url := "https://console.box.lenovo.com/v2/files/databox/"
	url += path
	url += "?"
	url += "X-LENOVO-SESS-ID="
	url += sessionId
	url += "&pathType"
	url += pathType
	if len(neid) > 0 {
		url += "&neid"
		url += neid
	}
	if len(from) > 0 {
		url += "&from"
		url += from
	}
	url += "&overwrite"
	url += strconv.FormatBool(overwrite)

	doUpload(url, filePath)
	// client := &fasthttp.Client{}
	// request := fasthttp.AcquireRequest()
	// response := fasthttp.AcquireResponse()

	// defer fasthttp.ReleaseRequest(request)
	// defer fasthttp.ReleaseResponse(response)

	// request.SetConnectionClose()
	// request.SetRequestURI(url)
	// request.Header.SetMethod("POST")
	// request.Header.SetContentType("application/x-www-form-urlencoded")
	// request.Header.SetUserAgent("LDClientMac_5.0.10.2_MacOSX_10.14.4")
	// request.Header.SetCookie("language", "zh-cn")
	// if err := client.Do(request, response); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(string(response.Header.Header()))
	// 	fmt.Println(string(response.Body()))
	// 	b := response.Body()
	// 	var v map[string]string
	// 	json.Unmarshal(b, &v)
	// 	sessionId = v["X-LENOVO-SESS-ID"]
	// 	fmt.Println(v)
	// 	fmt.Println(sessionId)
	// }
	return nil
}
func doUpload(url, filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res, err := http.Post(url, "application/x-www-form-urlencoded", file) // 第二个参数用来指定 "Content-Type"
	// resp, err := http.Post(Url, "image/jpeg", file)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
}
