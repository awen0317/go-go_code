package taillog

import (
	"fmt"
	"laonanhai/PhaseFive/day11/logagent/etcd"
	"time"
)
var (
	tskMgr *tailLogMgr
)
type tailLogMgr struct {
	LogEntry []*etcd.LogEntry
	tskMap map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		LogEntry: logEntryConf,
		tskMap:make(map[string]*TailTask,16),
		newConfChan:make(chan []*etcd.LogEntry),//无缓冲区的通道

	}
	for _,logEntry :=range logEntryConf{
		//初始化的时候起了多少哥tailtask
		tailObj :=NewTailTask(logEntry.Path,logEntry.Topic)
		mk :=fmt.Sprintf("%s_%s",logEntry.Path,logEntry.Topic)
		tskMgr.tskMap[mk] = tailObj
	}
	go tskMgr.run()
}

//监听自己的通道，有了新的配置过来之后就对影处理
//1配置新增
//2配置删除
//3配置变更
func (t *tailLogMgr)run()  {
	for true {
		select {
		case newConf :=<-t.newConfChan:
			for _,conf :=range newConf{
				mk :=fmt.Sprintf("%s_%s",conf.Path,conf.Topic)
				_,ok :=t.tskMap[mk]
				if ok{
					continue
				}else{
					//新增额
					tailObj :=NewTailTask(conf.Path,conf.Topic)
					t.tskMap[mk] = tailObj
				}
			}
			//找出原来t.logEntry有，但是newConf中没有，要删除
			for _,c1 :=range t.LogEntry{//从最开始配置中依次拿出配置项
				isDelete :=false
				for _,c2 :=range newConf{//去新的配置中逐一匹配进行比较
					if c2.Path ==c1.Path && c2.Topic ==c1.Topic{
						isDelete =false
						continue
					}
				}
				if isDelete{
					//把c1对应的这个tailObj给停掉
					mk :=fmt.Sprintf("%s_%s",c1.Path,c1.Topic)
					t.tskMap[mk].cancelFun()

				}

			}
			//1配置新增

			//2配置删除
			//3配置变更
			fmt.Println("新配置来了",newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}
//通过一个函数，向外暴漏tskMgr的newChanConf
func NewConfChan() chan<- []*etcd.LogEntry{
	return tskMgr.newConfChan
}
