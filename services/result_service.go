package services

import (
	"golang-lottery/dao"
	"golang-lottery/models"
)

// ResultService : 
type ResultService interface {
	GetAll(page, size int) []models.LtResult
	CountAll() int64
	GetNewPrize(size int, giftIDs []int) []models.LtResult
	SearchByGift(giftID, page, size int) []models.LtResult
	SearchByUser(uid, page, size int) []models.LtResult
	CountByGift(giftID int) int64
	CountByUser(uid int) int64
	Get(id int) *models.LtResult
	Delete(id int) error
	Update(data *models.LtResult, columns []string) error
	Create(data *models.LtResult) error
}

type resultService struct {
	dao *dao.ResultDao
}

// NewResultService :
func NewResultService() ResultService {
	return &resultService{
		dao: dao.NewResultDao(nil),
	}
}

// GetAll :
func (s *resultService) GetAll(page, size int) []models.LtResult {
	return s.dao.GetAll(page, size)
}

// CountAll :
func (s *resultService) CountAll() int64 {
	return s.dao.CountAll()
}

// GetNewPrize :
func (s *resultService) GetNewPrize(size int, giftIDs []int) []models.LtResult {
	return s.dao.GetNewPrize(size, giftIDs)
}

// SearchByGift :
func (s *resultService) SearchByGift(giftID, page, size int) []models.LtResult {
	return s.dao.SearchByGift(giftID, page, size)
}

// SearchByUser :
func (s *resultService) SearchByUser(uid, page, size int) []models.LtResult {
	return s.dao.SearchByUser(uid, page, size)
}

// CountByGift :
func (s *resultService) CountByGift(giftID int) int64 {
	return s.dao.CountByGift(giftID)
}

// CountByUser :
func (s *resultService) CountByUser(uid int) int64 {
	return s.dao.CountByUser(uid)
}

// Get :
func (s *resultService) Get(id int) *models.LtResult {
	return s.dao.Get(id)
}

// Delete :
func (s *resultService) Delete(id int) error {
	return s.dao.Delete(id)
}

// Update :
func (s *resultService) Update(data *models.LtResult, columns []string) error {
	return s.dao.Update(data, columns)
}

// Create :
func (s *resultService) Create(data *models.LtResult) error {
	return s.dao.Create(data)
}
