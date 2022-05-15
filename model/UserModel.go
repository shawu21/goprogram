package model

import (
	"Program/constants"
	"Program/helper"
)

// 添加用户
func AddUser(data User) helper.ReturnType {
	user := User{}
	err := db.
		Where("nickname = ?", data.Nickname).
		First(&user).
		Error

	// 判断昵称是否存在
	if err == nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "昵称已存在", Data: user}
	}

	// 创建用户
	err = db.Create(&data).Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: 1}
	}
}

// 添加新的用户权限
func AddUserLevel(data User) helper.ReturnType {
	userlevel := Userlevel{}
	userlevel.Level = 0

	err := db.
		Where("nickname = ?", data.Nickname).
		Find(&data).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查找用户失败，无法创建权限", Data: err.Error()}
	}
	userlevel.UserID = data.ID

	err = db.Create(&userlevel).Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建权限失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建权限成功", Data: 1}
	}
}

// 更新用户权限
func UploadUserLevel(ID int) {

}

// 用户登录
func UserLogin(data User) helper.ReturnType {
	var user User
	err := db.
		Where("nickname = ? AND password = ?", data.Nickname, data.Password).
		First(&user).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "用户名或密码错误", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "登录成功", Data: user}
	}
}

// 通过用户名查询用户id
func UseridCheck(name string) int {
	var user User
	err := db.
		Where("nickname = ?", name).
		First(&user).
		Error

	if err != nil {
		return constants.CodeError
	} else {
		return user.ID
	}
}
