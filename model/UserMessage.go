package model

import (
	"Program/constants"
	"Program/helper"
)

// 保存用户发送的消息
func SaveMessage(data UserMessage) helper.ReturnType {
	if err := db.Create(&data).Error; err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "保存信息失败!", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "保存信息成功!", Data: data.Message}
	}
}

// 在聊天室内展示聊天记录
func ShowMessageRecord() helper.ReturnType {
	var Usermessages []UserMessage
	err := db.
		Find(&Usermessages).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "调出信息失败!", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "调出信息成功!", Data: Usermessages}
	}
}

// 通过Nickname查询聊天记录
func ShowMessageRecordByname(name string) helper.ReturnType {
	var usermessages []UserMessage
	err := db.
		Where("nickname1 = ?", name).
		Find(&usermessages).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "调出信息失败!", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "调出信息成功!", Data: usermessages}
	}
}

