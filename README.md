# goStudy

## 学习文档
http://docscn.studygolang.com/doc/code.html
## go指南
https://tour.go-zh.org/


## 你可以离线使用本指南。
离线使用时所有代码样例均在你的机器上编译和运行，所以速度会更快。
要离线使用 Go 指南，请先在本地安装 Go 然后使用 go get 命令安装 gotour-zh:
go get github.com/Go-zh/tour/gotour
然后运行得到的 gotour 程序就可以了：
cd $GOPATH/bin
./gotour


## 启动本地教程报错
Couldn't find tour files: could not find go-tour content; check $GOROOT and $GOPATH
先把GOPATH设置为自己go工程的路径 export GOPATH=/Users/xiafei/code/goStudy/
然后依次执行
cd $GOPATH/bin
./gotour


last
http://127.0.0.1:3999/moretypes/15