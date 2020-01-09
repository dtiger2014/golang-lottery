package services

import (
	"golang-lottery/models"
	"golang-lottery/dao"
)

// GiftService :
type GiftService interface{
	GetAll(useCache bool) []models.LtGift
	CountAll() int64
	Get(id int, useCache bool) *models.LtGift
	Delete(id int) error
	Update(data *models.LtGift, columns []string) error
	Create(data *models.LtGift) error
	// GetAllUse(useCache bool) []models.ObjGiftPrize
	IncrLeftNum(id, num int) (int64, error)
	DecrLeftNum(id, num int) (int64, error)
}

type giftService struct {
	dao *dao.GiftDao
}

// NewGiftService : 
func NewGiftService() GiftService {
	return &giftService{
		dao: dao.NewGiftDao(nil),
	}
}

// GetAll :
func (s *giftService) GetAll(useCache bool) []models.LtGift{
	return s.dao.GetAll()
}

// CountAll :
func (s *giftService) CountAll() int64{
	return s.dao.CountAll()
}

// Get :
func (s *giftService) Get(id int, useCache bool) *models.LtGift{
	return s.dao.Get(id)
}

// Delete : 
func (s *giftService) Delete(id int) error{
	return s.dao.Delete(id)
}

// Update : 
func (s *giftService) Update(data *models.LtGift, columns []string) error{
	return s.dao.Update(data, columns)
}

// Create :
func (s *giftService) Create(data *models.LtGift) error{
	return s.dao.Create(data)
}

// GetAllUse : 
// func (s *giftService) GetAllUse(useCache bool) []models.ObjGiftPrize {
	
// }

// IncrLeftNum :
func (s *giftService) IncrLeftNum(id, num int) (int64, error) {
	return s.dao.IncrLeftNum(id, num)
}

// DecrLeftNum :
func (s *giftService) DecrLeftNum(id, num int) (int64, error) {
	return s.dao.DecrLeftNum(id, num)
}