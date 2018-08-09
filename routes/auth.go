package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/martini-contrib/render"

	"../db/users"
	"../utils"
)

func GetLoginHandler(rnd render.Render) {
	rnd.HTML(200, "login", nil)
}

func PostLoginHandler(rnd render.Render, r *http.Request, w http.ResponseWriter) {
	username := r.FormValue("username")
	id := utils.GenerateNameId(username)

	thisUser := users.UsersTable{}
	err := usersTables.FindId(id).One(&thisUser)
	if err != nil {
		fmt.Println("Нет такого юзера))")
		rnd.Redirect("/")
		return
	}

	password := r.FormValue("password")
	pass := thisUser.Password

	if pass != password {
		fmt.Println("Пароль не верный, иди нахуй")
		rnd.Redirect("/")
		return
	}

	sessionId := inMemorySession.Init(username)

	cookie := &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
	}

	http.SetCookie(w, cookie)

	rnd.Redirect("/")
}
