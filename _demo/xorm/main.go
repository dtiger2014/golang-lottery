/**
 * 应用程序
 * 同目录下多文件引用的问题解决方法：
 * https://blog.csdn.net/pingD/article/details/79143235
 * 方法1 1 go build ./ 2 运行编译后的文件
 * 方法2 go run *.go
 */
package main

import (
	"fmt"
	"log"
	"time"

	"golang-lottery/_demo/xorm/models"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

const (
	DriverName           = "mysql"
	MasterDataSourceName = "root:123456@tcp(127.0.0.1:3306)/test_golang?charset=utf8"
)

var engine *xorm.Engine

func main() {
	engine = newEngine()

	// query()
	// execute()
	// ormInsert()
	// ormGet()
	// ormGetCols()
	// ormCount()
	// ormFindRows()
	// ormUpdate()
	// ormOmitUpdate()
	ormMustColsUpdate()
}

// 连接到数据库
func newEngine() *xorm.Engine {
	engine, err := xorm.NewEngine(DriverName, MasterDataSourceName)

	if err != nil {
		log.Fatal(newEngine, err)
		return nil
	}

	engine.ShowSQL(true)
	return engine
}

// 通过query方法查询
func query() {
	sql := "SELECT * FROM user_info"
	results, err := engine.QueryString(sql)
	if err != nil {
		log.Fatal("query", sql, err)
		return
	}
	total := len(results)
	if total == 0 {
		fmt.Println("没有数据", sql)
	} else {
		for i, data := range results {
			fmt.Printf("%d = %+v\n", i, data)
		}
	}
}

// 通过execute方法执行更新
func execute() {
	sql := `INSERT INTO user_info values(NULL, 'fanfan3',0,0)`
	affected, err := engine.Exec(sql)
	if err != nil {
		log.Fatal("execute error", err)
	}

	id, err := affected.LastInsertId()
	rows, err := affected.RowsAffected()

	fmt.Println("execute id=", id, ", rows=", rows)
}

// 根据models的结构映射数据表
func ormInsert() {
	uinfo := &models.UserInfo{
		Id:         0,
		Name:       "饭饭",
		SysCreated: 0,
		SysUpdated: 0,
	}
	id, err := engine.Insert(uinfo)
	if err != nil {
		log.Fatal("ormInsert error", err)
	}
	fmt.Println("ormInsert id=", id)
	fmt.Printf("%v\n", *uinfo)
}

// 根据models的结构读取数据
func ormGet() {
	uinfo := &models.UserInfo{Id: 2}
	ok, err := engine.Get(uinfo)
	if ok {
		fmt.Printf("%+v\n", *uinfo)
	} else if err != nil {
		log.Fatal("ormGet error", err)
	} else {
		fmt.Println("ormGet empty id=", uinfo.Id)
	}
}

// 获取指定的字段
func ormGetCols() {
	UserInfo := &models.UserInfo{Id: 2}
	ok, err := engine.Cols("name,sys_created").Get(UserInfo)
	if ok {
		fmt.Printf("%+v\n", UserInfo)
	} else if err != nil {
		log.Fatal("ormGetCols error", err)
	} else {
		fmt.Println("ormGetCols empty id=2")
	}
}

// 统计
func ormCount() {
	//count, err := engine.Count(&UserInfo{})
	//count, err := engine.Where("name_zh=?", "饭饭").Count(&UserInfo{})
	count, err := engine.Count(&models.UserInfo{Name: "饭饭"})
	if err == nil {
		fmt.Printf("count=%v\n", count)
	} else {
		log.Fatal("ormCount error", err)
	}
}

// 查找多行数据
func ormFindRows() {
	list := make([]models.UserInfo, 0)
	//list := make(map[int]UserInfo)
	//err := engine.Find(&list)
	//err := engine.Where("id>?", 1).Limit(100, 0).Find(&list)
	err := engine.Cols("id", "name").Where("id>?", 0).
		Limit(10).Asc("id", "sys_created").Find(&list)

	//list := make([]map[string]string, 0)
	//err := engine.Table("star_info").Cols("id", "name_zh", "name_en").
	// Where("id>?", 1).Find(&list)

	if err == nil {
		fmt.Printf("%v\n", list)
	} else {
		log.Fatal("ormFindRows error", err)
	}
}

// 更新一个数据
func ormUpdate() {
	// 全部更新
	//UserInfo := &UserInfo{NameZh:"测试名"}
	//ok, err := engine.Update(UserInfo)
	// 指定ID更新
	UserInfo := &models.UserInfo{Name: "梅西"}
	ok, err := engine.ID(2).Update(UserInfo)
	fmt.Println(ok, err)
}

// 排除某字段
func ormOmitUpdate() {
	info := &models.UserInfo{Id: 1}
	ok, _ := engine.Get(info)
	if ok {
		if info.SysCreated > 0 {
			ok, _ := engine.ID(info.Id).Omit("sys_created").
				Update(&models.UserInfo{
					SysCreated: 0,
					SysUpdated: int(time.Now().Unix()),
				})
			fmt.Printf("ormOmitUpdate, rows=%d, sys_created=%d\n", ok, 0)
		} else {
			ok, _ := engine.ID(info.Id).Omit("sys_created").
				Update(&models.UserInfo{
					SysCreated: 1,
					SysUpdated: int(time.Now().Unix()),
				})
			fmt.Printf("ormOmitUpdate, rows=%d, sys_created=%d\n", ok, 0)
		}
	}
}

// 字段为空也可以更新（0, 空字符串等）
func ormMustColsUpdate() {
	info := &models.UserInfo{Id: 11}
	ok, _ := engine.Get(info)
	if ok {
		if info.SysCreated > 0 {
			ok, _ := engine.ID(info.Id).
				MustCols("sys_created").
				Update(&models.UserInfo{
					SysCreated: 0,
					SysUpdated: int(time.Now().Unix()),
				})
			fmt.Printf("ormMustColsUpdate, rows=%d, sys_created=%d\n", ok, 0)
		} else {
			ok, _ := engine.ID(info.Id).
				MustCols("sys_created").
				Update(&models.UserInfo{
					SysCreated: 1,
					SysUpdated: int(time.Now().Unix()),
				})
			fmt.Printf("ormMustColsUpdate, rows=%d, sys_created=%d\n", ok, 0)
		}
	}
}
