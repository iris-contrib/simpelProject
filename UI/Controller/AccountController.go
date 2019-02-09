package Controller

import (
	"SimpleProject/Domin/Data"
	dto "SimpleProject/Domin/Dto"
	"SimpleProject/Domin/Util"

	"github.com/kataras/iris"
)

func Login(ctx iris.Context) {
	login := dto.Login{}
	if err := ctx.ReadJSON(&login); err != nil {
		Response.Data = nil
		Response.Status = false
		Response.ErrorMessage = err.Error()
		ctx.JSON(Response)
		return
	}
	if err := Util.IsValid(login); err != nil {
		Response.Data = nil
		Response.Status = false
		Response.ErrorMessage = err.Error()
		ctx.JSON(Response)
		return
	}

	user, err := Data.GetUserByUserName(login.UserName)

	if err != nil {
		Response.Data = nil
		Response.Status = false
		Response.ErrorMessage = err.Error()
		ctx.JSON(Response)
		return
	}

	checkPassword := Util.ComparePasswords(user.Password, login.Password)

	if checkPassword == false {
		Response.Data = nil
		Response.Status = false
		Response.ErrorMessage = "Password is mistake"
		ctx.JSON(Response)
		return
	}

	token, err := Util.CreateTokenEndpoint(user)

	if err != nil {
		Response.Data = nil
		Response.Status = false
		Response.ErrorMessage = err.Error()
		ctx.JSON(Response)
		return
	}

	var Token dto.Token

	Token.AuthToken = token

	Response.Data = Token
	Response.Status = true
	Response.ErrorMessage = ""
	ctx.JSON(Response)
}

func GetUser(ctx iris.Context) {
	users, _ := Data.GetUser()
	ctx.JSON(users)
}
