package services

import (
	"golang-lottery/datasource"
	"golang-lottery/models"
	"golang-lottery/dao"
	"golang-lottery/comm"

	"sync"
	"fmt"
	"log"
	"github.com/gomodule/redigo/redis"
)

var cachedBlackipList = make(map[string]*models.LtBlackip)
var cachedBlackipLock = sync.Mutex{}

// BlackipService :
type BlackipService interface{
	GetAll(page, size int) []models.LtBlackip
	CountAll() int64
	Search(ip string) []models.LtBlackip
	Get(id int) *models.LtBlackip
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
		dao: dao.NewBlackipDao(datasource.InstanceDbMaster()),
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

// Update :
func (s *blackipService) Update(data *models.LtBlackip, columns []string) error{
	
	// 先更新缓存的数据
	s.updateByCache(data, columns)

	return s.dao.Update(data, columns)
}

// Create : 
func (s *blackipService) Create(data *models.LtBlackip) error{
	return s.dao.Create(data)
}

// GetByIP :
func (s *blackipService) GetByIP(ip string) *models.LtBlackip {
	// 先从缓存中读取数据
	data := s.getByCache(ip)
	if data == nil || data.Ip == "" {
		// 再从数据库中读取数据
		data = s.dao.GetByIP(ip)
		if data == nil || data.Ip == "" {
			data = &models.LtBlackip{Ip: ip}
		}
		s.setByCache(data)
	}

	return data
}

func (s *blackipService) getByCache(ip string) *models.LtBlackip  {
	// 集群模式，redis缓存
	key := fmt.Sprintf("info_blackip_%s", ip)
	rds := datasource.InstanceCache()
	dataMap, err := redis.StringMap(rds.Do("HGETALL", key))
	if err != nil {
		log.Println("blackip_service.getByCache HGETALL key=", key, ", error=", err)
		return nil
	}
	dataIp := comm.GetStringFromStringMap(dataMap, "Ip", "")
	if dataIp == "" {
		return nil
	}
	data := &models.LtBlackip{
		Id:         int(comm.GetInt64FromStringMap(dataMap, "Id", 0)),
		Ip:         dataIp,
		Blacktime:  int(comm.GetInt64FromStringMap(dataMap, "Blacktime", 0)),
		SysCreated: int(comm.GetInt64FromStringMap(dataMap, "SysCreated", 0)),
		SysUpdated: int(comm.GetInt64FromStringMap(dataMap, "SysUpdated", 0)),
	}
	return data
}

func (s *blackipService) setByCache(data *models.LtBlackip) {
	if data == nil || data.Ip == "" {
		return
	}
	// 集群模式，redis缓存
	key := fmt.Sprintf("info_blackip_%s", data.Ip)
	rds := datasource.InstanceCache()
	// 数据更新到redis缓存
	params := []interface{}{key}
	params = append(params, "Ip", data.Ip)
	if data.Id > 0 {
		params = append(params, "Blacktime", data.Blacktime)
		params = append(params, "SysCreated", data.SysCreated)
		params = append(params, "SysUpdated", data.SysUpdated)
	}
	_, err := rds.Do("HMSET", params...)
	if err != nil {
		log.Println("blackip_service.setByCache HMSET params=", params, ", error=", err)
	}
}

// 数据更新了，直接清空缓存数据
func (s *blackipService) updateByCache(data *models.LtBlackip, columns []string) {
	if data == nil || data.Ip == "" {
		return
	}
	// 集群模式，redis缓存
	key := fmt.Sprintf("info_blackip_%s", data.Ip)
	rds := datasource.InstanceCache()
	// 删除redis中的缓存
	rds.Do("DEL", key)
}
