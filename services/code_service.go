package services

import (
	"golang-lottery/datasource"
	"golang-lottery/models"
	"golang-lottery/dao"
)

// CodeService :
type CodeService interface{
	GetAll(page, size int) []models.LtCode
	CountAll() int64
	CountByGift(giftID int) int64
	Search(giftID int) []models.LtCode
	Get(id int) *models.LtCode
	Delete(id int) error
	Update(data *models.LtCode, columns []string) error
	Create(data *models.LtCode) error
	NextUsingCode(giftID, codeID int) *models.LtCode
	UpdateByCode(data *models.LtCode, columns []string) error
}

// codeService :
type codeService struct {
	dao *dao.CodeDao
}

// NewCodeService :
func NewCodeService() CodeService {
	return &codeService {
		dao: dao.NewCodeDao(datasource.InstanceDbMaster()),
	}
}

// GetAll :
func (s *codeService) GetAll(page, size int) []models.LtCode{
	return s.dao.GetAll(page, size)
}

// CountAll :
func (s *codeService) CountAll() int64{
	return s.dao.CountAll()
}

// CountByGift :
func (s *codeService) CountByGift(giftID int) int64 {
	return s.dao.CountByGift(giftID)
}

// Search :
func (s *codeService) Search(giftID int) []models.LtCode {
	return s.dao.Search(giftID)
}

// Get :
func (s *codeService) Get(id int) *models.LtCode {
	return s.dao.Get(id)
}

// Delete
func (s *codeService) Delete(id int) error {
	return s.dao.Delete(id)
}

// Update :
func (s *codeService) Update(data *models.LtCode, columns []string) error{
	return s.dao.Update(data, columns)
}

// Create : 
func (s *codeService) Create(data *models.LtCode) error{
	return s.dao.Create(data)
}

// NextUsingCode :
func (s *codeService) NextUsingCode(giftID, codeID int) *models.LtCode {
	return s.dao.NextUsingCode(giftID, codeID)
}

// UpdateByCode :
func (s *codeService) UpdateByCode(data *models.LtCode, columns []string) error {
	return s.dao.UpdateByCode(data, columns)
}