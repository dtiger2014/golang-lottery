package conf

import (
	"time"
)

const (
	SysTimeform = "2006-01-02 15:04:05"
	SysTimeformShort = "2006-01-02"
)

var (
	SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")

	SignSecret = []byte("123456")

	CookieSecret = "lottery123"
)

