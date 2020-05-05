package conf

type AppConf struct {
	KafkaConf `ini:"06kafka_demo"`
	EtcdConf `ini:"etcd"`

}
type KafkaConf struct {
	Address string `ini:"address"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int `ini:"timeout"`
}


// ---- unused ↓-----------

type TaillogConf struct {
	FileName string `ini:"filename"`
}