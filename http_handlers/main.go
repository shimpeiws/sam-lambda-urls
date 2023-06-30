package http_handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func DefaultHandler() (http.Handler, error) {
	return Handler()
}

func Handler() (http.Handler, error) {
	r := chi.NewRouter()
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		handleHello(w, r)
	})
	return r, nil
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Sprintf("r: %v", r)
	w.Write([]byte("OK!!!"))
}
