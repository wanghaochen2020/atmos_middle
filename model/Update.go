package model

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"middle/dataReceive/receive"
	"time"
)

const (
	sleepTime    = 1           // 轮询间隔
	spanHalfHour = 30          // 30分钟
	spanDay      = 60 * 24     // 一天
	spanWeek     = 60 * 24 * 7 // 一周
	//spanTest     = 5
)

// 保存各个数据表的更新时间，保存时间戳单位为ms
type Update struct {
	Id        int
	UpdatedAt string
	TableName string
}

// 检查时间是否超时
func checkTime(pTime string, nowTime time.Time, timeout int) bool {
	target, err := time.ParseInLocation("2006-01-02 15:04:05", pTime, time.Local)
	CheckErr(err)
	if t1 := nowTime.Sub(target).Minutes(); int(t1) >= timeout {
		return true
	}
	return false
}

// 输入表名判断是否需要更新
func checkUpdate(table Update, now time.Time) {
	switch table.TableName {
	case "trash_can_info":
		// 检查更新
		if checkTime(table.UpdatedAt, now, spanHalfHour) {
			// 超过时间间隔，更新
			newData := receive.GetDeviceInfo()
			data, err := json.Marshal(newData)
			CheckErr(err)
			msg := Message{"trash_can_info", data}
			msgData, _ := json.Marshal(msg)
			err = RbMQ.Ch.Publish(
				"",      // exchange
				"trash", // routing key
				false,   // mandatory
				false,   // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        msgData,
				})
			FailOnError(err, "Failed to publish a message")
			// 更新update表
			_, err = Db.Collection("update").UpdateOne(context.TODO(), bson.D{{
				Key:   "tableName",
				Value: table.TableName,
			}}, bson.D{
				{"$set", bson.D{{
					Key:   "updatedAt",
					Value: now.Format("2006-01-02 15:04:05"),
				}}},
			})
			CheckErr(err)
		}
	case "trash_can_state":
		if checkTime(table.UpdatedAt, now, spanHalfHour) {
			// 超过时间间隔，更新
			newData := receive.GetDeviceStatus()
			data, err := json.Marshal(newData)
			CheckErr(err)
			msg := Message{"trash_can_state", data}
			msgData, _ := json.Marshal(msg)
			err = RbMQ.Ch.Publish(
				"",      // exchange
				"trash", // routing key
				false,   // mandatory
				false,   // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        msgData,
				})
			FailOnError(err, "Failed to publish a message")
			// 更新update表
			_, err = Db.Collection("update").UpdateOne(context.TODO(), bson.D{{
				Key:   "tableName",
				Value: table.TableName,
			}}, bson.D{
				{"$set", bson.D{{
					Key:   "updatedAt",
					Value: now.Format("2006-01-02 15:04:05"),
				}}},
			})
			CheckErr(err)
		}
	case "trash_can_fault":
		if checkTime(table.UpdatedAt, now, spanHalfHour) {
			// 超过时间间隔，更新
			newData := receive.GetDeviceWarn()
			data, err := json.Marshal(newData)
			CheckErr(err)
			msg := Message{"trash_can_fault", data}
			msgData, _ := json.Marshal(msg)
			err = RbMQ.Ch.Publish(
				"",      // exchange
				"trash", // routing key
				false,   // mandatory
				false,   // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        msgData,
				})
			FailOnError(err, "Failed to publish a message")
			// 更新update表
			_, err = Db.Collection("update").UpdateOne(context.TODO(), bson.D{{
				Key:   "tableName",
				Value: table.TableName,
			}}, bson.D{
				{"$set", bson.D{{
					Key:   "updatedAt",
					Value: now.Format("2006-01-02 15:04:05"),
				}}},
			})
			CheckErr(err)
		}
	case "trash_can_score":
		if checkTime(table.UpdatedAt, now, spanDay) {
			// 超过时间间隔，更新
			newData := receive.GetConsumerScore()
			data, err := json.Marshal(newData)
			CheckErr(err)
			msg := Message{"trash_can_score", data}
			msgData, _ := json.Marshal(msg)
			err = RbMQ.Ch.Publish(
				"",      // exchange
				"trash", // routing key
				false,   // mandatory
				false,   // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        msgData,
				})
			FailOnError(err, "Failed to publish a message")
			// 更新update表
			_, err = Db.Collection("update").UpdateOne(context.TODO(), bson.D{{
				Key:   "tableName",
				Value: table.TableName,
			}}, bson.D{
				{"$set", bson.D{{
					Key:   "updatedAt",
					Value: now.Format("2006-01-02 15:04:05"),
				}}},
			})
			CheckErr(err)
		}
	case "trash_can_delivery":
		if checkTime(table.UpdatedAt, now, spanDay) {
			// 超过时间间隔，更新
			//newData := receive.GetDeliveryRecordByDate("","")
			newData := receive.GetDeliveryRecord(0) // 传0不过滤游客，非零过滤游客
			data, err := json.Marshal(newData)
			CheckErr(err)
			msg := Message{"trash_can_delivery", data}
			msgData, _ := json.Marshal(msg)
			err = RbMQ.Ch.Publish(
				"",      // exchange
				"trash", // routing key
				false,   // mandatory
				false,   // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        msgData,
				})
			FailOnError(err, "Failed to publish a message")
			// 更新update表
			_, err = Db.Collection("update").UpdateOne(context.TODO(), bson.D{{
				Key:   "tableName",
				Value: table.TableName,
			}}, bson.D{
				{"$set", bson.D{{
					Key:   "updatedAt",
					Value: now.Format("2006-01-02 15:04:05"),
				}}},
			})
			CheckErr(err)
		}
	case "trash_can_exchange":
		if checkTime(table.UpdatedAt, now, spanDay) {
			// 超过时间间隔，更新
			newData := receive.GetExchange()
			data, err := json.Marshal(newData)
			CheckErr(err)
			msg := Message{"trash_can_exchange", data}
			msgData, _ := json.Marshal(msg)
			err = RbMQ.Ch.Publish(
				"",      // exchange
				"trash", // routing key
				false,   // mandatory
				false,   // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        msgData,
				})
			FailOnError(err, "Failed to publish a message")
			// 更新update表
			_, err = Db.Collection("update").UpdateOne(context.TODO(), bson.D{{
				Key:   "tableName",
				Value: table.TableName,
			}}, bson.D{
				{"$set", bson.D{{
					Key:   "updatedAt",
					Value: now.Format("2006-01-02 15:04:05"),
				}}},
			})
			CheckErr(err)
		}
	case "unit":
		if checkTime(table.UpdatedAt, now, spanWeek) {
			// 超过时间间隔，更新
			newData := receive.GetUnitList()
			data, err := json.Marshal(newData.Data.Units)
			CheckErr(err)
			msg := Message{"unit", data}
			msgData, _ := json.Marshal(msg)
			err = RbMQ.Ch.Publish(
				"",      // exchange
				"trash", // routing key
				false,   // mandatory
				false,   // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        msgData,
				})
			FailOnError(err, "Failed to publish a message")
			// 更新update表
			_, err = Db.Collection("update").UpdateOne(context.TODO(), bson.D{{
				Key:   "tableName",
				Value: table.TableName,
			}}, bson.D{
				{"$set", bson.D{{
					Key:   "updatedAt",
					Value: now.Format("2006-01-02 15:04:05"),
				}}},
			})
			CheckErr(err)
		}
	case "recycle":
		if checkTime(table.UpdatedAt, now, spanDay) {
			// 超过时间间隔，更新

		}
	}
}

// 循环查询是否更新
func LoopQueryUpdate() {
	go func() {
		var updates []Update
		for {
			RbMQ.ConnCheck()
			// 根据规则过滤数据，这里过滤条件为空
			data, err := Db.Collection("update").Find(context.TODO(), bson.D{})
			CheckErr(err)
			// 解析数据到数组中
			err = data.All(context.TODO(), &updates)
			CheckErr(err)
			now := time.Now()
			for _, v := range updates {
				checkUpdate(v, now)
			}
			_ = data.Close(context.TODO())
			// 睡眠一分钟释放cpu
			time.Sleep(time.Minute * time.Duration(sleepTime))
		}
	}()
}

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s\n", msg, err)
	}
}
