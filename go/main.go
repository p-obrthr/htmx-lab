package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	tmpl, err := template.ParseFiles("../index.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Message string
		}{
			Message: "hello world",
		}
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("current time: " + time.Now().Format(time.RFC1123)))
	})

	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		someText := "some Text..."

		dynamicContent := fmt.Sprintf(`
			<div>
				<h3>Dynamic Content</h3>
				<p>%s</p>
			</div>`, someText)

		w.Write([]byte(dynamicContent))
	})

	fmt.Println("server on port:8080....")
	http.ListenAndServe(":8080", nil)
}
