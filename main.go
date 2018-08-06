package main

import (
	modules "./modules"
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
	postID := r.FormValue("postID")
	postTitle := r.FormValue("postTitle")
	postText := r.FormValue("postText")

	var post *modules.Post
	if postID != "" {
		post = posts[postID]
		post.title = postTitle
		post.content = postText
	} else {
		post := modules.newPost(postID, postTitle, postText)
		posts[post.id] = post
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
