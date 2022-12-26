package main

import (
	"net/http"

	"github.com/treeforest/coredemo/framework"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		// 请求监听地址
		Addr: "localhost:8080",
	}
	server.ListenAndServe()
}
