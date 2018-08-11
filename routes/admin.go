package routes

import (
	"net/http"

	"github.com/martini-contrib/render"

	"../db/users"
)

func AdminHandler(rnd render.Render, r *http.Request) {
	username := protect(r)
	if username == "" {
		rnd.Redirect("/notAuth")
	}

	thisUser := users.UsersTable{}
	usersTables.FindId(username).One(&thisUser)
	if thisUser.Permission != "admin" {
		rnd.Redirect("/notPerm")
	}

	rnd.HTML(200, "admin", username)
}
