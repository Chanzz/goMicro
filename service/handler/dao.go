package handler

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	protobuf "github.com/golang/protobuf/proto"
	"github.com/micro/go-micro"
	"strings"

	//"github.com/micro/go-micro/broker"
	"goMicro/service/common"
	proto "goMicro/service/proto"
	"log"
	"time"
)

type AppDao struct {
	DB  *sql.DB
	PUB micro.Publisher
}

//func publish(b broker.Broker, topic string, msg *broker.Message) {
//	if err := b.Publish(topic, msg); err != nil {
//		log.Println(err)
//	}
//}
func (d *AppDao) CreateDingDianTui(ctx context.Context, in *proto.DingDianTuiInfo, out *proto.Response) error {
	word := &proto.Word{}
	word.Type = "定点推"
	word.UserId = in.Token[0:32]
	word.Phone = in.Phone
	word.Info = "{\"活动地点\":\"" + in.Location + "\",\"活动时间\":\"" + in.HoldTime + "\",\"预热时间\":\"" + in.PreheatTime + "\",\"推送范围\":\"" + in.PushRange + "\"}"
	message, _ := protobuf.Marshal(word)
	conn := common.ConnectRedis()
	_, err := conn.Do("PUBLISH", "change_word", message)
	if err != nil {
		log.Println(err)
		out.ErrMsg = "提交失败"
		out.ErrCode = 400
		return nil
	}
	out.ErrMsg = "提交成功"
	out.ErrCode = 200
	return nil
}

func (d *AppDao) UpdateDingDianTui(ctx context.Context, in *proto.DingDianTuiInfo, out *proto.Response) error {
	//update_sql := "UPDATE order_info SET location=?,hold_time=?,preheat_time=?,push_range=? WHERE id=? AND userid=?"
	//_, err := d.DB.Exec(update_sql, in.Location, in.HoldTime, in.PreheatTime, in.PushRange, in.Id, in.Token[0:32])
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "更新失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//out.ErrMsg = "更新成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) DeleteDingDianTui(ctx context.Context, in *proto.DingDianTuiInfo, out *proto.Response) error {
	//delete_sql := "DELETE FROM order_info where id=? AND userid=?"
	//_, err := d.DB.Exec(delete_sql, in.Id, in.Token[0:32])
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "删除失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//out.ErrMsg = "删除成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) QueryDingDianTuis(ctx context.Context, in *proto.DingDianTuiInfo, out *proto.Response) error {
	//sql := "SELECT * FROM order_info WHERE userid=?"
	//rows, err := d.DB.Query(sql, in.Token[0:32])
	//defer rows.Close()
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "查询失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//var info string
	//info += "["
	//for rows.Next() {
	//	var token string
	//	var location string
	//	var hold_time string
	//	var preheat_time string
	//	var push_range string
	//	var id string
	//	var submit_time string
	//	err = rows.Scan(&token, &location, &hold_time, &preheat_time, &push_range, &submit_time, &id)
	//	info += "{'location':'" + location + "','hold_time':'" + hold_time + "','preheat_time':'" + preheat_time + "','push_range':'" + push_range + "','submit_time':'" + submit_time + "','id':'" + id + "'},"
	//}
	//info += "]"
	//out.Info = info
	//out.ErrMsg = "查询成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) QueryDingDianTui(ctx context.Context, in *proto.DingDianTuiInfo, out *proto.Response) error {
	//sql := "SELECT * FROM order_info WHERE id=? AND userid=?"
	//rows, err := d.DB.Query(sql, in.Id, in.Token[0:32])
	//defer rows.Close()
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "查询失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//var info string
	//for rows.Next() {
	//	var token string
	//	var location string
	//	var hold_time string
	//	var preheat_time string
	//	var push_range string
	//	var id string
	//	var submit_time string
	//	err = rows.Scan(&token, &location, &hold_time, &preheat_time, &push_range, &submit_time, &id)
	//	info = "{'location':'" + location + "','hold_time':'" + hold_time + "','preheat_time':'" + preheat_time + "','push_range':'" + push_range + "','submit_time':'" + submit_time + "','id':'" + id + "'}"
	//}
	//if info == "" {
	//	out.ErrMsg = "查询失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//out.Info = info
	//out.ErrMsg = "查询成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) CreateKeLiuJing(ctx context.Context, in *proto.KeLiuJingInfo, out *proto.Response) error {
	word := &proto.Word{}
	word.Type = "客流镜"
	word.UserId = in.Token[0:32]
	word.Phone = in.Phone
	word.Info = "{\"监控面积\":\"" + in.MonitorArea + "\",\"需求版本\":\"" + in.Version + "\",\"店铺名称\":\"" + in.Name + "\"}"
	message, _ := protobuf.Marshal(word)
	conn := common.ConnectRedis()
	_, err := conn.Do("PUBLISH", "change_word", message)

	if err != nil {
		log.Println(err)
		out.ErrMsg = "提交失败"
		out.ErrCode = 400
		return nil
	}
	out.ErrMsg = "提交成功"
	out.ErrCode = 200
	return nil
}

func (d *AppDao) UpdateKeLiuJing(ctx context.Context, in *proto.KeLiuJingInfo, out *proto.Response) error {
	//update_sql := "UPDATE order_info SET location=?,hold_time=?,preheat_time=?,push_range=? WHERE id=? AND userid=?"
	//_, err := d.DB.Exec(update_sql, in.Location, in.HoldTime, in.PreheatTime, in.PushRange, in.Id, in.Token[0:32])
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "更新失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//out.ErrMsg = "更新成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) DeleteKeLiuJing(ctx context.Context, in *proto.KeLiuJingInfo, out *proto.Response) error {
	//delete_sql := "DELETE FROM order_info where id=? AND userid=?"
	//_, err := d.DB.Exec(delete_sql, in.Id, in.Token[0:32])
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "删除失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//out.ErrMsg = "删除成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) QueryKeLiuJings(ctx context.Context, in *proto.KeLiuJingInfo, out *proto.Response) error {
	//sql := "SELECT * FROM order_info WHERE userid=?"
	//rows, err := d.DB.Query(sql, in.Token[0:32])
	//defer rows.Close()
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "查询失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//var info string
	//info += "["
	//for rows.Next() {
	//	var token string
	//	var location string
	//	var hold_time string
	//	var preheat_time string
	//	var push_range string
	//	var id string
	//	var submit_time string
	//	err = rows.Scan(&token, &location, &hold_time, &preheat_time, &push_range, &submit_time, &id)
	//	info += "{'location':'" + location + "','hold_time':'" + hold_time + "','preheat_time':'" + preheat_time + "','push_range':'" + push_range + "','submit_time':'" + submit_time + "','id':'" + id + "'},"
	//}
	//info += "]"
	//out.Info = info
	//out.ErrMsg = "查询成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) QueryKeLiuJing(ctx context.Context, in *proto.KeLiuJingInfo, out *proto.Response) error {
	//sql := "SELECT * FROM order_info WHERE id=? AND userid=?"
	//rows, err := d.DB.Query(sql, in.Id, in.Token[0:32])
	//defer rows.Close()
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "查询失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//var info string
	//for rows.Next() {
	//	var token string
	//	var location string
	//	var hold_time string
	//	var preheat_time string
	//	var push_range string
	//	var id string
	//	var submit_time string
	//	err = rows.Scan(&token, &location, &hold_time, &preheat_time, &push_range, &submit_time, &id)
	//	info = "{'location':'" + location + "','hold_time':'" + hold_time + "','preheat_time':'" + preheat_time + "','push_range':'" + push_range + "','submit_time':'" + submit_time + "','id':'" + id + "'}"
	//}
	//if info == "" {
	//	out.ErrMsg = "查询失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//out.Info = info
	//out.ErrMsg = "查询成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) CreateChaoShiXin(ctx context.Context, in *proto.ChaoShiXinInfo, out *proto.Response) error {
	word := &proto.Word{}
	word.Type = "超视信"
	word.UserId = in.Token[0:32]
	word.Phone = in.Phone
	word.Info = "{\"发送数量\":\"" + in.SendNum + "\",\"发送时间\":\"" + in.SendTime + "\",\"内容主题\":\"" + in.Theme + "\",\"触达线索\":\"" + in.Clew + "\"}"
	message, _ := protobuf.Marshal(word)
	conn := common.ConnectRedis()
	_, err := conn.Do("PUBLISH", "change_word", message)

	if err != nil {
		log.Println(err)
		out.ErrMsg = "提交失败"
		out.ErrCode = 400
		return nil
	}
	out.ErrMsg = "提交成功"
	out.ErrCode = 200
	return nil
}

func (d *AppDao) UpdateChaoShiXin(ctx context.Context, in *proto.ChaoShiXinInfo, out *proto.Response) error {
	//update_sql := "UPDATE order_info SET location=?,hold_time=?,preheat_time=?,push_range=? WHERE id=? AND userid=?"
	//_, err := d.DB.Exec(update_sql, in.Location, in.HoldTime, in.PreheatTime, in.PushRange, in.Id, in.Token[0:32])
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "更新失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//out.ErrMsg = "更新成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) DeleteChaoShiXin(ctx context.Context, in *proto.ChaoShiXinInfo, out *proto.Response) error {
	//delete_sql := "DELETE FROM order_info where id=? AND userid=?"
	//_, err := d.DB.Exec(delete_sql, in.Id, in.Token[0:32])
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "删除失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//out.ErrMsg = "删除成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) QueryChaoShiXins(ctx context.Context, in *proto.ChaoShiXinInfo, out *proto.Response) error {
	//sql := "SELECT * FROM order_info WHERE userid=?"
	//rows, err := d.DB.Query(sql, in.Token[0:32])
	//defer rows.Close()
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "查询失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//var info string
	//info += "["
	//for rows.Next() {
	//	var token string
	//	var location string
	//	var hold_time string
	//	var preheat_time string
	//	var push_range string
	//	var id string
	//	var submit_time string
	//	err = rows.Scan(&token, &location, &hold_time, &preheat_time, &push_range, &submit_time, &id)
	//	info += "{'location':'" + location + "','hold_time':'" + hold_time + "','preheat_time':'" + preheat_time + "','push_range':'" + push_range + "','submit_time':'" + submit_time + "','id':'" + id + "'},"
	//}
	//info += "]"
	//out.Info = info
	//out.ErrMsg = "查询成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) QueryChaoShiXin(ctx context.Context, in *proto.ChaoShiXinInfo, out *proto.Response) error {
	//sql := "SELECT * FROM order_info WHERE id=? AND userid=?"
	//rows, err := d.DB.Query(sql, in.Id, in.Token[0:32])
	//defer rows.Close()
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "查询失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//var info string
	//for rows.Next() {
	//	var token string
	//	var location string
	//	var hold_time string
	//	var preheat_time string
	//	var push_range string
	//	var id string
	//	var submit_time string
	//	err = rows.Scan(&token, &location, &hold_time, &preheat_time, &push_range, &submit_time, &id)
	//	info = "{'location':'" + location + "','hold_time':'" + hold_time + "','preheat_time':'" + preheat_time + "','push_range':'" + push_range + "','submit_time':'" + submit_time + "','id':'" + id + "'}"
	//}
	//if info == "" {
	//	out.ErrMsg = "查询失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//out.Info = info
	//out.ErrMsg = "查询成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) CreateAIYunHu(ctx context.Context, in *proto.AIYunHuInfo, out *proto.Response) error {
	word := &proto.Word{}
	word.Type = "AI云呼"
	word.UserId = in.Token[0:32]
	word.Phone = in.Phone
	word.Info = "{\"拨打数量\":\"" + in.DialNum + "\",\"拨打时间\":\"" + in.DialTime + "\",\"内容主题\":\"" + in.Theme + "\",\"触达线索\":\"" + in.Clew + "\"}"
	message, _ := protobuf.Marshal(word)
	conn := common.ConnectRedis()
	_, err := conn.Do("PUBLISH", "change_word", message)

	if err != nil {
		log.Println(err)
		out.ErrMsg = "提交失败"
		out.ErrCode = 400
		return nil
	}
	out.ErrMsg = "提交成功"
	out.ErrCode = 200
	return nil
}

func (d *AppDao) UpdateAIYunHu(ctx context.Context, in *proto.AIYunHuInfo, out *proto.Response) error {
	//update_sql := "UPDATE order_info SET location=?,hold_time=?,preheat_time=?,push_range=? WHERE id=? AND userid=?"
	//_, err := d.DB.Exec(update_sql, in.Location, in.HoldTime, in.PreheatTime, in.PushRange, in.Id, in.Token[0:32])
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "更新失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//out.ErrMsg = "更新成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) DeleteAIYunHu(ctx context.Context, in *proto.AIYunHuInfo, out *proto.Response) error {
	//delete_sql := "DELETE FROM order_info where id=? AND userid=?"
	//_, err := d.DB.Exec(delete_sql, in.Id, in.Token[0:32])
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "删除失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//out.ErrMsg = "删除成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) QueryAIYunHus(ctx context.Context, in *proto.AIYunHuInfo, out *proto.Response) error {
	//sql := "SELECT * FROM order_info WHERE userid=?"
	//rows, err := d.DB.Query(sql, in.Token[0:32])
	//defer rows.Close()
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "查询失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//var info string
	//info += "["
	//for rows.Next() {
	//	var token string
	//	var location string
	//	var hold_time string
	//	var preheat_time string
	//	var push_range string
	//	var id string
	//	var submit_time string
	//	err = rows.Scan(&token, &location, &hold_time, &preheat_time, &push_range, &submit_time, &id)
	//	info += "{'location':'" + location + "','hold_time':'" + hold_time + "','preheat_time':'" + preheat_time + "','push_range':'" + push_range + "','submit_time':'" + submit_time + "','id':'" + id + "'},"
	//}
	//info += "]"
	//out.Info = info
	//out.ErrMsg = "查询成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) QueryAIYunHu(ctx context.Context, in *proto.AIYunHuInfo, out *proto.Response) error {
	//sql := "SELECT * FROM order_info WHERE id=? AND userid=?"
	//rows, err := d.DB.Query(sql, in.Id, in.Token[0:32])
	//defer rows.Close()
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "查询失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//var info string
	//for rows.Next() {
	//	var token string
	//	var location string
	//	var hold_time string
	//	var preheat_time string
	//	var push_range string
	//	var id string
	//	var submit_time string
	//	err = rows.Scan(&token, &location, &hold_time, &preheat_time, &push_range, &submit_time, &id)
	//	info = "{'location':'" + location + "','hold_time':'" + hold_time + "','preheat_time':'" + preheat_time + "','push_range':'" + push_range + "','submit_time':'" + submit_time + "','id':'" + id + "'}"
	//}
	//if info == "" {
	//	out.ErrMsg = "查询失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//out.Info = info
	//out.ErrMsg = "查询成功"
	//out.ErrCode = 200
	return nil
}
func (d *AppDao) CreateCheShiChuang(ctx context.Context, in *proto.CheShiChuangInfo, out *proto.Response) error {
	word := &proto.Word{}
	word.Type = "车视窗"
	word.UserId = in.Token[0:32]
	word.Info = "{\"推广计划名称\":\"" + in.Name + "\",\"投放时间\":\"{" + strings.Join(in.Time, ",") + "}\",\"价格\":\"" + in.Price + "\",\"地域\":\"{" + strings.Join(in.Place, ",") + "}\",\"媒体资源\":\"{" + strings.Join(in.Media, ",") + "}\",\"广告模式\":\"" + in.ModelType + "\",\"投放策略\":\"" + in.Strategy + "\",\"性别\":\"{" + strings.Join(in.Sex, ",") + "}\",\"年龄\":\"{" + strings.Join(in.Age, ",") + "}\",\"竞品车型\":\"{" + strings.Join(in.Object, ",") + "}\",\"竞品市场\":\"{" + strings.Join(in.Makert, ",") + "}\",\"本品车型\":\"" + in.Car + "\"}"
	message, _ := protobuf.Marshal(word)
	conn := common.ConnectRedis()
	_, err := conn.Do("PUBLISH", "change_word", message)

	if err != nil {
		log.Println(err)
		out.ErrMsg = "提交失败"
		out.ErrCode = 400
		return nil
	}
	out.ErrMsg = "提交成功"
	out.ErrCode = 200
	return nil
}

func (d *AppDao) UpdateCheShiChuang(ctx context.Context, in *proto.CheShiChuangInfo, out *proto.Response) error {
	//update_sql := "UPDATE order_info SET location=?,hold_time=?,preheat_time=?,push_range=? WHERE id=? AND userid=?"
	//_, err := d.DB.Exec(update_sql, in.Location, in.HoldTime, in.PreheatTime, in.PushRange, in.Id, in.Token[0:32])
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "更新失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//out.ErrMsg = "更新成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) DeleteCheShiChuang(ctx context.Context, in *proto.CheShiChuangInfo, out *proto.Response) error {
	//delete_sql := "DELETE FROM order_info where id=? AND userid=?"
	//_, err := d.DB.Exec(delete_sql, in.Id, in.Token[0:32])
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "删除失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//out.ErrMsg = "删除成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) QueryCheShiChuangs(ctx context.Context, in *proto.CheShiChuangInfo, out *proto.Response) error {
	//sql := "SELECT * FROM order_info WHERE userid=?"
	//rows, err := d.DB.Query(sql, in.Token[0:32])
	//defer rows.Close()
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "查询失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//var info string
	//info += "["
	//for rows.Next() {
	//	var token string
	//	var location string
	//	var hold_time string
	//	var preheat_time string
	//	var push_range string
	//	var id string
	//	var submit_time string
	//	err = rows.Scan(&token, &location, &hold_time, &preheat_time, &push_range, &submit_time, &id)
	//	info += "{'location':'" + location + "','hold_time':'" + hold_time + "','preheat_time':'" + preheat_time + "','push_range':'" + push_range + "','submit_time':'" + submit_time + "','id':'" + id + "'},"
	//}
	//info += "]"
	//out.Info = info
	//out.ErrMsg = "查询成功"
	//out.ErrCode = 200
	return nil
}

func (d *AppDao) QueryCheShiChuang(ctx context.Context, in *proto.CheShiChuangInfo, out *proto.Response) error {
	//sql := "SELECT * FROM order_info WHERE id=? AND userid=?"
	//rows, err := d.DB.Query(sql, in.Id, in.Token[0:32])
	//defer rows.Close()
	//if err != nil {
	//	log.Println(err)
	//	out.ErrMsg = "查询失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//var info string
	//for rows.Next() {
	//	var token string
	//	var location string
	//	var hold_time string
	//	var preheat_time string
	//	var push_range string
	//	var id string
	//	var submit_time string
	//	err = rows.Scan(&token, &location, &hold_time, &preheat_time, &push_range, &submit_time, &id)
	//	info = "{'location':'" + location + "','hold_time':'" + hold_time + "','preheat_time':'" + preheat_time + "','push_range':'" + push_range + "','submit_time':'" + submit_time + "','id':'" + id + "'}"
	//}
	//if info == "" {
	//	out.ErrMsg = "查询失败"
	//	out.ErrCode = 400
	//	return nil
	//}
	//out.Info = info
	//out.ErrMsg = "查询成功"
	//out.ErrCode = 200
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
