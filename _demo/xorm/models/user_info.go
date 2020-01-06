package models

type UserInfo struct {
	Id         int    `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	Name       string `json:"name" xorm:"not null default '' comment('中文名') VARCHAR(50)"`
	SysCreated int    `json:"sys_created" xorm:"not null default 0 comment('创建时间') INT(10)"`
	SysUpdated int    `json:"sys_updated" xorm:"not null default 0 comment('最后修改时间') INT(10)"`
}
