package model

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

var RbMQ RabbitMQ

type Message struct {
	DataType string `json:"dataType"` // 消息类型，后端根据类型进行对应处理
	Data     []byte `json:"data"`     // 消息内容，为保存json的字节数组
}

func InitMq() {
	RbMQ = RabbitMQ{}
	conn, err := amqp.Dial("amqp://admin:admin@43.138.78.252:5672//")
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer Conn.Close()
	// 创建一个channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	RbMQ.Conn = conn
	RbMQ.Ch = ch
	//defer Ch.Close()

	// 声明一个队列
	_, err = RbMQ.Ch.QueueDeclare(
		"trash", // 队列名称
		false,   // 是否持久化
		false,   // 是否自动删除
		false,   // 是否独立
		false, nil,
	)
	failOnError(err, "Failed to declare a queue")
}

func (c *RabbitMQ) ConnCheck() {
	if c.Ch.IsClose() || c.Conn.IsClosed() {
		log.Println("rabbitmq连接断开，重新连接")
		_ = c.Conn.Close()
		_ = c.Ch.Close()
		InitMq()
	}
}

// 帮助函数检测每一个amqp调用
func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s\n", msg, err)
	}
}
