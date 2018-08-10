package models

type Post struct {
	Id              string
	Title           string
	ContentHTML     string
	ContentMarkdown string
	Time            CurrentTime
	Owner           string
}

func NewPost(Id, Title, ContentHTML, ContentMarkdown string, Time CurrentTime, Owner string) *Post {
	return &Post{Id, Title, ContentHTML, ContentMarkdown, Time, Owner}
}
