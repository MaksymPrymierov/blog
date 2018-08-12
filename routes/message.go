package routes

import (
	"strconv"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

/* Array for messages */
var messages = []string{
	"Unknow meggase",
	"Registration is successfully completed, you can log in.",
}

/* Message handler */
func MessageHandler(rnd render.Render, params martini.Params) {
	id, err := strconv.Atoi(params["id"])
	if err != nil || id >= len(errors) {
		rnd.HTML(200, "message", messages[0])
		return
	}

	rnd.HTML(200, "message", messages[id])
}

func getMessageHandler(rnd render.Render, id int) {
	rnd.Redirect("/message/" + strconv.Itoa(id))
}
