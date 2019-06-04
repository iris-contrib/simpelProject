package controller

import (
	"github.com/kataras/iris"
	"github.com/majidbigdeli/simpelProject/domin/data"
	"github.com/majidbigdeli/simpelProject/domin/dto"
	"github.com/majidbigdeli/simpelProject/domin/util"
)

//Create User ...
func Create(ctx iris.Context) {
	user := dto.UserDto{}

	if err := ctx.ReadJSON(&user); err != nil {
		resp := dto.NewResponse(false, nil, err)
		ctx.StatusCode(iris.StatusBadRequest)
		var _, _ = ctx.JSON(resp)
		return
	}

	if err := util.IsValid(user); err != nil {
		resp := dto.NewResponse(false, nil, err)
		ctx.StatusCode(iris.StatusBadRequest)
		var _, _ = ctx.JSON(resp)
		return
	}

	user.Password = util.HashAndSalt(user.Password)

	userID, err := data.CreateUser(user)

	if err != nil {
		resp := dto.NewResponse(false, nil, err)
		ctx.StatusCode(iris.StatusConflict)
		var _, _ = ctx.JSON(resp)
		return
	}

	returnID := dto.NewReturnID(userID)

	resp := dto.NewResponse(true, returnID, nil)
	ctx.StatusCode(iris.StatusCreated)
	var _, _ = ctx.JSON(resp)
	return
}

//Get All User ...
func Get(ctx iris.Context) {
	users, err := data.Get()
	if err != nil {
		resp := dto.NewResponse(false, nil, err)
		ctx.StatusCode(iris.StatusInternalServerError)
		var _, _ = ctx.JSON(resp)
		return
	}
	resp := dto.NewResponse(true, users, nil)
	ctx.StatusCode(iris.StatusOK)
	var _, _ = ctx.JSON(resp)
	return
}

//GetUser : Get User By Token ...
func GetUser(ctx iris.Context) {
	userID := util.GetUserIDFromToken(ctx)
	user, err := data.GetUser(userID)
	if err != nil {
		resp := dto.NewResponse(false, nil, err)
		ctx.StatusCode(iris.StatusInternalServerError)
		var _, _ = ctx.JSON(resp)
		return
	}
	resp := dto.NewResponse(true, user, nil)
	ctx.StatusCode(iris.StatusOK)
	var _, _ = ctx.JSON(resp)
	return
}
