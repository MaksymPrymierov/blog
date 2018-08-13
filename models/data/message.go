package data

import (
	"../../models"
)

type MessageData struct {
	UserData models.PublicUsersData
	Message  string
}

func NewMesssageData(UserData models.PublicUsersData, Message string) *MessageData {
	return &MessageData{UserData, Message}
}
