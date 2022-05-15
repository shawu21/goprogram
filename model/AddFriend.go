package model

import (
	"Program/constants"
	"Program/helper"
)

func AddFriend(name string, id int) helper.ReturnType {
	var user User
	var friendsid Friends
	friendsid.ID = id
	err := db.
		Where("nickname = ?", name).
		First(&user).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "添加好友失败!", Data: ""}
	}
	user.FriendsID = append(user.FriendsID, friendsid)
	db.Save(&user)
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "添加好友成功!", Data: ""}
}
