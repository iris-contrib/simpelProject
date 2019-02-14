package controller

import (
	dto "SimpleProject/Domin/Dto"
	"SimpleProject/Domin/data"
	"SimpleProject/Domin/util"

	"github.com/kataras/iris"
)

//Create User ...
func Create(ctx iris.Context) {
	user := dto.User{}

	if err := ctx.ReadJSON(&user); err != nil {
		response.Data = nil
		response.Status = false
		response.ErrorMessage = err.Error()
		ctx.JSON(response)
		return
	}

	if err := util.IsValid(user); err != nil {
		response.Data = nil
		response.Status = false
		response.ErrorMessage = err.Error()
		ctx.JSON(response)
		return
	}

	user.Password = util.HashAndSalt(user.Password)

	UserID, err := data.CreateUser(user)

	if err != nil {
		response.Data = nil
		response.Status = false
		response.ErrorMessage = err.Error()
		ctx.JSON(response)
		return
	}

	var CreateUser dto.CreateUser
	CreateUser.UserID = UserID
	response.Data = CreateUser
	response.Status = true
	response.ErrorMessage = ""
	ctx.JSON(response)
}

//Get All User ...
func Get(ctx iris.Context) {
	users, err := data.Get()
	if err != nil {
		response.Status = false
		response.Data = nil
		response.ErrorMessage = err.Error()
		ctx.JSON(response)
		return
	}
	response.Status = true
	response.Data = users
	response.ErrorMessage = ""
	ctx.JSON(response)
}

//GetUser : Get User By Token ...
func GetUser(ctx iris.Context) {
	userID := util.GetUserIDFromToken(ctx)

	user, err := data.GetUser(userID)

	if err != nil {
		response.Status = false
		response.Data = nil
		response.ErrorMessage = err.Error()
		ctx.JSON(response)
		return
	}

	response.Status = true
	response.Data = user
	response.ErrorMessage = ""
	ctx.JSON(response)

}
