package main

import (
	"html/template"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"./routes"
)

func unescape(x string) interface{} {
	return template.HTML(x)
}

func main() {
	routes.Init()
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

	m.Get("/", routes.IndexHandler)
	m.Get("/write", routes.WriteHandler)
	m.Post("/createPost", routes.CreatePostHandler)
	m.Get("/editPost/:id", routes.EditPostHandler)
	m.Get("/readPost/:id", routes.ReadPostHandler)
	m.Get("/deletePost/:id", routes.DeletePostHandler)
	m.Get("/login", routes.GetLoginHandler)
	m.Post("/login", routes.PostLoginHandler)
	m.Get("/register", routes.GetRegisterHandler)
	m.Post("/register", routes.PostRegisterHandler)
	m.Get("/notPerm", routes.NotPermHandler)
	m.Get("/alreadyAuth", routes.AlreadyAuthHandler)
	m.Get("/exit", routes.ExitSessionHandler)

	m.RunOnAddr(":80")
	m.Run()
}
