package mysql

import (
	"Program/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MySqlDb *gorm.DB
var MySqlDbErr error

func init() {
	dbConfig := config.GetDbConfig()

	dbDSN := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbConfig["username"],
		dbConfig["password"],
		dbConfig["hostname"],
		dbConfig["port"],
		dbConfig["databasename"],
		dbConfig["charset"],
		dbConfig["parseTime"],
		dbConfig["local"],
	)

	MySqlDb, MySqlDbErr = gorm.Open("mysql", dbDSN)

	if MySqlDbErr != nil {
		panic("database open error" + MySqlDbErr.Error())
	}
}
