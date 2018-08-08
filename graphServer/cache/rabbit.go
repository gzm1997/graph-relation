package cache

import (
	"github.com/streadway/amqp"
	"github.com/astaxie/beego"
	"fmt"
	"relation-graph/graphRelation/graphServer/models"
	"encoding/json"
	"relation-graph/graphRelation/createTriple/session"
)

var conn *amqp.Connection
var ch *amqp.Channel


const EXCHANGE = "cayleyExchange"

func init() {
	//加载rabbitMQ在配置文件中的配置
	rabbitUser := beego.AppConfig.String("rabbitUser")
	rabbitPsw := beego.AppConfig.String("rabbitPsw")
	rabbitIp := beego.AppConfig.String("rabbitIp")
	rabbitPort, _ := beego.AppConfig.Int("rabbitPort")
	dbUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", rabbitUser, rabbitPsw, rabbitIp, rabbitPort)
	var err error
	//建立链接
	if conn, err = amqp.Dial(dbUrl); err != nil {
		panic(err)
	}
	//声明一个channel
	if ch, err = conn.Channel(); err != nil {
		panic(err)
	}
	//声明一个类型为direct的交换区
	err = ch.ExchangeDeclare(
		EXCHANGE,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(nil)
	}
}

func PublishMsg(json []byte, bingdingKey BindingKey) error {
	err := ch.Publish(
		//指定我们要使用的direct类型的交换区
		EXCHANGE,
		//根据bingdingKey进行转发
		bingdingKey.String(),
		false,
		false,
		amqp.Publishing{
			//发送的信息时候持久的 即使发送方突然挂了 重启之后还会继续发
			DeliveryMode: amqp.Persistent,
			//发送的类型是json对象
			ContentType: "application/json",
			Body: json,
		})
	return err
}

func GetMsg() {
	//在这里释放rabbit的资源
	defer conn.Close()
	defer ch.Close()
	forever := make(chan bool)
	for _, bingdingKey := range AllBindingKeys {
		go func(b BindingKey) {
			fmt.Println("queue for", b)
			//声明一个匿名队列
			q, err := ch.QueueDeclare(
				"",
				true,
				false,
				true,
				false,
				nil,
			)
			if err != nil {
				panic(err)
			}
			//进行队列绑定
			err = ch.QueueBind(
				q.Name,
				b.String(),
				EXCHANGE,
				false,
				nil)
			if err != nil {
				panic(err)
			}
			//声明一个消费者 从这个匿名队列里面读取信息
			msgs, err := ch.Consume(
				q.Name, // queue
				"",     // consumer
				false,   // auto ack
				false,  // exclusive
				false,  // no local
				false,  // no wait
				nil,    // args
			)
			if err != nil {
				panic(err)
			}

			store := session.GetGraph()

			switch b {
			case CreateGroupShareLink:
				for d := range msgs {
					cgsl := models.CreateGroupShareLink{}
					if err := json.Unmarshal(d.Body, &cgsl); err == nil {
						fmt.Println("add CreateGroupShareLink", cgsl)
						fmt.Println(cgsl.AddCreateGroupShareLinkToCayley(store))
					} else {
						panic(err)
					}
					d.Ack(false)
				}
			case ClickGroupShareLink:
				for d := range msgs {
					cgsl := models.ClickGroupShareLink{}
					if err := json.Unmarshal(d.Body, &cgsl); err == nil {
						fmt.Println("add CreateGroupShareLink", cgsl)
						cgsl.AddClickGroupShareLinkToCayley(store)
					} else {
						panic(err)
					}
					d.Ack(false)
				}
			case CreateFileLink:
				for d := range msgs {
					cfl := models.CreateFileLink{}
					if err := json.Unmarshal(d.Body, &cfl); err == nil {
						fmt.Println("add CreateFileLink", cfl)
						cfl.AddCreateFileLinkToCayley(store)
					} else {
						panic(err)
					}
					d.Ack(false)
				}
			case ClickFileLink:
				for d := range msgs {
					cfl := models.ClickFileLink{}
					if err := json.Unmarshal(d.Body, &cfl); err == nil {
						fmt.Println("add ClickFileLink", cfl)
						cfl.AddClickFileLinkToCayley(store)
					} else {
						panic(err)
					}
					d.Ack(false)
				}
			case User:
				for d := range msgs {
					user := models.User{}
					if err := json.Unmarshal(d.Body, &user); err == nil {
						fmt.Println("add user", user)
						user.AddUserToCayley(store)
					} else {
						panic(err)
					}
					d.Ack(false)
				}
			case File:
				for d := range msgs {
					file := models.File{}
					if err := json.Unmarshal(d.Body, &file); err == nil {
						fmt.Println("add file", file)
						file.AddFileToCayley(store)
					} else {
						panic(err)
					}
					d.Ack(false)
				}
			case Group:
				for d := range msgs {
					group := models.Group{}
					if err := json.Unmarshal(d.Body, &group); err == nil {
						fmt.Println("add group", group)
						group.AddGroupToCayley(store)
					} else {
						panic(err)
					}
					d.Ack(false)
				}
			}
		}(bingdingKey)
	}
	<- forever
}