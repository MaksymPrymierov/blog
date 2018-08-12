package routes

import (
	"net/http"
	"time"

	"github.com/martini-contrib/render"

	"../utils"
)

/* Render login template */
func GetLoginHandler(rnd render.Render, r *http.Request) {
	/* Check user session */
	if getCurrentUserId(r) != "" {
		getErrorHandler(rnd, 2)
		return
	}

	/* Render html template */
	rnd.HTML(200, "login", nil)
}

/* Create new user session */
func PostLoginHandler(rnd render.Render, r *http.Request, w http.ResponseWriter) {
	/* Check user sesssion */
	if getCurrentUserId(r) != "" {
		getErrorHandler(rnd, 2)
		return
	}

	/* Verification login */
	username := r.FormValue("username")
	user, err := getPrivateUserData(utils.GenerateNameId(username))
	if err != nil {
		getErrorHandler(rnd, 3)
		return
	}

	/* Verification password */
	password := r.FormValue("password")
	if user.Password != password {
		getErrorHandler(rnd, 3)
		return
	}

	/* Create new user session */
	sessionId := inMemorySession.Init(username)
	cookie := &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, cookie)

	/* Redirect in main page */
	rnd.Redirect("/")
}

/* Exit which user session */
func ExitSessionHandler(rnd render.Render, r *http.Request) {
	/* Check user session */
	if getCurrentUserId(r) == "" {
		getErrorHandler(rnd, 1)
		return
	}

	/* Delete user session */
	cookie, _ := r.Cookie(COOKIE_NAME)
	inMemorySession.Delete(cookie.Value)

	/* Redirect in main page */
	rnd.Redirect("/")
}
