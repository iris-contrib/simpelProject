package main

import (
	"SimpleProject/Domin/data"
	"SimpleProject/Domin/util"
	"SimpleProject/UI/controller"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/kardianos/minwinsvc"
	"github.com/kataras/iris"
)

func main() {
	data.GetDB()
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
