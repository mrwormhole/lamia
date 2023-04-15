package rest

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

const maxFileSize = 8 * 1024 * 1024

func Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token", "Authorization"},
		MaxAge:         300,
	}))

	r.Post("/generate/invoice", generateInvoice)
	return r
}

func generateInvoice(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(maxFileSize)
	if err != nil {
		panic(err)
	}

	if r.MultipartForm != nil {
		for k, v := range r.MultipartForm.Value {
			log.Printf("multipart form values = %v:%v \n", k, v)
		}
		for k, v := range r.MultipartForm.File {
			log.Printf("multipart form files = %v:%v \n", k, v)
		}
	}

	rawReq, err := httputil.DumpRequest(r, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(rawReq))
	log.Println("------------------------------------------------------------")

	_, _ = w.Write([]byte("invoice generated"))
}
