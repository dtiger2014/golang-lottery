package bootstrap

import (
	"io/ioutil"
	"time"

	// "github.com/gorilla/securecookie"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	// "github.com/kataras/iris/v12/sessions"

	"golang-lottery/conf"
	// "golang-lottery/cron"
)

// Configurator : 
type Configurator func(*Bootstrapper)

// Bootstrapper :
// 使用Go内建的嵌入机制(匿名嵌入)，允许类型之前共享代码和数据
// （Bootstrapper继承和共享 iris.Application ）
// 参考文章： https://hackthology.com/golangzhong-de-mian-xiang-dui-xiang-ji-cheng.html
type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
}

// New : returns a new Bootstrapper
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	var (
		b *Bootstrapper
	)

	b = &Bootstrapper{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

// SetupViews : loads the templates
func (b *Bootstrapper) SetupViews(viewsDir string) {
	htmlEngine := iris.HTML(viewsDir, ".html").
		Layout("shared/layout.html")
	// 每次重新加载模板（测试用，线上关闭）
	htmlEngine.Reload(true)
	// 给模板内置各种定制的方法
	htmlEngine.AddFunc("FromUnixtimeShort", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeformShort)
	})
	htmlEngine.AddFunc("FromUnixtime", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeform)
	})
	b.RegisterView(htmlEngine)
}

// SetupErrorHandlers prepares the http error handlers
// `(context.StatusCodeNotSuccessful`,  
// which defaults to < 200 || >= 400 but you can change it).
func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app": b.AppName,
			"status": ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}

// Configure : accepts configurations and runs them 
// inside the Bootstraper's context.
func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

// 启动计划任务服务
func (b *Bootstrapper) setupCron() {
	
}

const(
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = "./public"
	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "/favicon.ico"
)

// Bootstrap : prepares our application.
//
// Returns itself.
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.SetupViews("./views")
	b.SetupErrorHandlers()

	// static files
	b.Favicon(StaticAssets + Favicon)
	// b.StaticWeb(StaticAssets[1:], StaticAssets)
	indexHTML, err := ioutil.ReadFile(StaticAssets + "/index.html")
	if err == nil {
		b.StaticContent(StaticAssets[1:] + "/", "text/html", indexHTML)
	}
	// 不要把目录末尾"/"省略掉
	iris.WithoutPathCorrectionRedirection(b.Application)

	// crontab

	// middleware, after static files
	b.Use(recover.New())
	b.Use(logger.New())

	return b
}

// Listen : starts the http server with the specified "addr"
func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}