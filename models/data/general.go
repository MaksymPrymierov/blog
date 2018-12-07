package data

import (
	"github.com/connor41/blog/models"
)

type GeneralData struct {
	UserData models.PublicUsersData
}

func NewGeneralData(UserData models.PublicUsersData) *GeneralData {
	return &GeneralData{UserData}
}
