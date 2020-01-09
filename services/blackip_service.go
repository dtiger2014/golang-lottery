package services

import (
	"golang-lottery/models"
	"golang-lottery/dao"
)

type BlackipService interface{
	GetAll() []models.LtBlackip
	CountAll() int64
	Get(id int) *models.LtBlackip
	Delete(id int) error
	Update(data *models.LtBlackip, columns []string) error
	Create(data *models.LtBlackip) error
}

type blackipService struct {
	dao *dao.BlackipDao
}

func NewBlackipService() BlackipService {
	return &blackipService{
		dao: *dao.NewBlackipDao(nil),
	}
}

func (s *BlackipService)GetAll() []models.LtGift{
	return s.dao.GetAll()
}
func (s *giftService)CountAll() int64{
	return s.dao.CountAll()
}
func (s *giftService)Get(id int) *models.LtGift{
	return s.dao.Get(id)
}
func (s *giftService)Delete(id int) error{
	return s.dao.Delete(id)
}
func (s *giftService)Update(data *models.LtGift, columns []string) error{
	return s.dao.Update(data, columns)
}
func (s *giftService)Create(data *models.LtGift) error{
	return s.dao.Create(data)
}