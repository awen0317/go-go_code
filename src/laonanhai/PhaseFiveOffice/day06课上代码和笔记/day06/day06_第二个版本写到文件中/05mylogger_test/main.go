package main

import (
	"time"

	"code.oldboyedu.com/day06/mylogger"
)

// 测试我们自己写 的日志库
func main() {
	// log := mylogger.NewLog("Info")
	log := mylogger.NewFileLogger("Info", "./", "zhoulinwan.log", 10*1024*1024)
	for {

		log.Debug("这是一条Debug日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条warning日志")
		id := 10010
		name := "理想"
		log.Error("这是一条Error日志,id:%d,name:%s", id, name)
		log.Fatal("这是一条Fatal日志")
		time.Sleep(2 * time.Second)
	}
}
