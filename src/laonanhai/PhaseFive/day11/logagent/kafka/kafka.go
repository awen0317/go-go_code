package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

type logData struct {
	topic string
	data  string
}

// 专门往kafka写日志
var (
	client      sarama.SyncProducer
	logDataChan chan *logData
)

func Init(addr []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	// tailf包的使用
	config.Producer.RequiredAcks = sarama.WaitForAll // 发送完数据需要确认leader和follow确认
	// 选择分区的方式
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出数据需要，轮询方式
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	// 链接kafka
	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("product closed,err", err)
		return
	}

	logDataChan = make(chan *logData, maxSize)
	go sendTiKafka()
	return
}


func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}
func sendTiKafka() {
	for {
		select {
		case ld:=<-logDataChan :
			// 构造一个消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			// 发送消息
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg faild err:", err)
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(50*time.Millisecond)
		}

	}

}
