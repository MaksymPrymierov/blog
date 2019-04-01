package models

/* Structure for posts */
type Post struct {
	Id              string
	Title           string
	ContentHTML     string
	ContentMarkdown string
	Time            CurrentTime
	Owner           string
	Type            string
}

/* Init */
func NewPost(Id, Title, ContentHTML, ContentMarkdown string, Time CurrentTime, Owner, Type string) *Post {
	return &Post{Id, Title, ContentHTML, ContentMarkdown, Time, Owner, Type}
}

type Comment struct {
	Id      string
	PostId  string
	Title   string
	Content string
	Time    CurrentTime
	Owner   string
}

func NewComment(Id, PostId, Title, Content string, Time CurrentTime, Owner string) *Comment {
	return &Comment{Id, PostId, Title, Content, Time, Owner}
}
