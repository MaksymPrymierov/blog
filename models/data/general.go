package data

import (
	"../../models"
)

type GeneralData struct {
	UserData models.PublicUsersData
}

func NewGeneralData(UserData models.PublicUsersData) *GeneralData {
	return &GeneralData{UserData}
}
