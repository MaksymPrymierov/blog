package modules

type Post struct {
	Id    string
	Title string
	Text  string
}

func NewPost(Id, Title, Text string) *Post {
	return &Post{Id, Title, Text}
}
