package model

import (
	"Program/mysql"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB = mysql.MySqlDb

func init() {
	db.AutoMigrate(
		&User{},
		&Userlevel{},
		&UserMessage{},
		&Friends{},
	)
}
