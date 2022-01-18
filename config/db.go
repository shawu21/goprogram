package config

var dbConfig map[string]interface{}

func init() {
	dbConfig = make(map[string]interface{})

	dbConfig["username"] = "root"
	dbConfig["password"] = ""
	dbConfig["hostname"] = "127.0.0.1"
	dbConfig["port"] = "3306"
	dbConfig["databasename"] = "program"
	dbConfig["charset"] = "utf8mb4"
	dbConfig["parseTime"] = "True"
	dbConfig["local"] = "Local"
}

func GetDbConfig() map[string]interface{} {
	return dbConfig
}
