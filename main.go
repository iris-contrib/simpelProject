package main

import (
	"SimpleProject/Domin/Util"
	"SimpleProject/UI/Controller"

	_ "github.com/denisenkom/go-mssqldb"
	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {

	myJwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(Util.JwtKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	app := iris.New()

	app.Use(recover.New())
	requestLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(requestLogger)

	//AccountController
	AccountController := app.Party("/Account")
	AccountController.Post("/Login", Controller.Login)

	//UserController
	UserController := app.Party("/User", myJwtMiddleware.Serve)
	UserController.Post("/Create", Controller.Create)

	app.Run(
		iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
