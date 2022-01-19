package model

import (
	"Program/mysql"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB = mysql.MySqlDb

