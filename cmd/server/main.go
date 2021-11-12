package main

import (
	"github.com/MrWormHole/perseusCMS/http/rest"
	"html/template"
	"log"
	"net/http"
)

func main() {
	t := template.Must(template.ParseGlob("./web/templates/*.go.tmpl"))
	t = template.Must(t.ParseGlob("./web/templates/components/*.go.tmpl"))
	t = template.Must(t.ParseGlob("./web/templates/pages/*.go.tmpl"))

	log.Fatal(http.ListenAndServe(":8080", rest.Handler(t)))
}
