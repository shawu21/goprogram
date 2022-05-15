package model

type User struct {
	ID        int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	FriendsID []Friends `gorm:"foreignKey:UserID"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
}

type Friends struct {
	UserID int
	ID     int
}

type Userlevel struct {
	ID     int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	UserID int
	Level  int
}

// 公共聊天室
type UserMessage struct {
	ID       int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Nickname string `json:"nickname"`
	Message  string `json:"message"`
}

type AllUserMessage struct {
	ID          int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Usermessage []UserMessage
}
