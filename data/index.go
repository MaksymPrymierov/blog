package data

import (
	"../models"
)

type IndexData struct {
	DataPosts []models.Post
	UserData  string
}

func NewIndexData(DataPosts []models.Post, UserData string) *IndexData {
	return &IndexData{DataPosts, UserData}
}
