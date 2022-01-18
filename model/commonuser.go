package model

type User struct {
	ID       int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Nickname string
	Email    string
	Password string
}

type Userlevel struct {
	ID     int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	UserID int
	Level  int
}
