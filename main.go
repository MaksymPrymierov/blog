package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHadler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func main() {
	fmt.Println("Listening on port :3000")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHadler)

	http.ListenAndServe(":3000", nil)
}
