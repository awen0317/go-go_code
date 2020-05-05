package conf

type AppConf struct {
	KafkaConf `ini:"06kafka_demo"`
	TaillogConf `ini:"taillog"`
}
type KafkaConf struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}


type TaillogConf struct {
	FileName string `ini:"filename"`
}