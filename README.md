# Golang 制作静态库

之前一直在考虑一个通用库的方法，也算是拓展自己的知识面吧。一直在看Go相关的东西，发现了一个go实现的[百度网盘](https://github.com/iikira/BaiduPCS-Go)，实现了跨平台的方案，想动手实现一下使用go打包静态、动态库，实现跨平台方案。
主要思路就是Golang源码，打包成静态库，供使用

# IDE
[VSCode](https://code.visualstudio.com/)

[GoLand](https://www.jetbrains.com/go/)

由于GFW的原因，我们要设置下代理
# 代理
以下是我的SS的设置，如果有不一样的参考 [VSCode官方文档](https://code.visualstudio.com/docs/setup/network)
## 设置VSCode 代理
` VSCode->首选项->设置->应用程序->代理服务器->Proxy`

    "http.proxy": "http://127.0.0.1:1087",

## 设置终端代理 
    export http_proxy=http://127.0.0.1:1087;
    export https_proxy=http://127.0.0.1:1087;
只当前窗口生效，关了就没了
# 包管理

Go Module 

module golib

### go.mod
go 1.12


require (
    github.com/spf13/viper v1.3.2  /// 外部package
    goLib/action v0.0.0-00010101000000-000000000000 // 本地package
)

replace goLib/action => ./action 

# 文件设置

> 函数可导设置 //export your_function_name
> 要导出的文件中 引入 import “C”
```golang
//export SayHello
func SayHello(name string) {
    fmt.Printf("func in Golang SayHello says: Hello, %s!\n", name)
}

```

# 静态库&动态库
我们需要在main包中导出C函数。对于C静态库构建方式来说，会忽略main包中的main函数，只是简单导出C函数。采用以下命令构建：

静态库
`$ go build -buildmode=c-archive -o number.a`
动态库
`$ go build -buildmode=c-shared -o number.so`

## 问题
目前还不支持 struct 结构体的导出 `头大啊`

# gomobile 
目前只支持Android和iOS
需要设置$GOPATH，并在$GOPATH的src中设置并打包 framework 
下载 gomobile 支持
`go get golang.org/x/mobile/cmd/gomobile`

在 $GOPATH/src 下创建自己要打包的package,然后执行

`gomobile bind -target=ios framework`

不同的framework里不能引入相同的文件，别问我怎么知道的

# 参考
- https://books.studygolang.com/advanced-go-programming-book/ch2-cgo/ch2-06-static-shared-lib.html
- https://www.kancloud.cn/cattong/go_command_tutorial/261347
- https://my.oschina.net/mickelfeng/blog/2252565
- http://blog.ralch.com/tutorial/golang-sharing-libraries/
- https://stackoverflow.com/questions/40573401/building-a-dll-with-go-1-7
- https://books.studygolang.com/advanced-go-programming-book/ch2-cgo/ch2-06-static-shared-lib.html
- https://github.com/golang/mobile
- https://godoc.org/golang.org/x/mobile/cmd/gomobile
