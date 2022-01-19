package controller

import (
	"Program/constants"
	"Program/helper"
	"Program/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var UserModel struct {
		model.User
		PasswordCheck string
	}
	if err := c.ShouldBind(&UserModel); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据绑定模型错误", err.Error()))
	}

	// 校验密码输入
	if UserModel.Password != UserModel.PasswordCheck {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "两次密码输入不一致", ""))
	}

	// 密码加密
	UserModel.Password = helper.Getmd5(UserModel.Password)

	res := model.AddUser(UserModel.User)
	if res.Status == constants.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	}
	res = model.AddUserLevel(UserModel.User)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}
