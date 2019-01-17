package routes

import (
	"net/http"
	"time"

	"github.com/connor41/blog/models/data"
	"github.com/connor41/blog/utils"
	"github.com/martini-contrib/render"
)

/* Render login template */
func GetLoginHandler(rnd render.Render, r *http.Request) {
	userData, err := getPublicCurrentUserData(r)
	/* Check user session */
	if err == nil {
		getErrorHandler(rnd, 2)
		return
	}

	data := data.GeneralData{userData}

	/* Render html template */
	rnd.HTML(200, "login", data)
}

/* Create new user session */
func PostLoginHandler(rnd render.Render, r *http.Request, w http.ResponseWriter) {
	errCode, username := checkUserData(r)
	if errCode != 0 {
		getErrorHandler(rnd, errCode)
	}

	/* Verification login */
	user, err := getPrivateUserData(utils.GenerateNameId(username))
	if err != nil {
		getErrorHandler(rnd, 3)
		return
	}

	/* Init password */
	password := r.FormValue("password")

	/* Check len password */
	if utils.CheckLen(password, 4, 120) != true {
		getErrorHandler(rnd, 8)
		return
	}

	/* Verification password */
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
