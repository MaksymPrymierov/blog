package models

/* Structure for posts */
type Post struct {
	Id              string
	Title           string
	ContentHTML     string
	ContentMarkdown string
	Time            CurrentTime
	Owner           string
}

/* Init */
func NewPost(Id, Title, ContentHTML, ContentMarkdown string, Time CurrentTime, Owner string) *Post {
	return &Post{Id, Title, ContentHTML, ContentMarkdown, Time, Owner}
}
