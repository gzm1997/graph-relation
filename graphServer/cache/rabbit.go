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
var queue amqp.Queue

const EXCHANGE = "cayleyExchange"

func init() {
	rabbitUser := beego.AppConfig.String("rabbitUser")
	rabbitPsw := beego.AppConfig.String("rabbitPsw")
	rabbitIp := beego.AppConfig.String("rabbitIp")
	rabbitPort, _ := beego.AppConfig.Int("rabbitPort")
	dbUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", rabbitUser, rabbitPsw, rabbitIp, rabbitPort)
	//fmt.Println("dbUrl", dbUrl)
	var err error
	if conn, err = amqp.Dial(dbUrl); err != nil {
		panic(err)
	}
	//defer conn.Close()
	if ch, err = conn.Channel(); err != nil {
		panic(err)
	}
	//defer ch.Close()
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
	//fmt.Println("laalllalaall")
}

func PublishMsg(json []byte, bingdingKey BindingKey) error {
	err := ch.Publish(
		EXCHANGE,
		bingdingKey.String(),
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "application/json",
			Body: json,
		})
	return err
}

func GetMsg() {
	defer conn.Close()
	defer ch.Close()
	forever := make(chan bool)
	for _, bingdingKey := range AllBindingKeys {
		go func(b BindingKey) {
			fmt.Println("queue for", b)
			q, err := ch.QueueDeclare(
				"",
				false,
				false,
				true,
				false,
				nil,
			)
			if err != nil {
				panic(err)
			}

			err = ch.QueueBind(
				q.Name,
				b.String(),
				EXCHANGE,
				false,
				nil)
			if err != nil {
				panic(err)
			}

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