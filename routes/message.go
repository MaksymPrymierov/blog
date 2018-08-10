package routes

import (
	"github.com/martini-contrib/render"
)

const (
	regSucc = "Registration is successfully completed, you can log in."
)

func RegSuccHandler(rnd render.Render) {
	rnd.HTML(200, "message", regSucc)
}
