package dao

import (
	"xorm.io/xorm"
	"golang-lottery/models"
)

// UserdayDao :
type UserdayDao struct {
	engine *xorm.Engine
}

// NewUserdayDao :
func NewUserdayDao(engine *xorm.Engine) *UserdayDao {
	return &UserdayDao {
		engine: engine,
	}
}

// Get :
func (d *UserdayDao) Get(id int) *models.LtUserday {
	data := &models.LtUserday{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Id = 0
	return data
}

// GetAll :
func (d *UserdayDao) GetAll(page, size int) []models.LtUserday {
	offset := (page - 1) * size
	datalist := make([]models.LtUserday, 0)
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
func (d *UserdayDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtUserday{})
	if err != nil {
		return 0
	}
	return num
}

// Search :
func (d *UserdayDao) Search(uid, day int) []models.LtUserday {
	datalist := make([]models.LtUserday, 0)
	err := d.engine.
		Where("uid=?", uid).
		Where("day=?", day).
		Desc("id").
		Find(&datalist)
	if err != nil {
		return datalist
	}
	return datalist
}

// Count :
func (d *UserdayDao) Count(uid, day int) int {
	info := &models.LtUserday{}
	ok, err := d.engine.
		Where("uid=?", uid).
		Where("day=?", day).
		Get(info)
	if !ok || err != nil {
		return 0
	}
	return info.Num
}

// Update :
func (d *UserdayDao) Update(data *models.LtUserday, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

// Create : 
func (d *UserdayDao) Create(data *models.LtUserday) error {
	_, err := d.engine.Insert(data)
	return err
}