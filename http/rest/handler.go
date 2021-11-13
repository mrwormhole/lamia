package rest

import (
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handler(t *template.Template) http.Handler {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))

	r.HandleFunc("/home", homeHandler(t))
	r.HandleFunc("/reload-browser", reloadBrowserHandler)

	return r
}

func homeHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// _, _ = w.Write([]byte("HELLO WORLD!"))
		err := t.ExecuteTemplate(w, "pages/home", nil)
		if err != nil {
			panic(err)
		}
	}
}

func reloadBrowserHandler(w http.ResponseWriter, req *http.Request) {
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer func() {
		_ = ws.Close()
	}()
	for {
		_, rawMsg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("[ERROR]: %v", err)
			break
		}
		log.Printf("[INFO]: %v", string(rawMsg))
	}
}
