package conf

// DriverName : driver name 
const DriverName = "mysql"

// DbConfig : 数据库结构
type DbConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	Database  string
	IsRunning bool
}

// DbMasterList : Mysql列表
var DbMasterList = []DbConfig{
	{
		Host:      "192.168.74.121",
		Port:      3306,
		User:      "root",
		Pwd:       "123456",
		Database:  "lottery",
		IsRunning: true,
	},
}

// DbMaster 实例化数据库
var DbMaster DbConfig = DbMasterList[0]
