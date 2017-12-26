package fileserver

import (
	"net/http"
)

func open() {
	http.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir("./"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello c"))
	})

	s := &http.Server{
		Addr: ":10086",
	}
	s.ListenAndServe()
}
