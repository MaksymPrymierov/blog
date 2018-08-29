package documents

import (
	"github.com/connor41/blog/models"
)

/* Structure for data base collection */
type PostDocument struct {
	Id              string `bson:"_id,omitempty"`
	Title           string
	ContentHtml     string
	ContentMarkdown string
	Time            models.CurrentTime
	Owner           string
}
