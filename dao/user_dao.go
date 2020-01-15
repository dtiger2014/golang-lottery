package dao

import (
	"xorm.io/xorm"
	"golang-lottery/models"
)

// UserDao :
type UserDao struct {
	engine *xorm.Engine
}

// NewUserDao :
func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao {
		engine: engine,
	}
}

// Get :
func (d *UserDao) Get(id int) *models.LtUser {
	data := &models.LtUser{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Id = 0
	return data
}

// GetAll : 
func (d *UserDao) GetAll(page, size int) []models.LtUser {
	offset := (page - 1) * size
	datalist := make([]models.LtUser, 0)
	err := d.engine.
		Desc("id").
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	}
	return datalist
}

// CountAll :
func (d *UserDao) CountAll() int {
	num, err := d.engine.Count(&models.LtUser{})
	if err != nil {
		return 0
	}
	return int(num)
}

// Update :
func (d *UserDao) Update(data *models.LtUser, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

// Create : 
func (d *UserDao) Create(data *models.LtUser) error {
	_, err := d.engine.Insert(data)
	return err
}