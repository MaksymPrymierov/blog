package routes

import (
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/MaksymPrymierov/blog/models/data"
)

/* Array for errors */
var errors = []string{
	"Unknow error", // id 0
	"To access this page, you need to login.",              // id 1
	"You are already authorized.",                          // id 2
	"Not a valid password or login.",                       // id 3
	"This login is already taken.",                         // id 4
	"This email is already taken.",                         // id 5
	"You do not have sufficient rights to view this page.", // id 6
	"Post not found.",                                      // id 7
	"Incorrect login.",                                     // id 8
	"Incorrect email.",                                     // id 9
	"Incorrect password",                                   // id 10
}

/* Error handlers */
func ErrorHandler(rnd render.Render, params martini.Params, r *http.Request) {
	id, err := strconv.Atoi(params["id"])
	if err != nil || id >= len(errors) {
		rnd.HTML(200, "error", errors[0])
		return
	}

	userData, _ := getPublicCurrentUserData(r)

	data := data.MessageData{userData, errors[id]}

	rnd.HTML(200, "error", data)
}

func getErrorHandler(rnd render.Render, id int) {
	rnd.Redirect("/error/" + strconv.Itoa(id))
}
