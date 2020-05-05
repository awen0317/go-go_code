package main

import (
	"imooc-rabbitmq/RabbitMq"
)

func main()  {
	rabbitmq := RabbitMq.NewRabbitNQSimple("imoocSimple")
	rabbitmq.ConsumerSimple()
}
