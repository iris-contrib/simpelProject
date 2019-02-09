package Controller

import (
	"SimpleProject/Domin/Data"
	dto "SimpleProject/Domin/Dto"
	"SimpleProject/Domin/Util"

	"github.com/kataras/iris"
)

func Create(ctx iris.Context) {
	user := dto.User{}

	if err := ctx.ReadJSON(&user); err != nil {
		Response.Data = nil
		Response.Status = false
		Response.ErrorMessage = err.Error()
		ctx.JSON(Response)
		return
	}

	if err := Util.IsValid(user); err != nil {
		Response.Data = nil
		Response.Status = false
		Response.ErrorMessage = err.Error()
		ctx.JSON(Response)
		return
	}

	user.Password = Util.HashAndSalt(user.Password)

	userId, err := Data.CreateUser(user)

	if err != nil {
		Response.Data = nil
		Response.Status = false
		Response.ErrorMessage = err.Error()
		ctx.JSON(Response)
		return
	}

	var CreateUser dto.CreateUser
	CreateUser.UserId = userId
	Response.Data = CreateUser
	Response.Status = true
	Response.ErrorMessage = ""
	ctx.JSON(Response)
}
