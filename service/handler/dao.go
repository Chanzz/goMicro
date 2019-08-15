package handler

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"goMicro/service/common"
	proto "goMicro/service/proto"
	"log"
)

type AppDao struct {
	DB *sql.DB
}

func (d *AppDao) CreateOrder(ctx context.Context, in *proto.OrderInfo, out *proto.Response) error {
	sql := "INSERT INTO order_info(userid,location,hold_time,preheat_time,push_range) VALUES(?,?,?,?,?)"
	message := "{\"userid\":\"" + in.Token[0:32] + "\",\"location\":\"" + in.Location + "\",\"hold_time\":\"" + in.HoldTime + "\",\"preheat_time\":\"" + in.PreheatTime + "\",\"push_range\":\"" + in.PushRange + "\"}"
	_, err := d.DB.Exec(sql, in.Token[0:32], in.Location, in.HoldTime, in.PreheatTime, in.PushRange)
	if err != nil {
		log.Println(err)
		out.ErrMsg = err.Error()
		out.ErrCode = 400
	}
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
		out.ErrMsg = err.Error()
		out.ErrCode = 400
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
		info += "{\"location\":\"" + location + "\",\"hold_time\":\"" + hold_time + "\",\"preheat_time\":\"" + preheat_time + "\",\"push_range\":\"" + push_range + "\",\"submit_time\":\"" + submit_time + "\",\"id\":\"" + id + "\"},"
	}
	info += "]"
	out.Info = info
	out.ErrMsg = "查询成功"
	out.ErrCode = 200
	return nil
}
func (d *AppDao) QueryOrder(ctx context.Context, in *proto.OrderInfo, out *proto.Response) error {
	sql := "SELECT * FROM order_info WHERE id=?,userid=?"
	rows, err := d.DB.Query(sql, in.Id, in.Token[0:32])
	defer rows.Close()
	if err != nil {
		log.Println(err)
		out.ErrMsg = "查询失败"
		out.ErrCode = 400
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
		info = "{\"location\":\"" + location + "\",\"hold_time\":\"" + hold_time + "\",\"preheat_time\":\"" + preheat_time + "\",\"push_range\":\"" + push_range + "\",\"submit_time\":\"" + submit_time + "\",\"id\":\"" + id + "\"}"
	}
	out.Info = info
	out.ErrMsg = "查询成功"
	out.ErrCode = 200
	return nil
}
