package rest

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/mrwormhole/lamia/pkg/invoicing"
)

const maxFileSize = 8 * 1024 * 1024

func Handler(svc invoicing.InvoiceService) http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token", "Authorization"},
		MaxAge:         300,
	}))

	r.Post("/generate/invoice", generateInvoice(svc))
	return r
}

func generateInvoice(svc invoicing.InvoiceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(maxFileSize)
		if err != nil {
			http.Error(w, "attached logo can not be bigger than 8Mb", http.StatusBadRequest)
			return
		}

		_, err = invoicing.FormToInvoice(r.MultipartForm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		rawReq, err := httputil.DumpRequest(r, true)
		if err != nil {
			panic(err)
		}
		log.Println(string(rawReq))
		log.Println("------------------------------------------------------------")

		_, _ = w.Write([]byte("invoice generated"))
	}
}
