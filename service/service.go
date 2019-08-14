package main

import (
	"github.com/micro/go-micro"
	"goMicro/service/handler"
	proto "goMicro/service/proto"
	"time"
)

func main() {
	serviceName := "goMicro.srv.app"

	db := handler.CreateMySqlConnection()
	defer db.Close()

	appDao := &handler.AppDao{DB: db} // 将数据源传给 service 对象
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5))

	// Init & Binding handling
	service.Init()
	_ = proto.RegisterOrderHandler(service.Server(), appDao)

	// Run Service
	if err := service.Run(); err != nil {
		println(err)
	}
}
