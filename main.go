package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
	"labix.org/v2/mgo"

	"./db/documents"
	"./modules"
	"./session"
	"./utils"
)

var postsCollection *mgo.Collection
var inMemorySession *session.Session

const (
	COOKIE_NAME = "sessionId"
)

func indexHandler(rnd render.Render, r *http.Request) {
	cookie, _ := r.Cookie(COOKIE_NAME)
	if cookie != nil {
		fmt.Println(inMemorySession.Get(cookie.Value))
	}

	postDocuments := []documents.PostDocument{}
	postsCollection.Find(nil).All(&postDocuments)

	posts := []modules.Post{}
	for _, doc := range postDocuments {
		post := modules.Post{doc.Id, doc.Title, doc.ContentHtml, doc.ContentMarkdown}
		posts = append(posts, post)
	}

	rnd.HTML(200, "index", posts)
}

func writeHandler(rnd render.Render) {
	post := modules.Post{}
	rnd.HTML(200, "write", post)
}

func createPostHandler(rnd render.Render, r *http.Request) {
	code := r.FormValue("code")
	if code != "дерпароль" {
		rnd.Redirect("/")
		return
	}
	p := bluemonday.NewPolicy()
	p.AllowStandardURLs()
	p.AllowElements("br")

	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("contentMarkdown")
	contentMarkdown = p.Sanitize(contentMarkdown)

	contentHTML := blackfriday.Run([]byte(contentMarkdown))
	contentHTML = []byte(strings.Replace(string(contentHTML), "\n", " <br> ", -1))

	postDocument := documents.PostDocument{id, title, string(contentHTML), contentMarkdown}
	if id != "" {
		fmt.Println("old post")
		postsCollection.UpdateId(id, postDocument)
	} else {
		fmt.Println("new post")
		id = utils.GenerateNameId(title)
		postDocument.Id = id
		err := postsCollection.Insert(postDocument)
		for err != nil {
			id = id + "c"
			postDocument.Id = id
			err = postsCollection.Insert(postDocument)
		}
	}

	rnd.Redirect("/")
}

func editPostHandler(rnd render.Render, params martini.Params) {
	id := params["id"]

	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		rnd.Redirect("/")
		return
	}
	post := modules.Post{postDocument.Id, postDocument.Title, postDocument.ContentHtml, postDocument.ContentMarkdown}

	post.ContentMarkdown = strings.Replace(post.ContentMarkdown, "<br>", "\n", -1)

	rnd.HTML(200, "write", post)
}

func readPostHandler(rnd render.Render, params martini.Params) {
	id := params["id"]

	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		rnd.Redirect("/")
		return
	}

	post := modules.Post{postDocument.Id, postDocument.Title, postDocument.ContentHtml, postDocument.ContentMarkdown}

	rnd.HTML(200, "read", post)
}

func deletePostHandler(rnd render.Render, params martini.Params) {
	id := params["id"]
	if id == "" {
		rnd.Redirect("/")
		return
	}

	postsCollection.RemoveId(id)

	rnd.Redirect("/")
}

func getLoginHandler(rnd render.Render) {
	rnd.HTML(200, "login", nil)
}

func postLoginHandler(rnd render.Render, r *http.Request, w http.ResponseWriter) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println(username)
	fmt.Println(password)

	sessionId := inMemorySession.Init(username)

	cookie := &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
	}

	http.SetCookie(w, cookie)

	rnd.Redirect("/")
}

func unescape(x string) interface{} {
	return template.HTML(x)
}

func main() {
	inMemorySession = session.NewSession()

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Now())

	postsCollection = session.DB("blog").C("posts")

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

	staticOpt := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOpt))

	m.Get("/", indexHandler)
	m.Get("/write", writeHandler)
	m.Post("/createPost", createPostHandler)
	m.Get("/editPost/:id", editPostHandler)
	m.Get("/readPost/:id", readPostHandler)
	m.Get("/deletePost/:id", deletePostHandler)
	m.Get("/login", getLoginHandler)
	m.Post("/login", postLoginHandler)

	m.Run()
}
