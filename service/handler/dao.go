package handler

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"goMicro/service/common"
	proto "goMicro/service/proto"
	"log"
	"time"
)

type AppDao struct {
	DB  *sql.DB
	PUB micro.Publisher
}

func publish(b broker.Broker, topic string, msg *broker.Message) {
	if err := b.Publish(topic, msg); err != nil {
		log.Println(err)
	}
}
func (d *AppDao) CreateOrder(ctx context.Context, in *proto.OrderInfo, out *proto.Response) error {
	sql := "INSERT INTO order_info(userid,location,hold_time,preheat_time,push_range) VALUES(?,?,?,?,?)"
	message := "{\"userid\":\"" + in.Token[0:32] + "\",\"location\":\"" + in.Location + "\",\"hold_time\":\"" + in.HoldTime + "\",\"preheat_time\":\"" + in.PreheatTime + "\",\"push_range\":\"" + in.PushRange + "\"}"
	_, err := d.DB.Exec(sql, in.Token[0:32], in.Location, in.HoldTime, in.PreheatTime, in.PushRange)
	if err != nil {
		log.Println(err)
		out.ErrMsg = "提交失败"
		out.ErrCode = 400
		return nil
	}
	//data,_:=protobuf.Marshal(in)
	//d.PUB.Publish(context.TODO(), in)
	//b := broker.NewBroker(broker.Addrs("127.0.0.1:6666"))
	//b.Init()
	//if err := b.Connect(); err != nil {
	//	log.Println(err)
	//}
	//publish(b, "test", &broker.Message{
	//	Body: []byte("hello"),
	//})
	//defer b.Disconnect()
	conn := common.ConnectRedis()
	_, err1 := conn.Do("PUBLISH", "change_word", message)

	if err1 != nil {
		log.Println(err1)
	}
	out.ErrMsg = "提交成功"
	out.ErrCode = 200
	return nil
}

func (d *AppDao) UpdateOrder(ctx context.Context, in *proto.OrderInfo, out *proto.Response) error {
	update_sql := "UPDATE order_info SET location=?,hold_time=?,preheat_time=?,push_range=? WHERE id=? AND userid=?"
	_, err := d.DB.Exec(update_sql, in.Location, in.HoldTime, in.PreheatTime, in.PushRange, in.Id, in.Token[0:32])
	if err != nil {
		log.Println(err)
		out.ErrMsg = "更新失败"
		out.ErrCode = 400
		return nil
	}
	out.ErrMsg = "更新成功"
	out.ErrCode = 200
	return nil
}

func (d *AppDao) DeleteOrder(ctx context.Context, in *proto.OrderInfo, out *proto.Response) error {
	delete_sql := "DELETE FROM order_info where id=? AND userid=?"
	_, err := d.DB.Exec(delete_sql, in.Id, in.Token[0:32])
	if err != nil {
		log.Println(err)
		out.ErrMsg = "删除失败"
		out.ErrCode = 400
		return nil
	}
	out.ErrMsg = "删除成功"
	out.ErrCode = 200
	return nil
}

func (d *AppDao) QueryOrders(ctx context.Context, in *proto.OrderInfo, out *proto.Response) error {
	sql := "SELECT * FROM order_info WHERE userid=?"
	rows, err := d.DB.Query(sql, in.Token[0:32])
	defer rows.Close()
	if err != nil {
		log.Println(err)
		out.ErrMsg = "查询失败"
		out.ErrCode = 400
		return nil
	}
	var info string
	info += "["
	for rows.Next() {
		var token string
		var location string
		var hold_time string
		var preheat_time string
		var push_range string
		var id string
		var submit_time string
		err = rows.Scan(&token, &location, &hold_time, &preheat_time, &push_range, &submit_time, &id)
		info += "{'location':'" + location + "','hold_time':'" + hold_time + "','preheat_time':'" + preheat_time + "','push_range':'" + push_range + "','submit_time':'" + submit_time + "','id':'" + id + "'},"
	}
	info += "]"
	out.Info = info
	out.ErrMsg = "查询成功"
	out.ErrCode = 200
	return nil
}
func (d *AppDao) QueryOrder(ctx context.Context, in *proto.OrderInfo, out *proto.Response) error {
	sql := "SELECT * FROM order_info WHERE id=? AND userid=?"
	rows, err := d.DB.Query(sql, in.Id, in.Token[0:32])
	defer rows.Close()
	if err != nil {
		log.Println(err)
		out.ErrMsg = "查询失败"
		out.ErrCode = 400
		return nil
	}
	var info string
	for rows.Next() {
		var token string
		var location string
		var hold_time string
		var preheat_time string
		var push_range string
		var id string
		var submit_time string
		err = rows.Scan(&token, &location, &hold_time, &preheat_time, &push_range, &submit_time, &id)
		info = "{'location':'" + location + "','hold_time':'" + hold_time + "','preheat_time':'" + preheat_time + "','push_range':'" + push_range + "','submit_time':'" + submit_time + "','id':'" + id + "'}"
	}
	if info == "" {
		out.ErrMsg = "查询失败"
		out.ErrCode = 400
		return nil
	}
	out.Info = info
	out.ErrMsg = "查询成功"
	out.ErrCode = 200
	return nil
}
func (d *AppDao) Login(ctx context.Context, in *proto.LoginInfo, out *proto.Response) error {
	sql := "SELECT * FROM dealer_user WHERE phone=?"
	err := d.DB.QueryRow(sql, in.PhoneNum)
	if err != nil {
		log.Print(err)
		out.ErrCode = 400
		return nil
	}
	sql = "UPDATE dealer_user SET last_login_ip=? ,last_login_time=? WHERE phone=?"
	d.DB.Exec(sql, in.Ip, time.Now(), in.PhoneNum)

	sql = "insert into customer_login(phone,login_time) values(?,?)"
	d.DB.Exec(sql, in.PhoneNum, time.Now())
	// 获取连接池一个连接
	conn := common.ConnectRedis()
	defer conn.Close()

	key := common.GenerateMD5(in.PhoneNum) // 获取对应用户名 md5 的字符串作为 key
	// 每次登陆都会刷新 Token
	// 为当前时间戳生成 md5
	token := common.GenerateMD5(string(time.Now().Unix()))
	// 生成用户的 token 并存储起来，设置一周过期时间
	conn.Do("set", key, token, "EX", "604800")
	out.ErrCode = 200
	out.Info = "{'token':'" + key + token + "'}"
	return nil
}
func (d *AppDao) SendCode(ctx context.Context, in *proto.LoginInfo, out *proto.Response) error {
	return nil
}
