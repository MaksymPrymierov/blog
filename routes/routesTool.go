package routes

import (
	"html/template"
	"net/http"

	"github.com/connor41/blog/db/users"
	"github.com/connor41/blog/models"
	"github.com/connor41/blog/session"
	"github.com/connor41/blog/utils"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2/bson"
	"labix.org/v2/mgo"
)

/* Global variables */
var postsCollection *mgo.Collection  // Variable for posts
var usersTables *mgo.Collection      // Variable for user
var inMemorySession *session.Session // Variable for session
var commentCollection *mgo.Collection

/* Global const */
const (
	COOKIE_NAME = "sessionId" // Const for cookie name
)

/* Init default server data */
func Init() *martini.ClassicMartini {
	/* Init session */
	inMemorySession = session.NewSession()

	/* Connect to database */
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	/* Connect collection from database */
	postsCollection = session.DB("blog").C("posts")
	usersTables = session.DB("blog").C("users")
	commentCollection = session.DB("blog").C("comments")

	/* Init martini framework */
	m := martini.Classic()

	/* Init martini render */
	unescapeFuncMap := template.FuncMap{"unescape": unescape, "checkGroup": checkGroup}
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                         // Specify what path to load the templates from.
		Layout:     "layout",                            // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"},          // Specify extensions to load for templates.
		Funcs:      []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8",                             // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                                // Output human readable JSON
	}))

	return m
}

func updateUserPermission(id string) {
	thisUser := users.UsersTable{}
	usersTables.FindId(id).One(&thisUser)
	oldUser := thisUser

	if thisUser.Permission == "admin" {
		thisUser.Permission = "user"
	} else {
		thisUser.Permission = "admin"
	}

	usersTables.Update(oldUser, thisUser)
}

func updateBan(id string) {
	thisUser := users.UsersTable{}
	usersTables.FindId(id).One(&thisUser)
	oldUser := thisUser

	if thisUser.Permission == "banned" {
		thisUser.Permission = "user"
	} else {
		thisUser.Permission = "banned"
	}

	usersTables.Update(oldUser, thisUser)
}

/* Function return id user of active current session */
func getCurrentUserId(r *http.Request) string {
	/* Get user name of current session */
	cookie, _ := r.Cookie(COOKIE_NAME)
	var username string
	if cookie != nil {
		username = inMemorySession.Get(cookie.Value)
	}

	/* Convert username on userId */
	userId := utils.GenerateNameId(username)

	return userId
}

/* Func return all user data on id and error */
func getPrivateUserData(userId string) (models.Users, error) {
	/* Init type Users */
	var privateUserData models.Users

	/* Get and check user on id */
	user := users.UsersTable{}
	err := usersTables.FindId(userId).One(&user)
	if err != nil {
		return privateUserData, err
	}

	/* Init all user data */
	privateUserData = models.Users{userId, user.Email, user.Username, user.Password, user.Permission}

	return privateUserData, nil
}

/* Func return public user data on id and error */
func getPublicUserData(userId string) (models.PublicUsersData, error) {
	/* Init type PublicUsersData */
	var publicUserData models.PublicUsersData

	/* Get and check user on id */
	user := users.UsersTable{}
	err := usersTables.FindId(userId).One(&user)
	if err != nil {
		return publicUserData, err
	}

	/* Init public user data */
	publicUserData = models.PublicUsersData{userId, user.Username, user.Permission}

	return publicUserData, nil
}

/* Function return all information about user current session */
func getPublicCurrentUserData(r *http.Request) (models.PublicUsersData, error) {
	return getPublicUserData(getCurrentUserId(r))
}

/* Function return public information about user current session */
func getPrivateCurrentUserData(r *http.Request) (models.Users, error) {
	return getPrivateUserData(getCurrentUserId(r))
}

/* Find user of custom data */
func findUserOfData(typeData, data string) (interface{}, error) {
	result := users.UsersTable{}
	err := usersTables.Find(bson.M{typeData: data}).One(&result)
	return result, err
}

func GetInfoUsers() int {
	count, _ := usersTables.Count()
	return count
}
