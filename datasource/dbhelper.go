package datasource

import(
	"xorm.io/xorm"
	_ "github.com/go-sql-driver/mysql"
	"golang-lottery/conf"
	"fmt"
	"log"
	"sync"
)

var dbLock sync.Mutex
var masterInstance *xorm.Engine

func InstanceDbMaster() *xorm.Engine{
	if masterInstance != nil {
		return masterInstance
	}

	dbLock.Lock()
	defer dbLock.Unlock()

	// 如果过个请求过来，第一个已经创建。但是并发情况下
	// 可能会多次创建。所以在下面再次判断是否已经创建masterInstance
	if masterInstance != nil {
		return masterInstance
	}

	return NewDbMaster()
}

func NewDbMaster() *xorm.Engine {
	sourcename := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
							conf.DbMaster.User,
							conf.DbMaster.Pwd,
							conf.DbMaster.Host,
							conf.DbMaster.Port,
							conf.DbMaster.Database)
	instance, err := xorm.NewEngine(conf.DriverName, sourcename)
	if err != nil {
		log.Fatal("dbhelper.NewDbMaster NewEngine error", err)
		return nil
	}

	instance.ShowSQL(true)

	masterInstance = instance
	return instance
}