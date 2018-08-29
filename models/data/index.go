package data

import (
	"github.com/connor41/blog/models"
)

/* Structure for template index */
type IndexData struct {
	DataPosts []models.Post
	UserData  models.PublicUsersData
}

/* Init */
func NewIndexData(DataPosts []models.Post, UserData models.PublicUsersData) *IndexData {
	return &IndexData{DataPosts, UserData}
}
