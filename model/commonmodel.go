package model

import (
	"Program/mysql"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB = mysql.MySqlDb

type ReturnType struct {
	Status int
	Msg    string
	Data   interface{}
}

func ApiReturn(status int, msg string, data interface{}) gin.H {
	return gin.H{
		"status":  status,
		"message": msg,
		"data":    data,
	}
}
