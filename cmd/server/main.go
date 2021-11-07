package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/MrWormHole/perseusCMS/pkg/http/rest"
)

func main() {
	t := template.Must(template.ParseGlob("./templates/*.go.tmpl"))
	t = template.Must(t.ParseGlob("./templates/components/*.go.tmpl"))

	log.Fatal(http.ListenAndServe(":8080", rest.Handler(t)))
}
