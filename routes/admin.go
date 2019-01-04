package routes

import (
	"net/http"

	"github.com/connor41/blog/models"
	"github.com/connor41/blog/models/data"
	"github.com/connor41/blog/utils"
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

	var info data.AdminInfoServerData

	info.Pages = data.AdminPages{
		true, false,
	}
	info.UserData = models.Users{
		userData.Id, "null", userData.Username, "null", userData.Permission,
	}
	info.Day, info.Hour, info.Minute = utils.GetUptimeServer()

	/* Render html template */
	rnd.HTML(200, "admin", info)
}

func AdminUsersHandler(rnd render.Render, r *http.Request) {
	userData, err := getPrivateCurrentUserData(r)
	if err != nil {
		getErrorHandler(rnd, 6)
	}

	var info data.AdminInfoServerData

	info.Pages = data.AdminPages{
		false, true,
	}
	info.UserData = userData
	info.Day, info.Hour, info.Minute = 0, 0, 0

	rnd.HTML(200, "admin", info)
}
