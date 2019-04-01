package data

import (
	"github.com/connor41/blog/models"
)

/* Structure for template which data post */
type PostsData struct {
	DataPosts models.Post
	UserData  models.PublicUsersData
	Comment   []models.Comment
}

/* Init */
func NewPostsData(DataPosts models.Post, UserData models.PublicUsersData, Comment []models.Comment) *PostsData {
	return &PostsData{DataPosts, UserData, Comment}
}
