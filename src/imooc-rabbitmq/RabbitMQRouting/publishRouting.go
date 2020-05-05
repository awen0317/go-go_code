package main

import (
	"fmt"
	"imooc-rabbitmq/RabbitMq"
	"strconv"
	"time"
)

func main()  {
	imoocOne :=RabbitMq.NewRabbitMQRouting("exImooc","imooc_one")
	imoocTwo :=RabbitMq.NewRabbitMQRouting("exImooc","imooc_two")

	for i:=0;i<10;i++{
		imoocOne.PublishRouting("main imcco one"+strconv.Itoa(i))
		imoocTwo.PublishRouting("main imcco two"+strconv.Itoa(i))
		time.Sleep(1*time.Second)
		fmt.Println(i)
	}
}
