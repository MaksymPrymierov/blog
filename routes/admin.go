package routes

import (
	"net/http"

	"github.com/connor41/blog/models/data"
	"github.com/martini-contrib/render"
)

/* Render admin template */
func AdminHandler(rnd render.Render, r *http.Request) {
	/* Check user session */
	userData, err := getPublicCurrentUserData(r)
	if err != nil {
		getErrorHandler(rnd, 1)
	}

	/* Check user permission */
	if userData.Permission != "admin" {
		getErrorHandler(rnd, 6)
	}

	data := data.AdminData{userData}

	/* Render html template */
	rnd.HTML(200, "admin", data)
}
