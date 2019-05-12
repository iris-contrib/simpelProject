package main

import (
	"github.com/kataras/iris"
	"github.com/majidbigdeli/simpelProject/domin/util"
	_ "github.com/majidbigdeli/simpelProject/startup"
	"github.com/majidbigdeli/simpelProject/ui/controller"
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
