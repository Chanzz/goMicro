package client

import (
	"context"
	"fmt"
	goMicro "github.com/micro/go-micro"
	appProto "goMicro/service/proto"
)

var apiService goMicro.Service // 一个全局变量保存初始化后的实例

var API_SERVICE_NAME = "goMicro.srv.api"
var APP_SERVICE_NAME = "goMicro.srv.app"

// 类似于静态函数，返回一个实例值。这里的单例与否需要上层来保证
// 这里是创建 api 服务，也就是 api web service。
func InitApiService() {
	apiService = goMicro.NewService(
		goMicro.Name(API_SERVICE_NAME),
		goMicro.Version("latest"))
	apiService.Init()
}

// 将本地调用，转换成 rpc 调用。如果数据结构不一致，这里会需要做数据结构的转换，包括传给 rpc 调用的参数，以及收到的返回结果。
func CreateOrder(orderInfo appProto.OrderInfo) *appProto.Response {
	// 声明要找哪个 RPC 服务，以及绑定到哪个 client 数据结构上
	// apiService 做发起 rpc 请求的一方，所以是 client
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : CreateOrder")
	// 实际调用 rpc
	resp, err := appService.CreateOrder(context.TODO(), &orderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryOrders(orderId appProto.OrderInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrders")
	resp, err := appService.QueryOrders(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryOrder(orderId appProto.OrderInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrder")
	resp, err := appService.QueryOrder(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func UpdateOrder(updateOrderInfo appProto.OrderInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : UpdateOrder")
	resp, err := appService.UpdateOrder(context.TODO(), &updateOrderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func DeleteOrder(orderId appProto.OrderInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : DeleteOrder")
	resp, err := appService.DeleteOrder(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func Login(loginInfo appProto.LoginInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : Login")
	resp, err := appService.Login(context.TODO(), &loginInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func SendCode(loginInfo appProto.LoginInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : SendCode")
	resp, err := appService.SendCode(context.TODO(), &loginInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
