package config

type AppConf struct {
	KafkaConf `ini:"kafka"`
	// TaillogConf `ini:"taillog"`
	EtcdConf `ini:"etcd"`
}
type KafkaConf struct {
	Addr string `ini:"address"`
	// Topic string `ini:"topic"`
	ChanMaxSize int `ini:"chan_max_size"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Key string `ini:"collect_log_key"`
	Timeout int `ini:"timeout"`
}
// type TaillogConf struct {
// 	Filename string  `ini:"filename"`
// }