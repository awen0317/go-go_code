package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"gopkg.in/ini.v1"
	"laonanhai/PhaseFive/day11/logagent/config"
	"laonanhai/PhaseFive/day11/logagent/etcd"
	"laonanhai/PhaseFive/day11/logagent/kafka"
	"laonanhai/PhaseFive/day11/logagent/taillog"
	"laonanhai/PhaseFive/day11/logagent/utils"
	"sync"
	"time"
)

// func run() {
// 	select {
// 	case line := <-taillog.ReadChan():
// 		kafka.SendTiKafka(cfg.KafkaConf.Topic, line.Text)
// 	default:
// 		time.Sleep(time.Second)
// 	}
//
// }
func ReadLog() {
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tailObj.Lines
		fmt.Println("line:", line.Text)
		fmt.Println("ok:", ok)
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tailObj.Filename)
			time.Sleep(time.Second)
			continue
		}

	}
}

// logagent入口

var (
	tailObj *tail.Tail
	cfg     = new(config.AppConf)
)

func main() {
	// 1.初始化配置文件
	err := ini.MapTo(cfg, "./config/config.ini")
	if err != nil {
		fmt.Println("init config failed", err)
		return
	}
	// 2.初始化一哥kafka的连接
	err = kafka.Init([]string{cfg.KafkaConf.Addr},cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Println("init kafka failed,ere \n ", err)
		return
	}
	fmt.Println("init kafka success")
	// 3.初始化一哥etcd的连接go mod
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Println("init etcd failed,err \n ", err)
		return
	}
	fmt.Println("init etcd success")
	//为了实现每一个logangent都拉去自己独有的配置，所以要以自己的IP地址作为区分
	ipStr,err :=utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey :=fmt.Sprintf(cfg.EtcdConf.Key,ipStr)
	//4.从etcd中获取日志收集项的信息
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Println("init GetConf failed,err \n ", err)
		return
	}
	fmt.Println("init GetConf form etcd success %v\n", logEntryConf)
	//遍历展示获取到的配置信息
	for index, value := range logEntryConf {
		fmt.Printf("index:%v,value:%s\n", index,value)
	}

	//5.收集日志发往kafka
	taillog.Init(logEntryConf)

	//2.2派一个哨兵去监视收集项的变化
	//因为newConfChan 访问了tskMgr的newConfChan，这个channel是taillog。Init（logEntryConf）执行的处初始化

	newConfChan :=taillog.NewConfChan()//从tailog包中获取对外暴漏的通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(cfg.EtcdConf.Key,newConfChan)//哨兵发现最新的配置信息会通知上面的通道
	wg.Wait()
	// 3.1收集日志发往kafka
	// 3.2循环每一个日志收集项
	// for _, logEntry := range logEntryConf {
	// 	tsk :=taillog.NewTailTask(logEntry.Path,logEntry.Topic)
	// 	// 从tailOBj。Lines的通道中
	// 	for {
	// 		select {
	// 		case line := <-tsk.ReadChan():
	// 			// 3.3发往kafka
	// 			kafka.SendTiKafka(logEntry.Topic,line.Text)
	//
	// 		}
	// 	}

	// }

	// 2.打开文件准备手机
	// err = taillog.Init(cfg.TaillogConf.Filename)
	// if err != nil {
	// 	fmt.Printf("init taillog error err", err)
	// }
	// fmt.Println("init taill success")
	//
	// run()
}
