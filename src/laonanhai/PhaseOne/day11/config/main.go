package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main()  {

	conf,err := config.NewConfig("ini","./logagent.log")
	if err != nil {
		fmt.Println("new config failed err",err)
		return
	}
	port,err :=conf.Int("server::port")
	if err != nil {
		fmt.Println("new config failed err",err)
		return
	}
	fmt.Println("Post:",port)
	log_level :=conf.String("log::log_level")
	if err != nil {
		fmt.Println("log::log_level err",err)
		return
	}
	fmt.Println("log_level:",log_level)
	log_path :=conf.String("log::log_path")
	if err != nil {
		fmt.Println("log::log_path err",err)
		return
	}
	fmt.Println("log_path:",log_path)
}
