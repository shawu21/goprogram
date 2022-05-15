package controller

import (
	"Program/constants"
	"Program/helper"
	"Program/model"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 实现多人聊天室
func DoUserMessage(c *gin.Context) {
	var usermessage model.UserMessage

	nickname, _ := c.Get("nickname")

	c.JSON(http.StatusOK, gin.H{
		"nickname": nickname,
	})

	// 判断token是否读出用户身份
	if nickname == "" {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "身份信息已失效", nickname))
		return
	} else {
		res := model.ShowMessageRecord() // 展示聊天记录
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	}

	usermessage.Nickname = nickname.(string)

	go ListenRoom()

	conn, err := net.Dial("tcp", constants.Port)
	if err != nil {
		fmt.Println("dial err=", err)
		return
	}

	usermessage.Message = c.PostForm("message")
	if usermessage.Message != "" {
		model.SaveMessage(usermessage)
		conn.Write([]byte(usermessage.Message))
	}

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Read err=", err)
				return
			}
			msg := string(buf[:n])
			if msg != "" {
				c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "信息接收成功", msg))
			}
		}
	}()

	return
}
