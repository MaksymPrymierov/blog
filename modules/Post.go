package modules

type Post struct {
	id    string
	title string
	text  string
}

func newPost(id, title, text string) *Post {
	return &Post{id, title, text}
}
