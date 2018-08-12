package main

import (
	"github.com/go-martini/martini"

	"./routes"
)

func main() {
	/* Init default server data */
	m := routes.Init()

	/* Static file */
	staticOpt := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOpt))

	/* Home Routes */
	m.Get("/", routes.IndexHandler)

	/* Posts Routes */
	m.Get("/write", routes.WriteHandler)
	m.Post("/createPost", routes.CreatePostHandler)
	m.Get("/editPost/:id", routes.EditPostHandler)
	m.Get("/readPost/:id", routes.ReadPostHandler)
	m.Get("/deletePost/:id", routes.DeletePostHandler)

	/* Auth Routes */
	m.Get("/login", routes.GetLoginHandler)
	m.Post("/login", routes.PostLoginHandler)
	m.Get("/exit", routes.ExitSessionHandler)

	/* Reg Routes */
	m.Get("/register", routes.GetRegisterHandler)
	m.Post("/register", routes.PostRegisterHandler)

	/* Error Routes */
	m.Get("/error/:id", routes.ErrorHandler)

	/* Message Routes */
	m.Get("/message/:id", routes.MessageHandler)

	/* Admin Panel Routes */
	m.Get("/admin", routes.AdminHandler)

	/* Listen port and run server*/
	m.RunOnAddr(":80")
	m.Run()
}
