package controller

import (
	"Program/constants"
	"Program/helper"
	"Program/model"
	"Program/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.User

	// if err := c.ShouldBind(&user); err != nil {
	// 	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据模型绑定错误", err.Error()))
	// }

	user.Nickname = c.PostForm("nickname")
	user.Password = c.PostForm("password")

	user.Password = helper.Getmd5(user.Password)

	res := model.UserLogin(user)
	if res.Status == constants.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	} else {
		// 生成token
		tokenString, jwterr := token.GenToken(user.Nickname)
		if jwterr != nil {
			print(jwterr.Error())
		}
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, gin.H{"token": tokenString}))
	}
	return
}
