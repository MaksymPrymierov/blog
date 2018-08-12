package users

/* Structure for database collection */
type UsersTable struct {
	Id         string `bson:"_id,omitempty"`
	Email      string `bson:"_email,omitempty"`
	Username   string `bson:"_username,omitempty"`
	Password   string `bson:"password"`
	Permission string `bson:"permission"`
}
