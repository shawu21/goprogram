package controller

import (
	"Program/constants"
	"Program/helper"
	"Program/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var UserModel model.User

	UserModel.Nickname = c.PostForm("nickname")
	UserModel.Password = c.PostForm("password")

	// 密码加密
	UserModel.Password = helper.Getmd5(UserModel.Password)

	res := model.AddUser(UserModel)
	if res.Status == constants.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	} else {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	}
	res = model.AddUserLevel(UserModel)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}
