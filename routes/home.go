package routes

import (
	"html/template"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"labix.org/v2/mgo"

	"../data"
	"../db/documents"
	"../db/users"
	"../models"
	"../session"
)

var postsCollection *mgo.Collection
var usersTables *mgo.Collection
var inMemorySession *session.Session

const (
	COOKIE_NAME = "sessionId"
)

func unescape(x string) interface{} {
	return template.HTML(x)
}

func Init() *martini.ClassicMartini {
	inMemorySession = session.NewSession()

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	postsCollection = session.DB("blog").C("posts")

	usersTables = session.DB("blog").C("users")

	m := martini.Classic()

	unescapeFuncMap := template.FuncMap{"unescape": unescape}

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                         // Specify what path to load the templates from.
		Layout:     "layout",                            // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"},          // Specify extensions to load for templates.
		Funcs:      []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8",                             // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                                // Output human readable JSON
	}))
	return m
}

func protect(r *http.Request) string {
	cookie, _ := r.Cookie(COOKIE_NAME)
	var currentSession string
	if cookie != nil {
		currentSession = (inMemorySession.Get(cookie.Value))
	}
	return currentSession
}

func UpdateUserPermission(rnd render.Render) {
	thisUser := users.UsersTable{}
	usersTables.FindId("connor41").One(&thisUser)
	oldUser := thisUser
	thisUser.Permission = "admin"
	usersTables.Update(oldUser, thisUser)

	rnd.Redirect("http://google.com")
}

func IndexHandler(rnd render.Render, r *http.Request) {
	currentSession := protect(r)

	postDocuments := []documents.PostDocument{}
	postsCollection.Find(nil).All(&postDocuments)

	posts := []models.Post{}
	for _, doc := range postDocuments {
		post := models.Post{
			doc.Id,
			doc.Title,
			doc.ContentHtml,
			doc.ContentMarkdown,
			doc.Time,
			doc.Owner,
		}
		posts = append(posts, post)
	}

	date := data.IndexData{posts, currentSession}

	rnd.HTML(200, "index", date)
}
