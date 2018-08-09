package models

type Users struct {
	Id         string
	Email      string
	Username   string
	Password   string
	Permission string
}

func NewUser(Id, Username, Email, Password, Permission string) *Users {
	return &Users{Id, Username, Email, Password, Permission}
}
