package handler

import "github.com/gin-gonic/gin"

func WebService() *gin.Engine {
	router := gin.Default()
	app := router.Group("/v1")
	appManager := new(AppManager)
	app.POST("/login", appManager.Login)
	app.POST("/send_code", appManager.SendCode)
	app.POST("/order", appManager.CreateOrder)
	app.GET("/orders", appManager.QueryOrders)
	app.GET("/order/:id", appManager.QueryOrder)
	app.PUT("/order/:id", appManager.UpdateOrder)
	app.DELETE("/order/:id", appManager.DeleteOrder)
	return router
}
