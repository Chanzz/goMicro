package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "goMicro/service/proto"
	"log"
)

func ProcessEvent(ctx context.Context, event *proto.OrderInfo) error {
	fmt.Printf("Got event %+v\n", event)
	fmt.Printf("Got event %+v\n", event.Location)
	return nil
}
func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.pubsub"),
	)
	service.Init()
	micro.RegisterSubscriber("test", service.Server(), ProcessEvent)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
