package main

import (
	"./modules"
	"fmt"
	"html/template"
	"net/http"
)

var posts map[string]*modules.Post

func indexHadler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", posts)
}

func writeHadler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
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
		post := modules.NewPost(id, title, text)
		posts[post.Id] = post
	}

	http.Redirect(w, r, "/", 302)
}

func main() {
	fmt.Println("Listening on port :3000")

	posts = make(map[string]*modules.Post, 0)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHadler)
	http.HandleFunc("/write", writeHadler)
	http.HandleFunc("/createPost", createPostHadler)

	http.ListenAndServe(":3000", nil)
}
