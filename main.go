package main

import (
	"./modules"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
)

var posts map[string]*modules.Post
var postsID int

func indexHandler(rnd render.Render) {
	rnd.HTML(200, "index", posts)
}

func writeHandler(rnd render.Render) {
	rnd.HTML(200, "write", nil)
}

func createPostHandler(rnd render.Render, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	text := r.FormValue("text")

	var post *modules.Post
	if id != "" {
		post = posts[id]
		post.Title = title
		post.Text = text
	} else {
		id = strconv.Itoa(postsID)
		post := modules.NewPost(id, title, text)
		posts[post.Id] = post
	}

	postsID = postsID + 1
	rnd.Redirect("/")
}

func editPostHandler(rnd render.Render, params martini.Params) {
	id := params["id"]
	post, found := posts[id]
	if !found {
		rnd.Redirect("/")
		return
	}

	rnd.HTML(200, "write", post)
}

func readPostHandler(rnd render.Render, params martini.Params) {
	id := params["id"]
	post, found := posts[id]
	if !found {
		rnd.Redirect("/")
		return
	}

	rnd.HTML(200, "read", post)
}

func deletePostHandler(rnd render.Render, params martini.Params) {
	id := params["id"]
	if id == "" {
		rnd.Redirect("/")
		return
	}

	delete(posts, id)

	rnd.Redirect("/")
}

func main() {
	posts = make(map[string]*modules.Post, 0)
	postsID = 1

	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		//	Funcs:           []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
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
