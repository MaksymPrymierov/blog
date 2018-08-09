package models

type Users struct {
	Id         string
	Username   string
	Password   string
	Permission string
}

func NewUser(Id, Username, Password, Permission string) *Users {
	return &Users{Id, Username, Password, Permission}
}
