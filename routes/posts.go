package routes

import (
	"net/http"
	"strings"

	"github.com/connor41/blog/db/documents"
	"github.com/connor41/blog/models"
	"github.com/connor41/blog/models/data"
	"github.com/connor41/blog/utils"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/russross/blackfriday.v2"
)

/* Render write template */
func WriteHandler(rnd render.Render, r *http.Request) {
	/* Init user data and check user session */
	userData, err := getPublicCurrentUserData(r)
	if err != nil {
		getErrorHandler(rnd, 1)
		return
	}

	if userData.Permission != "admin" {
		getErrorHandler(rnd, 6)
		return
	}

	/* Init post data */
	post := models.Post{}
	comment := []models.Comment{}

	/* Init PostsData */
	data := data.PostsData{post, userData, comment}

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

	if userData.Permission != "admin" {
		getErrorHandler(rnd, 6)
		return
	}

	/* Init and setting html cleaner */
	p := bluemonday.NewPolicy()
	p.AllowStandardURLs()
	p.AllowElements("br")

	/* Write data of form */
	id := r.FormValue("id")
	title := r.FormValue("title")
	typePost := r.FormValue("type")
	contentMarkdown := r.FormValue("contentMarkdown")

	/* Text preprocessing */
	contentMarkdown = strings.Replace(string(contentMarkdown), "<", "&lt;", -1)
	contentMarkdown = strings.Replace(string(contentMarkdown), ">", "&gt;", -1)

	/* Clear html tegs in text which textarea */
	contentMarkdown = p.Sanitize(contentMarkdown)

	/* Convert markdown tegs in html tegs */
	contentHTML := blackfriday.Run([]byte(contentMarkdown))
	contentHTML = []byte(strings.Replace(string(contentHTML), "amp;", "", -1))

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
		typePost,
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

/* Save post in database */

/* Render edite template */
func EditPostHandler(rnd render.Render, params martini.Params, r *http.Request) {
	userData, err := getPublicCurrentUserData(r)
	if err != nil {
		getErrorHandler(rnd, 1)
		return
	}

	if userData.Permission != "admin" {
		getErrorHandler(rnd, 6)
		return
	}

	/* Init Post data, and check post id */
	post, err := getPostData(params["id"])
	if err != nil {
		getErrorHandler(rnd, 7)
		return
	}

	commentsDocument := []documents.CommentsDocument{}
	commentCollection.FindId(params["id"]).All(&commentsDocument)

	comments := []models.Comment{}
	for _, doc := range commentsDocument {
		comment := models.Comment{
			doc.Id,
			doc.PostId,
			doc.Title,
			doc.Content,
			doc.Time,
			doc.Owner,
		}
		comments = append(comments, comment)
	}

	/* Text preprocessing */
	post.ContentMarkdown = strings.Replace(post.ContentMarkdown, "&lt;", "<", -1)
	post.ContentMarkdown = strings.Replace(post.ContentMarkdown, "&gt;", ">", -1)
	post.ContentMarkdown = strings.Replace(post.ContentMarkdown, "&#34;", "\"", -1)

	/* Init PostsData */
	data := data.PostsData{post, userData, comments}

	/* Render html template */
	rnd.HTML(200, "write", data)
}

/* Render read template */
func ReadPostHandler(rnd render.Render, params martini.Params, r *http.Request) {
	/* Init user data */
	userData, _ := getPublicCurrentUserData(r)

	/* Init post data and check post id */
	post, err := getPostData(params["id"])
	if err != nil {
		getErrorHandler(rnd, 7)
		return
	}

	commentsDocument := []documents.CommentsDocument{}
	commentCollection.Find(bson.M{"_postId": params["id"]}).All(&commentsDocument)

	comments := []models.Comment{}
	for _, doc := range commentsDocument {
		comment := models.Comment{
			doc.Id,
			doc.PostId,
			doc.Title,
			doc.Content,
			doc.Time,
			doc.Owner,
		}
		comments = append(comments, comment)
	}

	/* Init PostsData */
	data := data.PostsData{post, userData, comments}

	rnd.HTML(200, "read", data)
}

/* Delete post from database */
func DeletePostHandler(rnd render.Render, params martini.Params, r *http.Request) {
	/* Check user session */
	userData, err := getPublicCurrentUserData(r)
	if err != nil {
		getErrorHandler(rnd, 1)
		return
	}

	if userData.Permission != "admin" {
		getErrorHandler(rnd, 6)
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

func getPostData(id string) (models.Post, error) {
	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		post := models.Post{}
		return post, err
	}

	post := models.Post{
		postDocument.Id,
		postDocument.Title,
		postDocument.ContentHtml,
		postDocument.ContentMarkdown,
		postDocument.Time,
		postDocument.Owner,
		postDocument.Type,
	}

	return post, nil
}

func DeleteCommentHandler(rnd render.Render, params martini.Params, r *http.Request) {
	_, err := getPublicCurrentUserData(r)
	if err != nil {
		getErrorHandler(rnd, 1)
		return
	}

	id := params["id"]
	if id == "" {
		getErrorHandler(rnd, 7)
		return
	}

	commentCollection.RemoveId(id)

	rnd.Redirect("/")
}

func CreateCommentHandler(rnd render.Render, r *http.Request) {
	/* Get data in current user session */
	userData, err := getPublicCurrentUserData(r)
	if err != nil {
		getErrorHandler(rnd, 1)
		return
	}

	/* Write data of form */
	postId := r.FormValue("postId")
	title := r.FormValue("owner")
	content := r.FormValue("content")

	/* Get current time */
	currentTime := models.GetCurrentTime()

	/* Init posts data */
	commentDocument := documents.CommentsDocument{
		"",
		postId,
		title,
		content,
		currentTime,
		userData.Username,
	}

	/* Write data posts in data base */
	id := utils.GenerateNameId(postId)
	commentDocument.Id = id
	err = commentCollection.Insert(commentDocument)
	for err != nil {
		id = id + "c"
		commentDocument.Id = id
		err = commentCollection.Insert(commentDocument)
	}

	/* Redirect in main page */
	rnd.Redirect("/")
}
