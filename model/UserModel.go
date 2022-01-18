package model

import "Program/constants"

// 添加用户
func AddUser(data User) ReturnType {
	user := User{}
	err := db.
		Where("nickname = ?", data.Nickname).
		First(&user).
		Error
	// 判断昵称是否存在
	if err == nil {
		return ReturnType{Status: constants.CodeError, Msg: "昵称已存在", Data: user}
	}
	// 创建用户
	err = db.Create(&data).Error
	if err != nil {
		return ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: 1}
	}
}

// 添加新的用户权限
func AddUserLevel(data User) ReturnType {
	userlevel := Userlevel{}
	userlevel.Level = 0

	err := db.
		Where("nick_name = ?", data.Nickname).
		Find(&data).
		Error
	if err != nil {
		return ReturnType{Status: constants.CodeError, Msg: "查找用户失败，无法创建权限", Data: err.Error()}
	}
	userlevel.UserID = data.ID

	err = db.Create(&userlevel).Error
	if err != nil {
		return ReturnType{Status: constants.CodeError, Msg: "创建权限失败", Data: err.Error()}
	} else {
		return ReturnType{Status: constants.CodeSuccess, Msg: "创建权限成功", Data: 1}
	}
}

// 更新用户权限
func UploadUserLevel(ID int) {

}
