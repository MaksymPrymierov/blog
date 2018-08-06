package main

import (
	"./modules"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var posts map[string]*modules.Post
var postsID int

func indexHadler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "index", posts)
}

func writeHadler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "write", nil)
}

func createPostHadler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	text := r.FormValue("text")

	var post *modules.Post
	if id != "" {
		post = posts[id]
		post.Title = title
		post.Text = text
	} else {
		id = strconv.Itoa(postsID)
		post := modules.NewPost(id, title, text)
		posts[post.Id] = post
	}

	postsID = postsID + 1
	http.Redirect(w, r, "/", 302)
}

func editPostHadler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	id := r.FormValue("id")
	post, found := posts[id]
	if !found {
		http.NotFound(w, r)
		return
	}

	t.ExecuteTemplate(w, "write", post)
}

func readPostHadler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/read.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	id := r.FormValue("id")
	post, found := posts[id]
	if !found {
		http.NotFound(w, r)
		return
	}

	t.ExecuteTemplate(w, "read", post)
}

func deletePostHadler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.NotFound(w, r)
		return
	}

	delete(posts, id)

	http.Redirect(w, r, "/", 302)
}

func main() {
	fmt.Println("Listening on port :3000")

	posts = make(map[string]*modules.Post, 0)
	postsID = 1

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHadler)
	http.HandleFunc("/write", writeHadler)
	http.HandleFunc("/createPost", createPostHadler)
	http.HandleFunc("/editPost", editPostHadler)
	http.HandleFunc("/readPost", readPostHadler)
	http.HandleFunc("/deletePost", deletePostHadler)

	http.ListenAndServe(":3000", nil)
}
