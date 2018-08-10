package routes

import (
	"fmt"
	"net/http"

	"github.com/martini-contrib/render"

	"../db/users"
	"../utils"
)

func GetRegisterHandler(rnd render.Render, r *http.Request) {
	c := protect(r)
	if c != "" {
		rnd.Redirect("/alreadyAuth")
		return
	}

	rnd.HTML(200, "register", nil)
}

func PostRegisterHandler(rnd render.Render, r *http.Request) {
	c := protect(r)
	if c != "" {
		rnd.Redirect("/alreadyAuth")
		return
	}

	id := ""
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	perm := ""

	userTable := users.UsersTable{id, email, username, password, perm}
	fmt.Println("createuser")
	id = utils.GenerateNameId(username)
	perm = "user"
	userTable.Id = id
	userTable.Permission = perm
	err := usersTables.Insert(userTable)
	if err != nil {
		rnd.Redirect("/errLogin")
	}

	rnd.Redirect("/regSucc")
}
