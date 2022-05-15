package controller

import (
	"Program/constants"
	"Program/helper"
	"Program/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoAddFriend(c *gin.Context) {
	name := c.PostForm("nickname")

	nickname, _ := c.Get("nickname")
	myname := nickname.(string)

	id := model.UseridCheck(name)

	if id == constants.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "查无该用户!", name))
		return
	}

	res := model.AddFriend(myname, id)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}
