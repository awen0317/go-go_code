package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main()  {
	config :=sarama.NewConfig()
	//tailf包的使用
	config.Producer.RequiredAcks =sarama.WaitForAll//发送完数据需要确认leader和follow确认
	//选择分区的方式
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出数据需要，轮询方式
	config.Producer.Return.Successes =true //成功交付的消息将在success channel返回
	//构造一个消息
	msg :=&sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")
	//链接kafka
	client,err :=sarama.NewSyncProducer([]string{"127.0.0.1:9092"},config)
	if err != nil {
		fmt.Println("product closed,err",err)
		return
	}
	defer client.Close()
	//发送消息
	pid,offset,err :=client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg faild err:",err)
	}
	fmt.Printf("pid:%v offset:%v\n",pid,offset)
}