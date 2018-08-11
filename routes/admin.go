package routes

import (
	"net/http"

	"github.com/martini-contrib/render"
)

/* Render admin template */
func AdminHandler(rnd render.Render, r *http.Request) {
	/* Check user session */
	userData, err := getPublicCurrentUserData(r)
	if err != nil {
		rnd.Redirect("/notAuth")
	}

	/* Check user permission */
	if userData.Permission != "admin" {
		rnd.Redirect("/notPerm")
	}

	/* Render html template */
	rnd.HTML(200, "admin", userData.Username)
}
