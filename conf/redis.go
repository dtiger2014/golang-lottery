package conf

// RdsConfig Redis配置信息结构
type RdsConfig struct {
	Host string
	Port int
	User string
	Pwd string
	IsRunning bool
}

// RdsCacheList Redis配置信息
var RdsCacheList = []RdsConfig{
	{
		Host: "192.168.74.121",
		Port: 6379,
		User: "",
		Pwd: "",
		IsRunning: true,
	},
}

// RdsCache 实例化Redis缓存
var RdsCache RdsConfig = RdsCacheList[0]