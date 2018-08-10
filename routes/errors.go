package routes

import (
	"github.com/martini-contrib/render"
)

const (
	notAuth     = "To access this page, you need to login."
	alreadyAuth = "You are already authorized."
)

func NotPermHandler(rnd render.Render) {
	rnd.HTML(200, "error", notAuth)
}

func AlreadyAuthHandler(rnd render.Render) {
	rnd.HTML(200, "error", alreadyAuth)
}
