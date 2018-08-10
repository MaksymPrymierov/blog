package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"

	"../data"
	"../db/documents"
	"../models"
	"../utils"
)

func WriteHandler(rnd render.Render, r *http.Request) {
	username := protect(r)
	if username == "" {
		rnd.Redirect("/notPerm")
		return
	}

	post := models.Post{}

	date := data.PostsData{post, username}

	rnd.HTML(200, "write", date)
}

func CreatePostHandler(rnd render.Render, r *http.Request) {
	c := protect(r)
	if c == "" {
		rnd.Redirect("/notPerm")
		return
	}

	p := bluemonday.NewPolicy()
	p.AllowStandardURLs()
	p.AllowElements("br")

	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("contentMarkdown")
	contentMarkdown = p.Sanitize(contentMarkdown)

	contentHTML := blackfriday.Run([]byte(contentMarkdown))
	contentHTML = []byte(strings.Replace(string(contentHTML), "\n", " <br> ", -1))

	postDocument := documents.PostDocument{id, title, string(contentHTML), contentMarkdown}
	if id != "" {
		fmt.Println("old post")
		postsCollection.UpdateId(id, postDocument)
	} else {
		fmt.Println("new post")
		id = utils.GenerateNameId(title)
		postDocument.Id = id
		err := postsCollection.Insert(postDocument)
		for err != nil {
			id = id + "c"
			postDocument.Id = id
			err = postsCollection.Insert(postDocument)
		}
	}

	rnd.Redirect("/")
}

func EditPostHandler(rnd render.Render, params martini.Params, r *http.Request) {
	username := protect(r)
	if username == "" {
		rnd.Redirect("/notPerm")
		return
	}

	id := params["id"]

	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		rnd.Redirect("/")
		return
	}
	post := models.Post{postDocument.Id, postDocument.Title, postDocument.ContentHtml, postDocument.ContentMarkdown}

	post.ContentMarkdown = strings.Replace(post.ContentMarkdown, "<br>", "\n", -1)

	date := data.PostsData{post, username}

	rnd.HTML(200, "write", date)
}

func ReadPostHandler(rnd render.Render, params martini.Params, r *http.Request) {
	username := protect(r)

	id := params["id"]

	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		rnd.Redirect("/")
		return
	}

	post := models.Post{postDocument.Id, postDocument.Title, postDocument.ContentHtml, postDocument.ContentMarkdown}

	date := data.PostsData{post, username}

	rnd.HTML(200, "read", date)
}

func DeletePostHandler(rnd render.Render, params martini.Params, r *http.Request) {
	c := protect(r)
	if c == "" {
		rnd.Redirect("/notPerm")
		return
	}

	id := params["id"]
	if id == "" {
		rnd.Redirect("/")
		return
	}

	postsCollection.RemoveId(id)

	rnd.Redirect("/")
}
