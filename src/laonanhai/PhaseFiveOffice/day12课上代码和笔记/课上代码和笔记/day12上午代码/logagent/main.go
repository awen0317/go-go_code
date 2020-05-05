package main

import (
	"code.oldboyedu.com/logagent/conf"
	"code.oldboyedu.com/logagent/etcd"
	"code.oldboyedu.com/logagent/kafka"
	"fmt"
	"gopkg.in/ini.v1"
	"time"
)

// logAgent入口程序

var (
	cfg = new(conf.AppConf)
)

//func run(){
//	// 1. 读取日志
//	for {
//		select {
//			case line := <- taillog.ReadChan():
//				// 2. 发送到kafka
//				06kafka_demo.SendToKafka(cfg.KafkaConf.Topic, line.Text)
//		default:
//			time.Sleep(time.Second)
//		}
//	}
//}

func main(){
	// 0. 加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}

	// 1. 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init 06kafka_demo success.")
	// 2. 初始化ETCD
	// 5 * time.Second
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("init etcd success.")
	// 2.1 从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf("/xxx")
	// 2.2 拍一个哨兵去监视日志收集项的变化（有变化及时通知我的logAgent实现热加载配置）
	if err != nil {
		fmt.Printf("etcd.GetConf failed,err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success, %v\n", logEntryConf)
	for index, value := range logEntryConf{
		fmt.Printf("index:%v value:%v\n", index, value)
	}

	// 2. 打开日志文件准备收集日志
	//err = taillog.Init(cfg.TaillogConf.FileName)
	//if err != nil {
	//	fmt.Printf("Init taillog failed,err:%v\n", err)
	//	return
	//}
	//fmt.Println("init taillog success.")
	// 3. 具体的业务
	//run()

}
