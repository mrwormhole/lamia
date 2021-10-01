package main

import (
	"html/template"
	"net/http"
	"time"
)

var t *template.Template

func main() {
	router := SetupRouter()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    30 * time.Second,
		MaxHeaderBytes: 2048,
	}

	t = template.Must(template.ParseGlob("./templates/*.go.tmpl"))
	t = template.Must(t.ParseGlob("./templates/components/*.go.tmpl"))

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
