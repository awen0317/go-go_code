package main

import (
	"imooc-rabbitmq/RabbitMq"
)

func main()  {
	imoocTwo :=RabbitMq.NewRabbitMQRouting("exImooc","imooc_two")
	imoocTwo.RecieveRouting()
}
