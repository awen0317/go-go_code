package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"time"
)

//etcd watch
func main()  {

	cli,err :=clientv3.New(clientv3.Config{
		Endpoints:            []string{"127.0.0.1:2379"},
		DialTimeout:          5*time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// ctx ,cancel :=context.WithTimeout(context.Background(),time.Second)
	//	拍一个哨兵，监视lmh变化
	rch :=cli.Watch(context.Background(),"lmh")
	// cancel()
	//尝试取值
	//从通道尝试取值(监视的信息)
	fmt.Println("waiting...")
	for wresp :=range rch{
		for _,evt :=range wresp.Events{
			fmt.Printf("Type:%v key:%v value:%v",evt.Type,string(evt.Kv.Key),string(evt.Kv.Value))
		}
	}
}
