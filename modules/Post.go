package modules

type Post struct {
	Id              string
	Title           string
	ContentHTML     string
	ContentMarkdown string
}

func NewPost(Id, Title, ContentHTML, ContentMarkdown string) *Post {
	return &Post{Id, Title, ContentHTML, ContentMarkdown}
}
