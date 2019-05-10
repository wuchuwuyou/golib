# Golang 制作静态库

之前一直在考虑一个通用库的方法，也算是拓展自己的知识面吧

一只在看Go相关的东西，今天就动手实现一下使用go打包静态、动态库，实现跨平台方案




/*


设置VSCode 代理

"http.proxy": "http://127.0.0.1:1087",
"http.proxyStrictSSL": false,
"http.proxyAuthorization":null



设置终端代理 export http_proxy=http://127.0.0.1:1087;export https_proxy=http://127.0.0.1:1087;

module golib

go 1.12


require (
github.com/spf13/viper v1.3.2  /// 外部package
goLib/action v0.0.0-00010101000000-000000000000 // 本地package
)

replace goLib/action => ./action 


/// 编译成 .a 和头文件
go build -buildmode=c-archive -o main.a


https://books.studygolang.com/advanced-go-programming-book/ch2-cgo/ch2-06-static-shared-lib.html

https://www.kancloud.cn/cattong/go_command_tutorial/261347


https://my.oschina.net/mickelfeng/blog/2252565
*/
