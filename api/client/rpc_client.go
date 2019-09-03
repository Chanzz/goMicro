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
func CreateDingDianTui(orderInfo appProto.DingDianTuiInfo) *appProto.Response {
	// 声明要找哪个 RPC 服务，以及绑定到哪个 client 数据结构上
	// apiService 做发起 rpc 请求的一方，所以是 client
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : CreateOrder")
	// 实际调用 rpc
	resp, err := appService.CreateDingDianTui(context.TODO(), &orderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryDingDianTuis(orderId appProto.DingDianTuiInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrders")
	resp, err := appService.QueryDingDianTuis(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryDingDianTui(orderId appProto.DingDianTuiInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrder")
	resp, err := appService.QueryDingDianTui(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func UpdateDingDianTui(updateOrderInfo appProto.DingDianTuiInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : UpdateOrder")
	resp, err := appService.UpdateDingDianTui(context.TODO(), &updateOrderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func DeleteDingDianTui(orderId appProto.DingDianTuiInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : DeleteOrder")
	resp, err := appService.DeleteDingDianTui(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func CreateChaoShiXin(orderInfo appProto.ChaoShiXinInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : CreateOrder")
	// 实际调用 rpc
	resp, err := appService.CreateChaoShiXin(context.TODO(), &orderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryChaoShiXins(orderId appProto.ChaoShiXinInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrders")
	resp, err := appService.QueryChaoShiXins(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryChaoShiXin(orderId appProto.ChaoShiXinInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrder")
	resp, err := appService.QueryChaoShiXin(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func UpdateChaoShiXin(updateOrderInfo appProto.ChaoShiXinInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : UpdateOrder")
	resp, err := appService.UpdateChaoShiXin(context.TODO(), &updateOrderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func DeleteChaoShiXin(orderId appProto.ChaoShiXinInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : DeleteOrder")
	resp, err := appService.DeleteChaoShiXin(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func CreateAIYunHu(orderInfo appProto.AIYunHuInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : CreateOrder")
	// 实际调用 rpc
	resp, err := appService.CreateAIYunHu(context.TODO(), &orderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryAIYunHus(orderId appProto.AIYunHuInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrders")
	resp, err := appService.QueryAIYunHus(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryAIYunHu(orderId appProto.AIYunHuInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrder")
	resp, err := appService.QueryAIYunHu(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func UpdateAIYunHu(updateOrderInfo appProto.AIYunHuInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : UpdateOrder")
	resp, err := appService.UpdateAIYunHu(context.TODO(), &updateOrderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func DeleteAIYunHu(orderId appProto.AIYunHuInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : DeleteOrder")
	resp, err := appService.DeleteAIYunHu(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func CreateKeLiuJing(orderInfo appProto.KeLiuJingInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : CreateOrder")
	resp, err := appService.CreateKeLiuJing(context.TODO(), &orderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryKeLiuJings(orderId appProto.KeLiuJingInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrders")
	resp, err := appService.QueryKeLiuJings(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryKeLiuJing(orderId appProto.KeLiuJingInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrder")
	resp, err := appService.QueryKeLiuJing(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func UpdateKeLiuJing(updateOrderInfo appProto.KeLiuJingInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : UpdateOrder")
	resp, err := appService.UpdateKeLiuJing(context.TODO(), &updateOrderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func DeleteKeLiuJing(orderId appProto.KeLiuJingInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : DeleteOrder")
	resp, err := appService.DeleteKeLiuJing(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func CreatePinXiaoTong(orderInfo appProto.PinXiaoTongInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : CreateOrder")
	resp, err := appService.CreatePinXiaoTong(context.TODO(), &orderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryPinXiaoTongs(orderId appProto.PinXiaoTongInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrders")
	resp, err := appService.QueryPinXiaoTongs(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryPinXiaoTong(orderId appProto.PinXiaoTongInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrder")
	resp, err := appService.QueryPinXiaoTong(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func UpdatePinXiaoTong(updateOrderInfo appProto.PinXiaoTongInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : UpdateOrder")
	resp, err := appService.UpdatePinXiaoTong(context.TODO(), &updateOrderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func DeletePinXiaoTong(orderId appProto.PinXiaoTongInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : DeleteOrder")
	resp, err := appService.DeletePinXiaoTong(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func CreateCheShiChuang(orderInfo appProto.CheShiChuangInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : CreateOrder")
	resp, err := appService.CreateCheShiChuang(context.TODO(), &orderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryCheShiChuangs(orderId appProto.CheShiChuangInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrders")
	resp, err := appService.QueryCheShiChuangs(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryCheShiChuang(orderId appProto.CheShiChuangInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrder")
	resp, err := appService.QueryCheShiChuang(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func UpdateCheShiChuang(updateOrderInfo appProto.CheShiChuangInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : UpdateOrder")
	resp, err := appService.UpdateCheShiChuang(context.TODO(), &updateOrderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func DeleteCheShiChuang(orderId appProto.CheShiChuangInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : DeleteOrder")
	resp, err := appService.DeleteCheShiChuang(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func CreateYunHuaXiang(orderInfo appProto.YunHuaXiangInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : CreateOrder")
	resp, err := appService.CreateYunHuaXiang(context.TODO(), &orderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryYunHuaXiangs(orderId appProto.YunHuaXiangInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrders")
	resp, err := appService.QueryYunHuaXiangs(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryYunHuaXiang(orderId appProto.YunHuaXiangInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrder")
	resp, err := appService.QueryYunHuaXiang(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func UpdateYunHuaXiang(updateOrderInfo appProto.YunHuaXiangInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : UpdateOrder")
	resp, err := appService.UpdateYunHuaXiang(context.TODO(), &updateOrderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func DeleteYunHuaXiang(orderId appProto.YunHuaXiangInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : DeleteOrder")
	resp, err := appService.DeleteYunHuaXiang(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func CreateWaQianKe(orderInfo appProto.WaQianKeInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : CreateOrder")
	resp, err := appService.CreateWaQianKe(context.TODO(), &orderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryWaQianKes(orderId appProto.WaQianKeInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrders")
	resp, err := appService.QueryWaQianKes(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func QueryWaQianKe(orderId appProto.WaQianKeInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : QueryOrder")
	resp, err := appService.QueryWaQianKe(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func UpdateWaQianKe(updateOrderInfo appProto.WaQianKeInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : UpdateOrder")
	resp, err := appService.UpdateWaQianKe(context.TODO(), &updateOrderInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func DeleteWaQianKe(orderId appProto.WaQianKeInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : DeleteOrder")
	resp, err := appService.DeleteWaQianKe(context.TODO(), &orderId)
	if err != nil {
		println(err.Error())
	}
	return resp
}

func Register(registerInfo appProto.RegisterInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : Register")
	resp, err := appService.Register(context.TODO(), &registerInfo)
	if err != nil {
		println(err.Error())
	}
	return resp
}
func UserInfo(userInfo appProto.UserInfo) *appProto.Response {
	appService := appProto.NewOrderService(APP_SERVICE_NAME, apiService.Client())
	fmt.Printf("Create RPC client for : UserInfo")
	resp, err := appService.User(context.TODO(), &userInfo)
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
