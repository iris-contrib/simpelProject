package main

import (
	"SimpleProject/Domin/util"
	_ "SimpleProject/Startup"
	"SimpleProject/UI/controller"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	app.Use(util.Cors)
	//AccountController
	accountController := app.Party("/api/Account")
	accountController.Post("/Login", controller.Login)

	//UserController
	userController := app.Party("/api/User", util.MyJwtMiddleware.Serve)
	userController.Post("/Create", controller.Create)
	userController.Get("/Get", controller.Get)
	userController.Get("/GetUser", controller.GetUser)

	var _ = app.Run(
		iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		iris.WithoutPathCorrectionRedirection,
	)
}
