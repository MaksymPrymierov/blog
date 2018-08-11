package routes

import (
	"net/http"

	"github.com/martini-contrib/render"

	"../db/users"
	"../utils"
)

/* Render register template */
func GetRegisterHandler(rnd render.Render, r *http.Request) {
	/* Check user session */
	if getCurrentUserId(r) != "" {
		rnd.Redirect("/alreadyAuth")
		return
	}

	/* Render html template */
	rnd.HTML(200, "register", nil)
}

/* Save user in database */
func PostRegisterHandler(rnd render.Render, r *http.Request) {
	/* Check user session */
	if getCurrentUserId(r) != "" {
		rnd.Redirect("/alreadyAuth")
		return
	}

	/* User data init */
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	perm := "user"
	id := utils.GenerateNameId(username)
	userTable := users.UsersTable{id, email, username, password, perm}

	/* Check login and save user data on data base */
	err := usersTables.Insert(userTable)
	if err != nil {
		rnd.Redirect("/errLogin")
	}

	/* Redirect in message */
	rnd.Redirect("/regSucc")
}
