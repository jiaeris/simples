package web

import (
	"net/http"
	"html/template"
	"fmt"
	"time"
)

func openServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			panic(err)
		}
		fmt.Println(t.Name(), time.Now().Format("2006-01-02"))
		data := struct {
			Title string
			Name  string
		}{Title: "标题", Name: "homepage"}

		t.Execute(w, data)
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("HelloC"))
	})

	s := http.Server{
		Addr:        ":8080",
		ReadTimeout: 10,
	}
	s.ListenAndServe()

}
