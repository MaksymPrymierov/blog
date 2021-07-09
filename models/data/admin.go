package data

import (
	"github.com/MaksymPrymierov/blog/models"
)

type AdminData struct {
	UserData models.PublicUsersData
}

func NewAdminData(UserData models.PublicUsersData) *AdminData {
	return &AdminData{UserData}
}
