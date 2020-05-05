package main

import (
	"fmt"
	"imooc-rabbitmq/RabbitMq"
	"strconv"
	"time"
)

func main()  {
	rabbitmq := RabbitMq.NewRabbitNQSimple(""+"imoocSimple")
	for i:=0;i<=100;i++{
		rabbitmq.PublishSimple("Hello imooc"+strconv.Itoa(i))
		time.Sleep(1*time.Second)
		fmt.Println(i)

	}
}
