package main

import (
	"fmt"

	"github.com/majidbigdeli/simpelProject/domin/data"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/kardianos/minwinsvc"
	"github.com/kataras/iris"
	"github.com/majidbigdeli/simpelProject/domin/util"
	"github.com/majidbigdeli/simpelProject/ui/controller"
	"github.com/spf13/viper"
)

func init() {
	//	path, err := osext.ExecutableFolder()

	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error ExecutableFolder: %s \n", err))

	// }

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}

	util.InitValidator()
	data.GetDB()

}

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
