package routes

import (
	"net/http"

	"github.com/connor41/blog/db/users"
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
	info.UserData = userData
	info.Day, info.Hour, info.Minute = utils.GetUptimeServer()

	/* Render html template */
	rnd.HTML(200, "admin", info)
}

func AdminUsersHandler(rnd render.Render, r *http.Request) {
	userData, err := getPublicCurrentUserData(r)
	if err != nil {
		getErrorHandler(rnd, 6)
	}

	var info data.AdminUsersData

	info.Pages = data.AdminPages{
		false, true,
	}
	info.UserData = userData

	usersTable := []users.UsersTable{}
	usersTables.Find(nil).All(&usersTable)
	for _, data := range usersTable {
		user := models.Users{
			data.Id,
			data.Email,
			data.Username,
			data.Password,
			data.Permission,
		}
		info.UsersData = append(info.UsersData, user)
	}

	rnd.HTML(200, "admin", info)
}
