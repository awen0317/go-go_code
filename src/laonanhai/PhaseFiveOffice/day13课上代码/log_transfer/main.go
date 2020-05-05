package main

import (
	"fmt"

	"code.oldboy.com/log_transfer/conf"
	"code.oldboy.com/log_transfer/es"
	"code.oldboy.com/log_transfer/kafka"
	"gopkg.in/ini.v1"
)

// log transfer
// 将日志数据从kafka取出来发往ES

func main() {
	// 0. 加载配置文件
	var cfg = new(conf.LogTransferCfg)
	err := ini.MapTo(cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Printf("init config failed, err:%v\n", err)
		return
	}
	fmt.Printf("cfg:%v\n", cfg)

	// 1. 初始化ES
	// 1.1 初始化一个ES连接的client
	// 1.2 对外提供一个往ES写入数据 的一个函数
	err = es.Init(cfg.ESCfg.Address, cfg.ESCfg.ChanSize, cfg.ESCfg.Nums)
	if err != nil {
		fmt.Printf("init ES client failed,err:%v\n", err)
		return
	}
	fmt.Println("init es success.")

	// 2. 初始化kafka
	// 2.1 连接kafka, 创建分区的消费者
	// 2.2 每个分区的消费者分别取出数据 通过SendToES()将数据发往ES
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("init 06kafka_demo consumer failed,err:%v\n", err)
		return
	}
	select {}

}
