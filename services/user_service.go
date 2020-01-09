package services


import (
	// "fmt"
	"golang-lottery/dao"
	"golang-lottery/models"
)

// UserService :
type UserService interface{
	GetAll(page, size int) []models.LtUser
	CountAll() int64
	Get(id int) *models.LtUser
	Update(data *models.LtUser, columns []string) error
	Create(data *models.LtUser) error
}

type userService struct {
	dao *dao.UserDao
}

func NewUserService() UserService {
	return &userService {
		dao: dao.NewUserDao(nil),
	}
}

// GetAll :
func (s *userService) GetAll(page, size int) []models.LtUser{
	return s.dao.GetAll(page, size)
}

// CountAll :
func (s *userService) CountAll() int64 {
	return s.dao.CountAll()
}

// Get :
func (s *userService) Get(id int) *models.LtUser {
	return s.dao.Get(id)
}

// Update :
func (s *userService)Update(data *models.LtUser, columns []string) error {
	return s.dao.Update(data, columns)
}

// Create :
func (s *userService)Create(data *models.LtUser) error {
	return s.dao.Create(data)
}