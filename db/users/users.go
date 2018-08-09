package users

type UsersTable struct {
	Id         string `bson:"_id,omitempty"`
	Email      string
	Username   string
	Password   string
	Permission string
}
