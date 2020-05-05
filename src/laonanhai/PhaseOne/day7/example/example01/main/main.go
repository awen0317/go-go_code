package main

import (
	"fmt"
	"laonanhai/PhaseOne/day7/example/example01/balance"
	"math/rand"
	"os"
	"time"
)

func main() {
	//insts := make([]*balance.Instance)
	var insts []*balance.Instance
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8000)
		//append对host自动扩容的额
		insts = append(insts, one)
	}
	//var balancer balance.Balancer
	var balanceName = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}
	//if conf == "random" {
	//	balancer = &balance.RandomBalance{}
	//	fmt.Println("use radndom")
	//} else if conf == "roundrobin" {
	//	balancer = &balance.RoundRobinBalance{}
	//	fmt.Println("use roundrobin")
	//}
	for {
		//inst, err := balancer.DoBalance(insts)
		inst, err := balance.DoBalance(balanceName,insts)
		if err != nil {
			fmt.Println("do balance err", err)
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
