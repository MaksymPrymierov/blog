package routes

import (
	_ "fmt"
	"net/http"

	"github.com/martini-contrib/render"
	"labix.org/v2/mgo"

	"../data"
	"../db/documents"
	"../models"
	"../session"
)

var postsCollection *mgo.Collection
var usersTables *mgo.Collection
var inMemorySession *session.Session

const (
	COOKIE_NAME = "sessionId"
)

func Init() {
	inMemorySession = session.NewSession()

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	postsCollection = session.DB("blog").C("posts")
	usersTables = session.DB("blog").C("users")
}

func protect(r *http.Request) string {
	cookie, _ := r.Cookie(COOKIE_NAME)
	var currentSession string
	if cookie != nil {
		currentSession = (inMemorySession.Get(cookie.Value))
	}
	return currentSession
}

func IndexHandler(rnd render.Render, r *http.Request) {
	currentSession := protect(r)

	postDocuments := []documents.PostDocument{}
	postsCollection.Find(nil).All(&postDocuments)

	posts := []models.Post{}
	for _, doc := range postDocuments {
		post := models.Post{doc.Id, doc.Title, doc.ContentHtml, doc.ContentMarkdown, doc.Time}
		posts = append(posts, post)
	}

	date := data.IndexData{posts, currentSession}

	rnd.HTML(200, "index", date)
}
