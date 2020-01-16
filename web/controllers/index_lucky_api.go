package controllers

import(
	"fmt"
	"log"
	"time"
	"strconv"

	"golang-lottery/services"
	"golang-lottery/comm"
	"golang-lottery/conf"
	"golang-lottery/models"
	"golang-lottery/web/utils"
)

// LuckyApi 
type LuckyApi struct {
}

func (api *LuckyApi) luckyDo(uid int, username, ip string) (int, string, *models.ObjGiftPrize) {
	// 2 用户抽奖分布式锁定
	ok := utils.LockLucky(uid)
	if ok {
		defer utils.UnlockLucky(uid)
	} else {
		return 102, "正在抽奖，请稍后重试", nil
	}

	// 3 验证用户今日参与次数
	userDayNum := utils.IncrUserLuckyNum(uid)
	if userDayNum > conf.UserPrizeMax {
		return 103, "今日的抽奖次数已用完，明天再来吧", nil
	} else {
		ok = api.checkUserday(uid, userDayNum)
		if !ok {
			return 103, "今日的抽奖次数已用完，明天再来吧", nil
		}
	}

	// 4 验证IP今日的参与次数
	ipDayNum := utils.IncrIpLuckyNum(ip)
	if ipDayNum > conf.IpLimitMax {
		return 104, "相同IP参与次数太多，明天再来参与吧", nil
	}

	limitBlack := false // 黑名单
	if ipDayNum > conf.IpPrizeMax {
		limitBlack = true
	}

	// 5 验证IP黑名单
	var blackipInfo *models.LtBlackip
	if !limitBlack {
		ok, blackipInfo = api.checkBlackip(ip)
		if !ok {
			fmt.Println("黑名单中的IP", ip, limitBlack)
			limitBlack = true
		}
	}

	// 6 验证用户黑名单
	var userInfo *models.LtUser
	if !limitBlack {
		ok, userInfo = api.checkBlackUser(uid)
		if !ok {
			limitBlack = true
		}
	}

	// 7 获得抽奖编码
	prizeCode := comm.Random(10000)
	
	// 8 匹配奖品是否中奖
	prizeGift := api.prize(prizeCode, limitBlack)
	if prizeGift == nil ||
		prizeGift.PrizeNum < 0 ||
		(prizeGift.PrizeNum > 0 && prizeGift.LeftNum <= 0) {
		return 205, "很遗憾，没有中奖，请下次再试", nil
	}

	// 9 有限制奖品发放
	if prizeGift.PrizeNum > 0 {
		if utils.GetGiftPoolNum(prizeGift.Id) <= 0 {
			return 206, "很遗憾，没有中奖，请下次再试", nil
		}
		ok = utils.PrizeGift(prizeGift.Id, prizeGift.LeftNum)
		if !ok {
			return 207, "很遗憾，没有中奖，请下次再试", nil
		}
	}

	// 10 不同编码的优惠券的发放
	if prizeGift.Gtype == conf.GtypeCodeDiff {
		code := utils.PrizeCodeDiff(prizeGift.Id, services.NewCodeService())
		if code == "" {
			return 208, "很遗憾，没有中奖，请下次再试", nil
		}
		prizeGift.Gdata = code
	}

	// 11 记录中奖记录
	result := models.LtResult{
		GiftId:     prizeGift.Id,
		GiftName:   prizeGift.Title,
		GiftType:   prizeGift.Gtype,
		Uid:        uid,
		Username:   username,
		PrizeCode:  prizeCode,
		GiftData:   prizeGift.Gdata,
		SysCreated: comm.NowUnix(),
		SysIp:      ip,
		SysStatus:  0,
	}
	err := services.NewResultService().Create(&result)
	if err != nil {
		log.Println("index_lucky.GetLucky ServiceResult.Create ", result,
			", error=", err)
		return 209, "很遗憾，没有中奖，请下次再试", nil
	}
	if prizeGift.Gtype == conf.GtypeGiftLarge {
		// 如果获得了实物大奖，需要将用户、IP设置成黑名单一段时间
		api.prizeLarge(ip, uid, username, userInfo, blackipInfo)
	}

	// 12 返回抽奖结果
	return 0, "", prizeGift
}



// checkUserday : 
func (api *LuckyApi) checkUserday(uid int, num int64) bool {
	userdayService := services.NewUserdayService()
	userdayInfo := userdayService.GetUserToday(uid)
	if userdayInfo != nil && userdayInfo.Uid == uid {
		// 今天存在抽奖记录
		if userdayInfo.Num >= conf.UserPrizeMax {
			if int(num) < userdayInfo.Num {
				utils.InitUserLuckyNum(uid, int64(userdayInfo.Num))
			}
			return false
		} else {
			userdayInfo.Num++
			if int(num) < userdayInfo.Num {
				utils.InitUserLuckyNum(uid, int64(userdayInfo.Num))
			}
			err103 := userdayService.Update(userdayInfo, nil)
			if err103 != nil {
				log.Println("index_lucky_check_userday ServiceUserDay.Update " +
					"err103=", err103)
			}
		}
	} else {
		// 创建今天的用户参与记录
		y, m, d := time.Now().Date()
		strDay := fmt.Sprintf("%d%02d%02d", y, m, d)
		day, _ := strconv.Atoi(strDay)
		userdayInfo = &models.LtUserday{
			Uid:        uid,
			Day:        day,
			Num:        1,
			SysCreated: int(time.Now().Unix()),
		}
		err103 := userdayService.Create(userdayInfo)
		if err103 != nil {
			log.Println("index_lucky_check_userday ServiceUserDay.Create " +
				"err103=", err103)
		}
		utils.InitUserLuckyNum(uid, 1)
	}
	return true	
}

// checkBlackip :
func (api *LuckyApi) checkBlackip(ip string)(bool, *models.LtBlackip) {
	info := services.NewBlackipService().GetByIP(ip)
	if info == nil || info.Ip == "" {
		return true, nil
	}

	if info.Blacktime > int(time.Now().Unix()) {
		// IP黑名单存在，并且还在黑名单有效期内
		return false, info
	}
	return true, info
}

// checkBlackUser : 
func (api *LuckyApi) checkBlackUser(uid int) (bool, *models.LtUser) {
	info := services.NewUserService().Get(uid)
	if info != nil && info.Blacktime > int(time.Now().Unix()) {
		// 黑名单存在并且有效
		return false, info
	}
	return true, info
}

// prize :
func (api *LuckyApi) prize(prizeCode int, limitBlack bool) *models.ObjGiftPrize {
	var prizeGift *models.ObjGiftPrize
	giftList := services.NewGiftService().GetAllUse(true)
	for _, gift := range giftList {
		if gift.PrizeCodeA <= prizeCode &&
			gift.PrizeCodeB >= prizeCode {
			// 中奖编码区间满足条件，说明可以中奖
			if !limitBlack || gift.Gtype < conf.GtypeGiftSmall {
				prizeGift = &gift
				break
			}
		}
	}
	return prizeGift
}

// prizeLarge
func (api *LuckyApi) prizeLarge(ip string, uid int, username string, userinfo *models.LtUser, blackipInfo *models.LtBlackip) {
	userService := services.NewUserService()
	blackipService := services.NewBlackipService()
	nowTime := comm.NowUnix()
	blackTime := 30 * 86400

	// 更新用户的黑名单信息
	if userinfo == nil || userinfo.Id <= 0 || userinfo.Username == "" {
		userinfo = &models.LtUser{
			Id:			uid,
			Username:   username,
			Blacktime:  nowTime+blackTime,
			SysCreated: nowTime,
			SysIp:      ip,
		}
		userService.Create(userinfo)
	} else {
		userinfo = &models.LtUser{
			Id: uid,
			Blacktime:nowTime+blackTime,
			SysUpdated:nowTime,
		}
		userService.Update(userinfo, nil)
	}

	// 更新要IP的黑名单信息
	if blackipInfo == nil || blackipInfo.Id <= 0 {
		blackipInfo = &models.LtBlackip{
			Ip:         ip,
			Blacktime:  nowTime+blackTime,
			SysCreated: nowTime,
		}
		blackipService.Create(blackipInfo)
	} else {
		blackipInfo.Blacktime = nowTime + blackTime
		blackipInfo.SysUpdated = nowTime
		blackipService.Update(blackipInfo, nil)
	}
}