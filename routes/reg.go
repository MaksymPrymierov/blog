package routes

import (
	"net/http"

	"github.com/MaksymPrymierov/blog/db/users"
	"github.com/MaksymPrymierov/blog/models/data"
	"github.com/MaksymPrymierov/blog/utils"
	"github.com/martini-contrib/render"
)

/* Render register template */
func GetRegisterHandler(rnd render.Render, r *http.Request) {
	userData, err := getPublicCurrentUserData(r)
	if err == nil {
		getErrorHandler(rnd, 2)
		return
	}

	data := data.GeneralData{userData}

	/* Render html template */
	rnd.HTML(200, "register", data)
}

/* Save user in database */
func PostRegisterHandler(rnd render.Render, r *http.Request) {

	/* Check user session */
	if getCurrentUserId(r) != "" {
		getErrorHandler(rnd, 2)
		return
	}

	/* Write username */
	username := r.FormValue("username")

	/* Check len username */
	if utils.CheckLen(username, 4, 30) != true {
		getErrorHandler(rnd, 8)
		return
	}

	/* Check username */
	_, checkUser := findUserOfData("_username", username)
	if checkUser == nil {
		getErrorHandler(rnd, 4)
		return
	}

	/* Write email */
	email := r.FormValue("email")

	/* Check len email */
	if utils.CheckLen(email, 4, 60) != true {
		getErrorHandler(rnd, 9)
		return
	}

	/* Check email */
	_, checkEmail := findUserOfData("_email", email)
	if checkEmail == nil {
		getErrorHandler(rnd, 5)
		return
	}

	/* Write password */
	password := r.FormValue("password")

	/* Check len password */
	if utils.CheckLen(password, 4, 120) != true {
		getErrorHandler(rnd, 10)
		return
	}

	/* Set permission */
	perm := "user"

	/* Init user data */
	id := utils.GenerateNameId(username)
	userTable := users.UsersTable{id, email, username, password, perm}

	/* Save user data in data base */
	err := usersTables.Insert(userTable)
	if err != nil {
		getErrorHandler(rnd, 4)
	}

	/* Redirect in message */
	getMessageHandler(rnd, 1)
	return
}
