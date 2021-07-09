package data

import (
	"github.com/MaksymPrymierov/blog/models"
)

/* Structure for template which data post */
type PostsData struct {
	DataPosts models.Post
	UserData  models.PublicUsersData
}

/* Init */
func NewPostsData(DataPosts models.Post, UserData models.PublicUsersData) *PostsData {
	return &PostsData{DataPosts, UserData}
}
