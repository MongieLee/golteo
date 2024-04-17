package service

import (
	"fmt"
	"ginl/config"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMq struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string //队列名称
	Exchange  string // 交换机名称
	key       string // bind key
	Mqurl     string // 链接信息
}

func NewRabbitMq(queueName, exchange, key string) *RabbitMq {
	mqUrl := fmt.Sprintf("amqp://%v:%v@%v:%v/%v", config.CustomConfig.Rmq.User, config.CustomConfig.Rmq.Password, config.CustomConfig.Rmq.Host, config.CustomConfig.Rmq.Port, config.CustomConfig.Rmq.VirtualHost)
	return &RabbitMq{
		QueueName: queueName,
		Exchange:  exchange,
		key:       key,
		Mqurl:     mqUrl,
	}
}

func NewRabbitMqSimple(queueName string) *RabbitMq {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	rabbitMq := NewRabbitMq(queueName, "", "")
	var err error
	rabbitMq.conn, err = amqp.Dial(rabbitMq.Mqurl)
	if err != nil {
		rabbitMq.FailOnErr(err, "failed to connect rabbitmq")
	}
	rabbitMq.channel, err = rabbitMq.conn.Channel()
	if err != nil {
		rabbitMq.FailOnErr(err, "failed to open channel")
	}
	return rabbitMq
}

func (r *RabbitMq) Destory() {
	r.conn.Close()
	r.channel.Close()
}

func (r *RabbitMq) FailOnErr(err error, message string) {
	if err != nil {
		//log.Fatalf("%s,%s", message, err)
		log.Printf("%s,%s\n", message, err)
		//panic(fmt.Sprintf("%s,%s", message, err))
	}
}

func (r *RabbitMq) PublishSimple(message string) {
	if r == nil {
		return
	}
	// 申请队列，如果不存在则创建，存在则跳过
	_, err := r.channel.QueueDeclare(r.QueueName,
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞处理
		false,
		// 额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		// 如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		// 如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

func (r *RabbitMq) ConsumeSimple() {
	q, err := r.channel.QueueDeclare(r.QueueName,
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞处理
		false,
		// 额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	consume, err := r.channel.Consume(
		q.Name,
		// 用来区分多个消费者
		"",
		// 是否自动应答
		true,
		// 是否独有
		false,
		// true，表示不能将同一个Connection中生产者发送的消息传递给这个Connection中 的消费
		false,
		// 是否阻塞
		false, nil)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		for d := range consume {
			log.Printf("received a message: %s", d.Body)
		}
	}()
	//ch := make(chan bool)
	//<-ch
}

func InitAndSenDMsgRmq() {
	rmq := NewRabbitMqSimple("testFuck")
	rmq.PublishSimple("Hello testFuck!")
}
