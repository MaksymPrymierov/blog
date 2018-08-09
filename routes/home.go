package routes

import (
	"fmt"
	"net/http"

	"github.com/martini-contrib/render"
	"labix.org/v2/mgo"

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

func IndexHandler(rnd render.Render, r *http.Request) {
	cookie, _ := r.Cookie(COOKIE_NAME)
	if cookie != nil {
		fmt.Println(inMemorySession.Get(cookie.Value))
	}

	postDocuments := []documents.PostDocument{}
	postsCollection.Find(nil).All(&postDocuments)

	posts := []models.Post{}
	for _, doc := range postDocuments {
		post := models.Post{doc.Id, doc.Title, doc.ContentHtml, doc.ContentMarkdown}
		posts = append(posts, post)
	}

	rnd.HTML(200, "index", posts)
}
