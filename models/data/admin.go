package data

import (
	"github.com/connor41/blog/models"
)

type AdminPages struct {
	InfoServer bool
	Users      bool
}

func NewAdminPages(InfoServer, Users bool) *AdminPages {
	return &AdminPages{InfoServer, Users}
}

type AdminInfoServerData struct {
	Pages    AdminPages
	UserData models.PublicUsersData
	Day      int
	Hour     int
	Minute   int
}

func NewAdminInfoServerData(Pages AdminPages, UserData models.PublicUsersData, Day, Hour, Minute int) *AdminInfoServerData {
	return &AdminInfoServerData{Pages, UserData, Day, Hour, Minute}
}
