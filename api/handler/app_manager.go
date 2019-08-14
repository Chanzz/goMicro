package handler

import (
	"github.com/gin-gonic/gin"
	"goMicro/api/client"
	appProto "goMicro/service/proto"
	"log"
)

type AppManager struct{}

func (a *AppManager) CreateOrder(c *gin.Context) {
	var orderInfo appProto.OrderInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.CreateOrder(orderInfo)
	c.JSON(200, resp)
}

func (a *AppManager) QueryOrders(c *gin.Context) {
	var orderInfo appProto.OrderInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.QueryOrders(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) QueryOrder(c *gin.Context) {
	var orderInfo appProto.OrderInfo
	orderInfo.Id = c.Param("id")
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.QueryOrder(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) UpdateOrder(c *gin.Context) {
	//var updateOrderInfo appProto.UpdateOrderInfo
	var orderInfo appProto.OrderInfo
	orderInfo.Id = c.Param("id")
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	//updateOrderInfo.OrderId = new(appProto.OrderId)
	//updateOrderInfo.OrderId.OrderId=c.Param("id")
	resp := client.UpdateOrder(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) DeleteOrder(c *gin.Context) {
	var orderInfo appProto.OrderInfo
	orderInfo.Id = c.Param("id")
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.DeleteOrder(orderInfo)
	c.JSON(200, resp)
}
