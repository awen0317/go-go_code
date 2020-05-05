package main

import (
	"imooc-rabbitmq/RabbitMq"
)

func main()  {
	imoocOne :=RabbitMq.NewRabbitMQRouting("exImooc","imooc_one")
	imoocOne.RecieveRouting()
}