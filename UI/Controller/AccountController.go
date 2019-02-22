package controller

import (
	"SimpleProject/Domin/Dto"
	"SimpleProject/Domin/data"
	"SimpleProject/Domin/util"

	"github.com/kataras/iris"
)

//Login ...
func Login(ctx iris.Context) {
	login := dto.Login{}
	if err := ctx.ReadJSON(&login); err != nil {
		response.Data = nil
		response.Status = false
		response.ErrorMessage = err.Error()
		var _, _ = ctx.JSON(response)
		return
	}
	if err := util.IsValid(login); err != nil {
		response.Data = nil
		response.Status = false
		response.ErrorMessage = err.Error()
		var _, _ = ctx.JSON(response)
		return
	}

	user, err := data.GetUserByUserName(login.UserName)

	if err != nil {
		response.Data = nil
		response.Status = false
		response.ErrorMessage = err.Error()
		var _, _ = ctx.JSON(response)
		return
	}

	checkPassword := util.ComparePasswords(user.Password, login.Password)

	if checkPassword == false {
		response.Data = nil
		response.Status = false
		response.ErrorMessage = "Password is mistake"
		var _, _ = ctx.JSON(response)
		return
	}

	token, err := util.CreateTokenEndpoint(user)

	if err != nil {
		response.Data = nil
		response.Status = false
		response.ErrorMessage = err.Error()
		var _, _ = ctx.JSON(response)
		return
	}

	var Token dto.Token

	Token.AuthToken = token

	response.Data = Token
	response.Status = true
	response.ErrorMessage = ""
	var _, _ = ctx.JSON(response)
}
