package routes

import (
	"net/http"

	"github.com/connor41/blog/db/users"
	"github.com/connor41/blog/models/data"
	"github.com/connor41/blog/utils"
	"github.com/martini-contrib/render"
)

func checkUserData(r *http.Request) (int, string) {
	if getCurrentUserId(r) != "" {
		return 2, "null"
	}

	username := r.FormValue("username")

	if utils.CheckLen(username, 4, 30) != true {
		return 8, "null"
	}

	return 0, username
}

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
	errCode, username := checkUserData(r)
	if errCode != 0 {
		getErrorHandler(rnd, errCode)
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
