package services

import(
	"fmt"
	"strconv"
	"time"
	"golang-lottery/dao"
	"golang-lottery/models"
	"golang-lottery/datasource"
)

// UserdayService :
type UserdayService interface {
	GetAll(page, size int) []models.LtUserday
	CountAll() int64
	Search(uid, day int) []models.LtUserday
	Count(uid, day int) int
	Get(id int) *models.LtUserday
	Update(data *models.LtUserday, columns []string) error
	Create(data *models.LtUserday) error
	GetUserToday(uid int) *models.LtUserday
}

type userdayService struct {
	dao *dao.UserdayDao
}

// NewUserdayService : 
func NewUserdayService() UserdayService {
	return &userdayService {
		dao: dao.NewUserdayDao(datasource.InstanceDbMaster()),
	}
}

// GetAll :
func (s *userdayService) GetAll(page, size int) []models.LtUserday {
	return s.dao.GetAll(page, size)
}

// CountAll : 
func (s *userdayService) CountAll() int64 {
	return s.dao.CountAll()
}

// Search : 
func (s *userdayService) Search(uid, day int) []models.LtUserday {
	return s.dao.Search(uid, day)
}

// Count :
func (s *userdayService) Count(uid, day int) int {
	return s.dao.Count(uid, day)
}

// Get :
func (s *userdayService) Get(id int) *models.LtUserday {
	return s.dao.Get(id)
}

// Update : 
func (s *userdayService) Update(data *models.LtUserday, columns []string) error{
	return s.dao.Update(data, columns)
}

// Create : 
func (s *userdayService) Create(data *models.LtUserday) error {
	return s.dao.Create(data)
}

// GetUserToday :
func (s *userdayService) GetUserToday(uid int) *models.LtUserday {
	y, m, d := time.Now().Date()
	strDay := fmt.Sprintf("%d%02d%02d", y, m, d)
	day, _ := strconv.Atoi(strDay)
	list := s.dao.Search(uid, day)
	if list != nil && len(list) > 0 {
		return &list[0]
	}

	return nil
}