package main

import (
	"fmt"
	"golang-lottery/bootstrap"
	"golang-lottery/web/routes"
)

var port = 8080

func newApp() *bootstrap.Bootstrapper {
	// 初始化应用
	app := bootstrap.New("Go抽奖系统", "饭饭")
	app.Bootstrap()
	app.Configure(routes.Configure)

	return app
}

func main() {
	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))
}