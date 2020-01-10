package comm

import(
	"net"
	"net/http"
	"net/url"
	"strconv"
	"log"
	"fmt"

	"golang-lottery/models"
	"golang-lottery/conf"
)

// ClientIP : 获取客户端IP地址
func ClientIP(request *http.Request) string {
	host, _, _ := net.SplitHostPort(request.RemoteAddr)
	return host
}


// Redirect : 跳转URL
func Redirect(writer http.ResponseWriter, url string) {
	writer.Header().Add("Location", url)
	writer.WriteHeader(http.StatusFound)
}

// GetLoginUser : 从cookie中的到当前登录的用户
func GetLoginUser(request *http.Request) *models.ObjLoginuser {
	c, err := request.Cookie("lottery_loginuser")
	if err != nil {
		return nil
	}

	params, err := url.ParseQuery(c.Value)
	if err != nil {
		return nil
	}

	uid, err := strconv.Atoi(params.Get("uid"))
	if err != nil || uid < 1 {
		return nil
	}

	// Cookie 最长使用时长，最大30天
	now, err := strconv.Atoi(params.Get("now"))
	if err != nil || NowUnix()-now > 86400*30 {
		return nil
	}

	// IP修改了是不是要重新登录
	// ip := params.Get("ip")
	// if ip != ClientIP(request) {
	// 	return nil
	// }

	// 登录信息
	loginuser := &models.ObjLoginuser{
		UID : uid,
		Username : params.Get("username"),
		Now : now,
		IP : ClientIP(request),
		Sign : params.Get("sign"),
	}

	sign := createLoginuserSign(loginuser)
	if sign != loginuser.Sign {
		log.Println("func_web GetLoginUser createLoginuserSign not sign", 
			sign, loginuser.Sign)
	}
	
	return loginuser
}

// SetLoginuser : 将登陆的用户信息设置到cookie中
func SetLoginuser(writer http.ResponseWriter, loginuser *models.ObjLoginuser) {
	if loginuser == nil || loginuser.UID < 1 {
		c := &http.Cookie{
			Name: "lottery_loginuser",
			Value: "",
			Path: "/",
			MaxAge: -1,
		}
		http.SetCookie(writer, c)
		return
	}

	if loginuser.Sign == "" {
		loginuser.Sign = createLoginuserSign(loginuser)
	}
	params := url.Values{}
	params.Add("uid", strconv.Itoa(loginuser.UID))
	params.Add("username", loginuser.Username)
	params.Add("now", strconv.Itoa(loginuser.Now))
	params.Add("ip", loginuser.IP)
	params.Add("sign", loginuser.Sign)
	c := &http.Cookie{
		Name: "lottery_loginuser",
		Value: params.Encode(),
		Path: "/",
	}
	http.SetCookie(writer, c)
}


// 将登陆的用户信息设置到cookie中
func createLoginuserSign(loginuser *models.ObjLoginuser) string {
	str := fmt.Sprintf("uid=%d&username=%s&secret=%s", 
		loginuser.UID, loginuser.Username, conf.CookieSecret)
	return CreateSign(str)
}