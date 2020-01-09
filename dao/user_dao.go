package dao

import (
	"log"
	"xorm.io/xorm"
	"golang-lottery/models"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao {
		engine: engine,
	}
}

func (d *UserDao) Get(id int) *models.LtUser {
	data := &models.LtUser{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *UserDao) GetAll() []models.LtUser {
	datalist := make([]models.LtUser, 0)
	err := d.engine.
	Asc("sys_status").
	Asc("displayorder").
	Find(&datalist)
	if err != nil {
		log.Println("User_data.GetAll error=", err)
		return datalist
	}
	return datalist
}

func (d *UserDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtUser{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

// func (d *UserDao) Delete(id int) error {
// 	data := &models.LtUser{Id:id}
// 	_, err :=d.engine.Id(data.Id).Update(data)
// 	return err
// }

func (d *UserDao) Update(data *models.LtUser, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *UserDao) Create(data *models.LtUser) error {
	_, err := d.engine.Insert(data)
	return err
}