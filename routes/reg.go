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

	/* Write username */
	username := r.FormValue("username")

	/* Check username */
	_, checkUser := findUserOfData("_username", username)
	if checkUser == nil {
		rnd.Redirect("/errLogin")
		return
	}

	/* Write email */
	email := r.FormValue("email")

	/* Check email */
	_, checkEmail := findUserOfData("_email", email)
	if checkEmail == nil {
		rnd.Redirect("/errEmail")
		return
	}

	/* Write password */
	password := r.FormValue("password")

	/* Set permission */
	perm := "user"

	/* Init user data */
	id := utils.GenerateNameId(username)
	userTable := users.UsersTable{id, email, username, password, perm}

	/* Save user data in data base */
	err := usersTables.Insert(userTable)
	if err != nil {
		rnd.Redirect("/errLogin")
	}

	/* Redirect in message */
	rnd.Redirect("/regSucc")
}
