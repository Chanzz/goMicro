package handler

import (
	"github.com/gin-gonic/gin"
	"goMicro/api/client"
	"goMicro/service/common"
	appProto "goMicro/service/proto"
	"log"
)

type AppManager struct {
}

func (a *AppManager) CreateDingDianTui(c *gin.Context) {
	var orderInfo appProto.DingDianTuiInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.CreateDingDianTui(orderInfo)
	c.JSON(200, resp)
}

func (a *AppManager) QueryDingDianTuis(c *gin.Context) {
	var orderInfo appProto.DingDianTuiInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.QueryDingDianTuis(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) QueryDingDianTui(c *gin.Context) {
	var orderInfo appProto.DingDianTuiInfo
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
	resp := client.QueryDingDianTui(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) UpdateDingDianTui(c *gin.Context) {
	//var updateOrderInfo appProto.UpdateOrderInfo
	var orderInfo appProto.DingDianTuiInfo
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
	resp := client.UpdateDingDianTui(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) DeleteDingDianTui(c *gin.Context) {
	var orderInfo appProto.DingDianTuiInfo
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
	resp := client.DeleteDingDianTui(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) CreateAIYunHu(c *gin.Context) {
	var orderInfo appProto.AIYunHuInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.CreateAIYunHu(orderInfo)
	c.JSON(200, resp)
}

func (a *AppManager) QueryAIYunHus(c *gin.Context) {
	var orderInfo appProto.AIYunHuInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.QueryAIYunHus(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) QueryAIYunHu(c *gin.Context) {
	var orderInfo appProto.AIYunHuInfo
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
	resp := client.QueryAIYunHu(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) UpdateAIYunHu(c *gin.Context) {
	var orderInfo appProto.AIYunHuInfo
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
	resp := client.UpdateAIYunHu(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) DeleteAIYunHu(c *gin.Context) {
	var orderInfo appProto.AIYunHuInfo
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
	resp := client.DeleteAIYunHu(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) CreateKeLiuJing(c *gin.Context) {
	var orderInfo appProto.KeLiuJingInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.CreateKeLiuJing(orderInfo)
	c.JSON(200, resp)
}

func (a *AppManager) QueryKeLiuJings(c *gin.Context) {
	var orderInfo appProto.KeLiuJingInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.QueryKeLiuJings(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) QueryKeLiuJing(c *gin.Context) {
	var orderInfo appProto.KeLiuJingInfo
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
	resp := client.QueryKeLiuJing(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) UpdateKeLiuJing(c *gin.Context) {
	var orderInfo appProto.KeLiuJingInfo
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
	resp := client.UpdateKeLiuJing(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) DeleteKeLiuJing(c *gin.Context) {
	var orderInfo appProto.KeLiuJingInfo
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
	resp := client.DeleteKeLiuJing(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) CreateChaoShiXin(c *gin.Context) {
	var orderInfo appProto.ChaoShiXinInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.CreateChaoShiXin(orderInfo)
	c.JSON(200, resp)
}

func (a *AppManager) QueryChaoShiXins(c *gin.Context) {
	var orderInfo appProto.ChaoShiXinInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.QueryChaoShiXins(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) QueryChaoShiXin(c *gin.Context) {
	var orderInfo appProto.ChaoShiXinInfo
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
	resp := client.QueryChaoShiXin(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) UpdateChaoShiXin(c *gin.Context) {
	var orderInfo appProto.ChaoShiXinInfo
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
	resp := client.UpdateChaoShiXin(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) DeleteChaoShiXin(c *gin.Context) {
	var orderInfo appProto.ChaoShiXinInfo
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
	resp := client.DeleteChaoShiXin(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) CreatePinXiaoTong(c *gin.Context) {
	var orderInfo appProto.PinXiaoTongInfo
	orderInfo.Token = c.Query("token")
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.CreatePinXiaoTong(orderInfo)
	c.JSON(200, resp)
}

func (a *AppManager) QueryPinXiaoTongs(c *gin.Context) {
	var orderInfo appProto.PinXiaoTongInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.QueryPinXiaoTongs(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) QueryPinXiaoTong(c *gin.Context) {
	var orderInfo appProto.PinXiaoTongInfo
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
	resp := client.QueryPinXiaoTong(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) UpdatePinXiaoTong(c *gin.Context) {
	var orderInfo appProto.PinXiaoTongInfo
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
	resp := client.UpdatePinXiaoTong(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) DeletePinXiaoTong(c *gin.Context) {
	var orderInfo appProto.PinXiaoTongInfo
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
	resp := client.DeletePinXiaoTong(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) CreateCheShiChuang(c *gin.Context) {
	var orderInfo appProto.CheShiChuangInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.CreateCheShiChuang(orderInfo)
	c.JSON(200, resp)
}

func (a *AppManager) QueryCheShiChuangs(c *gin.Context) {
	var orderInfo appProto.CheShiChuangInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.QueryCheShiChuangs(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) QueryCheShiChuang(c *gin.Context) {
	var orderInfo appProto.CheShiChuangInfo
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
	resp := client.QueryCheShiChuang(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) UpdateCheShiChuang(c *gin.Context) {
	var orderInfo appProto.CheShiChuangInfo
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
	resp := client.UpdateCheShiChuang(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) DeleteCheShiChuang(c *gin.Context) {
	var orderInfo appProto.CheShiChuangInfo
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
	resp := client.DeleteCheShiChuang(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) CreateWaQianKe(c *gin.Context) {
	var orderInfo appProto.WaQianKeInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.CreateWaQianKe(orderInfo)
	c.JSON(200, resp)
}

func (a *AppManager) QueryWaQianKes(c *gin.Context) {
	var orderInfo appProto.WaQianKeInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.QueryWaQianKes(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) QueryWaQianKe(c *gin.Context) {
	var orderInfo appProto.WaQianKeInfo
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
	resp := client.QueryWaQianKe(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) UpdateWaQianKe(c *gin.Context) {
	var orderInfo appProto.WaQianKeInfo
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
	resp := client.UpdateWaQianKe(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) DeleteWaQianKe(c *gin.Context) {
	var orderInfo appProto.WaQianKeInfo
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
	resp := client.DeleteWaQianKe(orderInfo)
	c.JSON(200, resp)
}

func (a *AppManager) CreateYunHuaXiang(c *gin.Context) {
	var orderInfo appProto.YunHuaXiangInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.CreateYunHuaXiang(orderInfo)
	c.JSON(200, resp)
}

func (a *AppManager) QueryYunHuaXiangs(c *gin.Context) {
	var orderInfo appProto.YunHuaXiangInfo
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.QueryYunHuaXiangs(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) QueryYunHuaXiang(c *gin.Context) {
	var orderInfo appProto.YunHuaXiangInfo
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
	resp := client.QueryYunHuaXiang(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) UpdateYunHuaXiang(c *gin.Context) {
	var orderInfo appProto.YunHuaXiangInfo
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
	resp := client.UpdateYunHuaXiang(orderInfo)
	c.JSON(200, resp)
}
func (a *AppManager) DeleteYunHuaXiang(c *gin.Context) {
	var orderInfo appProto.YunHuaXiangInfo
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
	resp := client.DeleteYunHuaXiang(orderInfo)
	c.JSON(200, resp)
}

func (a *AppManager) Login(c *gin.Context) {
	var loginInfo appProto.LoginInfo
	loginInfo.Ip = c.ClientIP()
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.Login(loginInfo)
	c.JSON(200, resp)
}
func (a *AppManager) UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	path := "../attach/" + common.Create_captcha() + file.Filename
	err := c.SaveUploadedFile(file, path)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, err)
		c.Abort()
	}
	resp := new(appProto.Response)
	resp.ErrMsg = "成功"
	resp.ErrCode = 200
	resp.Info = path
	c.JSON(200, resp)
}
func (a *AppManager) UserInfo(c *gin.Context) {
	var userInfo appProto.UserInfo
	userInfo.Token = c.Query("token")
	resp := client.UserInfo(userInfo)
	c.JSON(200, resp)
}
func (a *AppManager) Register(c *gin.Context) {
	var RegisterInfo appProto.RegisterInfo
	RegisterInfo.Ip = c.ClientIP()
	if err := c.ShouldBindJSON(&RegisterInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.Register(RegisterInfo)
	c.JSON(200, resp)
}
func (a *AppManager) SendCode(c *gin.Context) {
	var loginInfo appProto.LoginInfo
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		log.Println("err  :", err)
		resp := new(appProto.Response)
		resp.ErrMsg = "请求错误"
		resp.ErrCode = 400
		c.JSON(400, resp)
		c.Abort()
		return
	}
	resp := client.SendCode(loginInfo)
	c.JSON(200, resp)
}
