package config

type AppConf struct {
	KafkaConf   `ini:"kafka"`
	EtcdConf    `ini:"etcd"`
	TaillogConf `ini:"taillog"`
}

type KafkaConf struct {
	Address     string `ini:"address"`
	Top         string `ini:"topic"`
	ChanMaxSize int    `ini:"chan_max_size"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key     string `ini:"collection_log_key"`
}

type TaillogConf struct {
	FileName string `ini:"fileName"`
}
