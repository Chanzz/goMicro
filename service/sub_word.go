package main

import (
	"baliance.com/gooxml/document"
	"encoding/json"
	"fmt"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/gomodule/redigo/redis"
	"goMicro/service/common"
	"goMicro/service/handler"
	proto "goMicro/service/proto"
	"gopkg.in/gomail.v2"
	"log"
	"strconv"
)

func subs() {
	conn := common.ConnectRedis()
	defer conn.Close()
	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe("change_word")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
			docxPath, excelPath := new_word(v.Data)
			err := send_email(docxPath, excelPath)
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
func new_word(message []byte) (string, string) {
	word := &proto.Word{}
	protobuf.Unmarshal(message, word)
	dataArr := map[string]string{}
	json.Unmarshal([]byte(word.Info), &dataArr)
	db := handler.CreateMySqlConnection()
	defer db.Close()
	sql := "SELECT user_name,user_company,phone,manufacturer_id FROM dealer_user WHERE id=?"
	rows, _ := db.Query(sql, word.UserId)
	var user_info string
	var user_name string
	var user_company string
	var phone string
	var manufacturer_id string
	var manufacturer_name string
	for rows.Next() {
		rows.Scan(&user_name, &user_company, &phone, &manufacturer_id)
		sql1 := "SELECT name FROM manufacturer_name WHERE id=?"
		rows1, _ := db.Query(sql1, manufacturer_id)
		for rows1.Next() {
			rows1.Scan(&manufacturer_name)
			user_info = "品牌名称： " + manufacturer_name + "  用户名称： " + user_name + "  用户公司： " + user_company + "  联系方式： " + phone
		}
	}
	doc := document.New()
	para := doc.AddParagraph()
	para.SetStyle("Subtitle")
	if word.UserId != "" {
		para.AddRun().AddText(user_info)
	}
	para = doc.AddParagraph()
	para.AddRun().AddText("电话：" + word.Phone)
	for i, j := range dataArr {
		para = doc.AddParagraph()
		para.AddRun().AddText(i + "：")
		para.AddRun().AddText(j)
	}
	docxPath := "../docx/" + word.Type + "--" + manufacturer_name + common.Create_captcha() + ".docx"
	doc.SaveToFile(docxPath)
	insert_sql := "INSERT INTO order_info(user_id,phone,user_name,file_path,order_type) VALUES(?,?,?,?,?)"
	_, err1 := db.Exec(insert_sql, word.UserId, word.Phone, user_name, docxPath, word.Type)
	if err1 != nil {
		log.Println("插入失败")
		log.Println(err1)
	}
	return docxPath, word.Annex
}
func send_email(docxPath string, excelPath string) error {
	configMap := common.GetConfigMap("email")
	user := configMap["user"]
	password := configMap["passwd"]
	host := configMap["host"]
	port, _ := strconv.Atoi(configMap["port"]) //转换端口类型为int
	m := gomail.NewMessage()
	m.SetHeader("From", "<"+"haoji_chen@data-spark.cn"+">")
	m.SetHeader("To", "295218321@qq.com") //发送给多个用户
	subject := "订单请求--" + docxPath[8:]
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", "请查收附件") //设置邮件正文
	d := gomail.NewDialer(host, port, user, password)
	m.Attach(docxPath)
	if excelPath != "" {
		m.Attach(excelPath)
	}
	err := d.DialAndSend(m)
	return err
}
func main() {
	subs()
}
