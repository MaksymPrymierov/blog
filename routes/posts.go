package routes

import (
	"net/http"
	"strings"

	"github.com/MaksymPrymierov/blog/db/documents"
	"github.com/MaksymPrymierov/blog/models"
	"github.com/MaksymPrymierov/blog/models/data"
	"github.com/MaksymPrymierov/blog/utils"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

/* Render write template */
func WriteHandler(rnd render.Render, r *http.Request) {
	/* Init user data and check user session */
	userData, err := getPublicCurrentUserData(r)
	if err != nil {
		getErrorHandler(rnd, 1)
		return
	}

	/* Init post data */
	post := models.Post{}

	/* Init PostsData */
	data := data.PostsData{post, userData}

	/* Render html template */
	rnd.HTML(200, "write", data)
}

/* Save post in database */
func CreatePostHandler(rnd render.Render, r *http.Request) {
	/* Get data in current user session */
	userData, err := getPublicCurrentUserData(r)
	if err != nil {
		getErrorHandler(rnd, 1)
		return
	}

	/* Init and setting html cleaner */
	p := bluemonday.NewPolicy()
	p.AllowStandardURLs()
	p.AllowElements("br")

	/* Write data of form */
	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("contentMarkdown")

	/* Clear html tegs in text which textarea */
	contentMarkdown = p.Sanitize(contentMarkdown)

	/* Convert markdown tegs in html tegs */
	contentHTML := blackfriday.Run([]byte(contentMarkdown))
	contentHTML = []byte(strings.Replace(string(contentHTML), "\"", "\"", -1))

	/* Get current time */
	currentTime := models.GetCurrentTime()

	/* Init posts data */
	postDocument := documents.PostDocument{
		id,
		title,
		string(contentHTML),
		contentMarkdown,
		currentTime,
		userData.Username,
	}

	/* Write data posts in data base */
	if id != "" {
		postsCollection.UpdateId(id, postDocument)
	} else {
		id = utils.GenerateNameId(title)
		postDocument.Id = id
		err := postsCollection.Insert(postDocument)
		for err != nil {
			id = id + "c"
			postDocument.Id = id
			err = postsCollection.Insert(postDocument)
		}
	}

	/* Redirect in main page */
	rnd.Redirect("/")
}

/* Render edite template */
func EditPostHandler(rnd render.Render, params martini.Params, r *http.Request) {
	userData, err := getPublicCurrentUserData(r)
	if err != nil {
		getErrorHandler(rnd, 1)
		return
	}

	/* Init Post data, and check post id */
	id := params["id"]
	postDocument := documents.PostDocument{}
	err = postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		getErrorHandler(rnd, 7)
		return
	}
	post := models.Post{
		postDocument.Id,
		postDocument.Title,
		postDocument.ContentHtml,
		postDocument.ContentMarkdown,
		postDocument.Time,
		postDocument.Owner,
	}

	/* Replate html teg <br> on symbol '\n' */
	//	post.ContentMarkdown = strings.Replace(post.ContentMarkdown, "<br>", "\n", -1)

	/* Init PostsData */
	data := data.PostsData{post, userData}

	/* Render html template */
	rnd.HTML(200, "write", data)
}

/* Render read template */
func ReadPostHandler(rnd render.Render, params martini.Params, r *http.Request) {
	/* Init user data */
	userData, _ := getPublicCurrentUserData(r)

	/* Init post data and check post id */
	id := params["id"]
	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		getErrorHandler(rnd, 7)
		return
	}
	post := models.Post{
		postDocument.Id,
		postDocument.Title,
		postDocument.ContentHtml,
		postDocument.ContentMarkdown,
		postDocument.Time,
		postDocument.Owner,
	}

	/* Init PostsData */
	data := data.PostsData{post, userData}

	rnd.HTML(200, "read", data)
}

/* Delete post from database */
func DeletePostHandler(rnd render.Render, params martini.Params, r *http.Request) {
	/* Check user session */
	if getCurrentUserId(r) == "" {
		getErrorHandler(rnd, 1)
		return
	}

	/* Check post id and save id in "params" */
	id := params["id"]
	if id == "" {
		getErrorHandler(rnd, 7)
		return
	}

	/* Remove post */
	postsCollection.RemoveId(id)

	/* Redirect in main page */
	rnd.Redirect("/")
}
