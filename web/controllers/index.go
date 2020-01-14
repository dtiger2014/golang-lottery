package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"

	"golang-lottery/comm"
	"golang-lottery/conf"
	"golang-lottery/models"
	"golang-lottery/services"
)


// IndexController : index controller
type IndexController struct {
	Ctx            iris.Context
	ServiceUser    services.UserService
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceReuslt  services.ResultService
	ServiceUserday services.UserdayService
	ServiceBlackip services.BlackipService
}

// Get http://localhost:8080
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Go抽奖系统，<a href='/public/index.html'>开始抽奖</a>"
}

// GetGifts http://localhost:8080/gifts
func (c *IndexController) GetGifts() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	datalist := c.ServiceGift.GetAll(true)
	list := make([]models.LtGift, 0)
	for _, data := range datalist {
		// 正常状态的才需要放进来
		if data.SysStatus == 0 {
			list = append(list, data)
		}
	}
	rs["gifts"] = list
	return rs
}

// GetNewprize http://localhost:8080/newprize
func (c *IndexController) GetNewprize() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""

	gifts := c.ServiceGift.GetAll(true)
	giftIds := []int{}
	for _, data := range gifts {
		// 虚拟券或者实物奖才需要放到外部榜单中展示
		if data.Gtype > 1 {
			giftIds = append(giftIds, data.Id)
		}
	}

	list := c.ServiceReuslt.GetNewPrize(50, giftIds)
	rs["prize_list"] = list
	return rs
}

// GetMyprize http://localhost:8080/myprize
func (c *IndexController) GetMyprize() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	// 验证登录
	loginuser := comm.GetLoginUser(c.Ctx.Request())
	if loginuser == nil || loginuser.UID < 1 {
		rs["code"] = 101
		rs["msg"] = "请先登录，再来抽奖"
		return rs
	}
	// 只读取出来最新的100次中奖记录
	list := c.ServiceReuslt.SearchByUser(loginuser.UID, 1, 100)
	rs["prize_list"] = list
	// 今天抽奖次数
	day, _ := strconv.Atoi(comm.FormatFromUnixTimeShort(time.Now().Unix()))
	num := c.ServiceUserday.Count(loginuser.UID, day)
	rs["prize_num"] = conf.UserPrizeMax - num
	return rs
}

// GetLogin 登录 GET /login
func (c *IndexController) GetLogin() {
	// 每次随机生成一个登录用户信息
	uid := comm.Random(100000)
	loginuser := models.ObjLoginuser{
		UID:      uid,
		Username: fmt.Sprintf("admin-%d", uid),
		Now:      comm.NowUnix(),
		IP:       comm.ClientIP(c.Ctx.Request()),
	}
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/public/index.html?from=login"
	}
	comm.SetLoginuser(c.Ctx.ResponseWriter(), &loginuser)
	comm.Redirect(c.Ctx.ResponseWriter(), refer)
}

// GetLogout 退出 GET /logout
func (c *IndexController) GetLogout() {
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/public/index.html?from=logout"
	}
	comm.SetLoginuser(c.Ctx.ResponseWriter(), nil)
	comm.Redirect(c.Ctx.ResponseWriter(), refer)
}
