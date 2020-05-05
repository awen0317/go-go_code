package main

import (
	"fmt"
	"imooc-rabbitmq/RabbitMq"
	"strconv"
	"time"
)

func main()  {
	rabbitmq := RabbitMq.NewRabbitMQPubSub("NewProduct")
	for i:=0;i<=10;i++ {
		rabbitmq.PublishPub("Hello imooc"+strconv.Itoa(i))
		time.Sleep(1*time.Second)
		fmt.Println("Hello imooc"+strconv.Itoa(i))
	}
	fmt.Println("发送成功")
}
