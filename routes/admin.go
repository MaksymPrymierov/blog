package routes

import (
	"net/http"

	"github.com/martini-contrib/render"

	"../db/users"
)

func AdminHandler(rnd render.Render, r *http.Request) {
	username := protect(r)
	if username == "" {
		rnd.Redirect("http://google.com")
	}

	thisUser := users.UsersTable{}
	usersTables.FindId(username).One(&thisUser)
	if thisUser.Permission != "admin" {
		rnd.Redirect("http://google.com")
	}

	rnd.HTML(200, "admin", nil)
}
