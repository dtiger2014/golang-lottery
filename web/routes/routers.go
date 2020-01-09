package routes

import (
	"github.com/kataras/iris/v12/mvc"

	"golang-lottery/bootstrap"
	"golang-lottery/services"
	"golang-lottery/web/controllers"
	
)

// Configure registers the necessary routes to the app
func Configure(b *bootstrap.Bootstrapper) {
	userService := services.NewUserService()
	giftService := services.NewGiftService()
	codeService := services.NewCodeService()
	resultService := services.NewResultService()
	userdayService := services.NewUserdayService()
	blackipService := services.NewBlackipService()

	index := mvc.New(b.Party("/"))
	index.Register(userService, giftService, codeService, resultService, userdayService, blackipService)
	index.Handle(new(controllers.IndexController))
}