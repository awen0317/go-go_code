package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

type LogConfig struct {
	KafkaAddr  string
	ESAddr     string
	LogPath    string
	LogLevel   string
	KafkaTopic string
}

var (
	logConfig *LogConfig
)

func initConfig(confType string, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	logConfig = &LogConfig{}
	logConfig.LogLevel = conf.String("logs::log_level")
	if len(logConfig.LogLevel) == 0 {
		logConfig.LogLevel = "debug"
	}

	logConfig.LogPath = conf.String("logs::log_path")
	if len(logConfig.LogPath) == 0 {
		logConfig.LogPath = "./logs"
	}

	logConfig.KafkaAddr = conf.String("06kafka_demo::server_addr")
	if len(logConfig.KafkaAddr) == 0 {
		err = fmt.Errorf("invalid 06kafka_demo addr")
		return
	}

	logConfig.ESAddr = conf.String("es::addr")
	if len(logConfig.ESAddr) == 0 {
		err = fmt.Errorf("invalid es addr")
		return
	}

	logConfig.KafkaTopic = conf.String("06kafka_demo::topic")
	if len(logConfig.ESAddr) == 0 {
		err = fmt.Errorf("invalid es addr")
		return
	}
	return
}
