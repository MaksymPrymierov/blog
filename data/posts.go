package data

import (
	"../models"
)

type PostsData struct {
	DataPosts models.Post
	UserData  string
}

func NewPostsData(DataPosts models.Post, UserData string) *PostsData {
	return &PostsData{DataPosts, UserData}
}
