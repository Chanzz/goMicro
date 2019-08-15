package main

import (
	"baliance.com/gooxml/document"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"goMicro/service/common"
	"goMicro/service/handler"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type orderInfo struct {
	Userid       string `json:"userid"`
	Location     string `json:"location"`
	Hold_time    string `json:"hold_time"`
	Preheat_time string `json:"preheat_time"`
	Push_range   string `json:"push_range"`
}

func create_captcha() string {
	return fmt.Sprintf("%03v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000))
}
func subs() {
	conn := common.ConnectRedis()
	defer conn.Close()
	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe("change_word")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
			docxPath := new_word(v.Data)
			err := send_email(docxPath)
			if err != nil {
				log.Println(err)
			}
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}
func new_word(message []byte) string {
	var order_info orderInfo
	err := json.Unmarshal(message, &order_info)
	if err != nil {
		log.Fatalf("error opening Word: %s", err)
	}
	db := handler.CreateMySqlConnection()
	defer db.Close()
	sql := "SELECT user_name,user_company,phone,manufacturer_id FROM dealer_user WHERE id=?"
	rows, _ := db.Query(sql, order_info.Userid)
	var user_info string
	var user_name string
	var user_company string
	var phone string
	var manufacturer_id string
	var manufacturer_name string
	for rows.Next() {
		err = rows.Scan(&user_name, &user_company, &phone, &manufacturer_id)
		sql1 := "SELECT name FROM manufacturer_name WHERE id=?"
		rows1, _ := db.Query(sql1, manufacturer_id)
		for rows1.Next() {
			err = rows1.Scan(&manufacturer_name)
			user_info = "品牌名称 " + manufacturer_name + "  用户名称 " + user_name + "  用户公司 " + user_company + "  联系方式 " + phone
		}
	}
	doc := document.New()
	para := doc.AddParagraph()
	para.SetStyle("Subtitle")
	para.AddRun().AddText("用户信息：" + user_info)
	para = doc.AddParagraph()
	para.AddRun().AddText("活动地点：" + order_info.Location)
	para = doc.AddParagraph()
	para.AddRun().AddText("活动时间：" + order_info.Hold_time)
	para = doc.AddParagraph()
	para.AddRun().AddText("预热时间：" + order_info.Preheat_time)
	para = doc.AddParagraph()
	para.AddRun().AddText("推送范围：" + order_info.Push_range)
	docxPath := "../docx/" + manufacturer_name + create_captcha() + ".docx"
	doc.SaveToFile(docxPath)
	return docxPath
}
func send_email(docxPath string) error {
	configMap := common.GetConfigMap("email")
	user := configMap["user"]
	password := configMap["passwd"]
	host := configMap["host"]
	port, _ := strconv.Atoi(configMap["port"]) //转换端口类型为int
	m := gomail.NewMessage()
	m.SetHeader("From", "<"+"haoji_chen@data-spark.cn"+">")
	m.SetHeader("To", "295218321@qq.com") //发送给多个用户
	m.SetHeader("Subject", "测试")          //设置邮件主题
	m.SetBody("text/html", "测试")          //设置邮件正文
	d := gomail.NewDialer(host, port, user, password)
	m.Attach(docxPath)
	err := d.DialAndSend(m)
	return err
}
func main() {
	subs()
}
