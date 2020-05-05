package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	cli *clientv3.Client
)
type LogEntry struct {
	Path string `json:"path"`
	Topic string `json:"topic"`
}

func Init(addr string,timeout time.Duration)(err error)  {
	cli,err =clientv3.New(clientv3.Config{
		Endpoints:            []string{addr},
		DialTimeout:          timeout,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	return
}
//从etcd中获取配置项
func GetConf(key string) (logEntryConf []*LogEntry,err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	// fmt.Println(resp)
	// os.Exit(1)
	for _, ev :=range resp.Kvs{
		// fmt.Printf("key:%s=>value:%s",ev.Key,ev.Value)
		err =json.Unmarshal(ev.Value,&logEntryConf)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed,err:%v\n",err)
			return
		}
	}
	return
}

//etcd watch
func WatchConf(key string,newConfCh chan<- []*LogEntry)  {
	rch :=cli.Watch(context.Background(),key)
	// cancel()
	//尝试取值
	//从通道尝试取值(监视的信息)
	for wresp :=range rch{
		for _,evt :=range wresp.Events{
			fmt.Printf("Type:%v key:%v value:%v",evt.Type,string(evt.Kv.Key),string(evt.Kv.Value))
			//通知别人
			//先判断操作的类型
			var newConf []*LogEntry
			if evt.Type != clientv3.EventTypeDelete{
				//如果是删除操作
				err :=json.Unmarshal(evt.Kv.Value,&newConf)
				if err != nil {
					fmt.Println("unmarshal failed,err :%v",err)
					continue
				}
			}
			fmt.Printf("get new conf :%v\n",newConf)
			//slice类型
			newConfCh<-newConf
		}
	}
	
}