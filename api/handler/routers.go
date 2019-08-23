package handler

import "github.com/gin-gonic/gin"

func WebService() *gin.Engine {
	router := gin.Default()
	dingdiantui := router.Group("/dingdiantui")
	appManager := new(AppManager)
	dingdiantui.POST("/login", appManager.Login)
	dingdiantui.POST("/send_code", appManager.SendCode)
	dingdiantui.POST("/order", appManager.CreateDingDianTui)
	dingdiantui.GET("/orders", appManager.QueryDingDianTuis)
	dingdiantui.GET("/order/:id", appManager.QueryDingDianTui)
	dingdiantui.PUT("/order/:id", appManager.UpdateDingDianTui)
	dingdiantui.DELETE("/order/:id", appManager.DeleteDingDianTui)
	keliujing := router.Group("/keliujing")
	keliujing.POST("/order", appManager.CreateKeLiuJing)
	keliujing.GET("/orders", appManager.QueryKeLiuJings)
	keliujing.GET("/order/:id", appManager.QueryKeLiuJing)
	keliujing.PUT("/order/:id", appManager.UpdateKeLiuJing)
	keliujing.DELETE("/order/:id", appManager.DeleteKeLiuJing)
	chaoshixin := router.Group("/chaoshixin")
	chaoshixin.POST("/order", appManager.CreateChaoShiXin)
	chaoshixin.GET("/orders", appManager.QueryChaoShiXins)
	chaoshixin.GET("/order/:id", appManager.QueryChaoShiXin)
	chaoshixin.PUT("/order/:id", appManager.UpdateChaoShiXin)
	chaoshixin.DELETE("/order/:id", appManager.DeleteChaoShiXin)
	aiyunhu := router.Group("/aiyunhu")
	aiyunhu.POST("/order", appManager.CreateAIYunHu)
	aiyunhu.GET("/orders", appManager.QueryAIYunHus)
	aiyunhu.GET("/order/:id", appManager.QueryAIYunHu)
	aiyunhu.PUT("/order/:id", appManager.UpdateAIYunHu)
	aiyunhu.DELETE("/order/:id", appManager.DeleteAIYunHu)
	return router
}
