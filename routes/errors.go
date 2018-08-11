package routes

import (
	"github.com/martini-contrib/render"
)

const (
	notAuth     = "To access this page, you need to login."
	alreadyAuth = "You are already authorized."
	errAuth     = "Not a valid password or login."
	errRegLogin = "This login is already taken."
	errRegEmail = "This email is already taken."
	notPerm     = "You do not have sufficient rights to view this page."
)

func NotAuthHandler(rnd render.Render) {
	rnd.HTML(200, "error", notAuth)
}

func AlreadyAuthHandler(rnd render.Render) {
	rnd.HTML(200, "error", alreadyAuth)
}

func ErrAuthHandler(rnd render.Render) {
	rnd.HTML(200, "error", errAuth)
}

func ErrRegLoginHandler(rnd render.Render) {
	rnd.HTML(200, "error", errRegLogin)
}

func ErrRegEmailHandler(rnd render.Render) {
	rnd.HTML(200, "error", errRegEmail)
}

func NotPermHandler(rnd render.Render) {
	rnd.HTML(200, "error", notPerm)
}
