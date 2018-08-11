package routes

import (
	"github.com/martini-contrib/render"
)

/* Text for notice */
const (
	regSucc = "Registration is successfully completed, you can log in."
)

/* Notice handler */
func RegSuccHandler(rnd render.Render) {
	rnd.HTML(200, "message", regSucc)
}
