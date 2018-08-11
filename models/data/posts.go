package data

import (
	"../../models"
)

/* Structure for template which data post */
type PostsData struct {
	DataPosts models.Post
	UserData  string
}

/* Init */
func NewPostsData(DataPosts models.Post, UserData string) *PostsData {
	return &PostsData{DataPosts, UserData}
}
