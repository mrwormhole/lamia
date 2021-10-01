package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		//_, _ = w.Write([]byte("HELLO WORLD!"))
		err := t.ExecuteTemplate(w, "home", nil)
		if err != nil {
			panic(err)
		}
	})

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	return r
}
