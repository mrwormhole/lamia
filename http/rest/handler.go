package rest

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(t *template.Template) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		// _, _ = w.Write([]byte("HELLO WORLD!"))
		err := t.ExecuteTemplate(w, "pages/home", nil)
		if err != nil {
			panic(err)
		}
	})
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))
	return r
}
