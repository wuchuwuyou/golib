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

文件见设置 //export 名称
# 手动静态库

## .a静态库

/// 编译成 .a 和头文件
go build -buildmode=c-archive -o main.a


https://books.studygolang.com/advanced-go-programming-book/ch2-cgo/ch2-06-static-shared-lib.html

https://www.kancloud.cn/cattong/go_command_tutorial/261347


https://my.oschina.net/mickelfeng/blog/2252565
*/

# gomobile 

go get golang.org/x/mobile/cmd/gomobile

https://github.com/golang/mobile