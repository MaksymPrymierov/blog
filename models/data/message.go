package data

import (
	"github.com/connor41/blog/models"
)

type MessageData struct {
	UserData models.PublicUsersData
	Message  string
}

func NewMesssageData(UserData models.PublicUsersData, Message string) *MessageData {
	return &MessageData{UserData, Message}
}
