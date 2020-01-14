package conf

import (
	"time"
)

const (

	// 用户每天最多抽奖次数
	UserPrizeMax = 3000
	// 同一个IP每天最多抽奖次数
	IpPrizeMax = 30000
	// 同一个IP每天最多抽奖次数
	IpLimitMax = 300000

	GtypeVirtual   = 0 // 虚拟币
	GtypeCodeSame  = 1 // 虚拟券，相同的码
	GtypeCodeDiff  = 2 // 虚拟券，不同的码
	GtypeGiftSmall = 3 // 实物小奖
	GtypeGiftLarge = 4 // 实物大奖

	// SysTimeform 格式化 Y-m-d H:i:s
	SysTimeform = "2006-01-02 15:04:05"
	// SysTimeformShort 格式化 Y-m-d
	SysTimeformShort = "2006-01-02"
)

var (
	// RunningCrontabService : 是否需要启动全局计划任务服务
	RunningCrontabService = false

	// PrizeDataRandomDayTime : 定义24小时的奖品分配权重
	PrizeDataRandomDayTime = [100]int{
		// 24 * 3 = 72   平均3%的机会
		// 100 - 72 = 28 剩余28%的机会
		// 7 * 4 = 28    剩下的分别给7个时段增加4%的机会
		0, 0, 0,
		1, 1, 1,
		2, 2, 2,
		3, 3, 3,
		4, 4, 4,
		5, 5, 5,
		6, 6, 6,
		7, 7, 7,
		8, 8, 8,
		9, 9, 9, 9, 9, 9, 9,
		10, 10, 10, 10, 10, 10, 10,
		11, 11, 11,
		12, 12, 12,
		13, 13, 13,
		14, 14, 14,
		15, 15, 15, 15, 15, 15, 15,
		16, 16, 16, 16, 16, 16, 16,
		17, 17, 17, 17, 17, 17, 17,
		18, 18, 18,
		19, 19, 19,
		20, 20, 20, 20, 20, 20, 20,
		21, 21, 21, 21, 21, 21, 21,
		22, 22, 22,
		23, 23, 23,
	}


	// SysTimeLocation 时区，设置为上海
	SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")

	// SignSecret sign密码
	SignSecret = []byte("123456")

	// CookieSecret cookie密码
	CookieSecret = "lottery123"
)
