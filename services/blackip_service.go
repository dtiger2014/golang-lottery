package services

import (
	"golang-lottery/models"
	"golang-lottery/dao"
)

// BlackipService :
type BlackipService interface{
	GetAll(page, size int) []models.LtBlackip
	CountAll() int64
	Search(ip string) []models.LtBlackip
	Get(id int) *models.LtBlackip
	// Delete(id int) error
	Update(data *models.LtBlackip, columns []string) error
	Create(data *models.LtBlackip) error
	GetByIP(ip string) *models.LtBlackip
}

// blackipService :
type blackipService struct {
	dao *dao.BlackipDao
}

// NewBlackipService :
func NewBlackipService() BlackipService {
	return &blackipService {
		dao: dao.NewBlackipDao(nil),
	}
}

// GetAll :
func (s *blackipService) GetAll(page, size int) []models.LtBlackip{
	return s.dao.GetAll(page, size)
}

// CountAll :
func (s *blackipService) CountAll() int64{
	return s.dao.CountAll()
}

// Search :
func (s *blackipService) Search(ip string) []models.LtBlackip {
	return s.dao.Search(ip)
}

// Get :
func (s *blackipService) Get(id int) *models.LtBlackip{
	return s.dao.Get(id)
}

// func (s *giftService)Delete(id int) error{
// 	return s.dao.Delete(id)
// }

// Update :
func (s *blackipService) Update(data *models.LtBlackip, columns []string) error{
	
	return s.dao.Update(data, columns)
}

// Create : 
func (s *blackipService) Create(data *models.LtBlackip) error{
	return s.dao.Create(data)
}

// GetByIP :
func (s *blackipService) GetByIP(ip string) *models.LtBlackip {
	return s.dao.GetByIP(ip)
}