package main

import (
	"SimpleProject/Domin/data"
	"SimpleProject/Domin/util"
	"SimpleProject/UI/controller"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/kardianos/minwinsvc"
	"github.com/kataras/iris"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
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
