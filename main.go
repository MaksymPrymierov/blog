package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
	"labix.org/v2/mgo"

	"./db/documents"
	"./modules"
)

var postsCollection *mgo.Collection
var postsID int

func indexHandler(rnd render.Render) {
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
	rnd.HTML(200, "write", nil)
}

func createPostHandler(rnd render.Render, r *http.Request) {
	p := bluemonday.NewPolicy()
	p.AllowStandardURLs()
	p.AllowElements("br")

	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("contentMarkdown")
	contentMarkdown = p.Sanitize(contentMarkdown)
	contentMarkdown = strings.Replace(contentMarkdown, "\n", "<br>", -1)

	contentHTML := blackfriday.Run([]byte(contentMarkdown))

	postDocument := documents.PostDocument{id, title, string(contentHTML), contentMarkdown}
	if id != "" {
		postsCollection.UpdateId(id, postDocument)
	} else {
		id = strconv.Itoa(postsID)
		postDocument.Id = id
		postsCollection.Insert(postDocument)
	}

	postsID = postsID + 1
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

func unescape(x string) interface{} {
	return template.HTML(x)
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	postsCollection = session.DB("blog").C("posts")
	postsID = 1

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

	m.Run()
}
