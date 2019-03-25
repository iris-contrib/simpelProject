package controller

import (
	"github.com/kataras/iris"
	"github.com/majidbigdeli/simpelProject/domin/data"
	"github.com/majidbigdeli/simpelProject/domin/dto"
	"github.com/majidbigdeli/simpelProject/domin/util"
)

//Create User ...
func Create(ctx iris.Context) {
	user := dto.User{}

	if err := ctx.ReadJSON(&user); err != nil {
		resp := dto.NewResponse(false, nil, err.Error())
		var _, _ = ctx.JSON(resp)
		return
	}

	if err := util.IsValid(user); err != nil {
		resp := dto.NewResponse(false, nil, err.Error())
		var _, _ = ctx.JSON(resp)
		return
	}

	user.Password = util.HashAndSalt(user.Password)

	UserID, err := data.CreateUser(user)

	if err != nil {
		resp := dto.NewResponse(false, nil, err.Error())
		var _, _ = ctx.JSON(resp)
		return
	}

	CreateUser := dto.NewCreateUser(UserID)

	resp := dto.NewResponse(true, CreateUser, "")
	var _, _ = ctx.JSON(resp)
	return
}

//Get All User ...
func Get(ctx iris.Context) {
	users, err := data.Get()
	if err != nil {
		resp := dto.NewResponse(false, nil, err.Error())
		var _, _ = ctx.JSON(resp)
		return
	}
	resp := dto.NewResponse(true, users, "")
	var _, _ = ctx.JSON(resp)
	return
}

//GetUser : Get User By Token ...
func GetUser(ctx iris.Context) {
	userID := util.GetUserIDFromToken(ctx)
	user, err := data.GetUser(userID)
	if err != nil {
		resp := dto.NewResponse(false, nil, err.Error())
		var _, _ = ctx.JSON(resp)
		return
	}
	resp := dto.NewResponse(true, user, "")
	var _, _ = ctx.JSON(resp)
	return
}
