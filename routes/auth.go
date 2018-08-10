package routes

import (
	"net/http"
	"time"

	"github.com/martini-contrib/render"

	"../db/users"
	"../utils"
)

func GetLoginHandler(rnd render.Render, r *http.Request) {
	s := protect(r)
	if s != "" {
		rnd.Redirect("/alreadyAuth")
		return
	}

	rnd.HTML(200, "login", nil)
}

func PostLoginHandler(rnd render.Render, r *http.Request, w http.ResponseWriter) {
	s := protect(r)
	if s != "" {
		rnd.Redirect("/alreadyAuth")
		return
	}

	username := r.FormValue("username")
	id := utils.GenerateNameId(username)

	thisUser := users.UsersTable{}
	err := usersTables.FindId(id).One(&thisUser)
	if err != nil {
		rnd.Redirect("/errAuth")
		return
	}

	password := r.FormValue("password")
	pass := thisUser.Password

	if pass != password {
		rnd.Redirect("errAuth")
		return
	}

	sessionId := inMemorySession.Init(username)

	cookie := &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, cookie)

	rnd.Redirect("/")
}

func ExitSessionHandler(rnd render.Render, r *http.Request) {
	s := protect(r)
	if s == "" {
		rnd.Redirect("/notAuth")
		return
	}

	cookie, _ := r.Cookie(COOKIE_NAME)
	inMemorySession.Delete(cookie.Value)

	rnd.Redirect("/")
}
