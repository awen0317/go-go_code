package main

import (
	"imooc-rabbitmq/RabbitMq"
)

func main()  {

	rabbitmq := RabbitMq.NewRabbitMQPubSub("NewProduct")
	rabbitmq.RevieveSub()
}
