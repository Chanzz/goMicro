package main

import (
	"fmt"
	"github.com/micro/go-micro/web"
	"goMicro/api/client"
	"goMicro/api/handler"
	"time"
)

var (
	listenPort = "0.0.0.0:18667" // Restful 接口的 IP 和端口
)

func main() {
	serviceName := client.API_SERVICE_NAME

	fmt.Printf("start service:%s version:%s", serviceName)

	// consul 一般是 service.NewService，web 是一种特别的 service
	service := web.NewService(
		web.Name(serviceName),
		web.Address(listenPort), // web 服务的地址
		web.RegisterTTL(time.Second*10),
		web.RegisterInterval(time.Second*5)) // consul 的 Options，包括 consul 的地址
	_ = service.Init()

	// 初始化控制层
	gin := handler.WebService() // 拿到一个 gin Engine

	// 初始化 micro service
	client.InitApiService()

	// 绑定 gin 到 "/" 路径上处理所有请求
	service.Handle("/", gin)

	// 启动 service
	if err := service.Run(); err != nil {
		println(err) // 出错了
	}
}
