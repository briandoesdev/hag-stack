package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/briandoesdev/hag-stack/internal/components"
	"github.com/briandoesdev/hag-stack/internal/pages"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	port := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	router := chi.NewRouter()

	// Router Middleware
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second)) // timeout request after 60 seconds

	// Routes
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pages.Index().Render(r.Context(), w)
	})

	router.Post("/clicked", func(w http.ResponseWriter, r *http.Request) {
		is_htmx := r.Header.Get("HX-Request")
		log.Println("HTMX Request:", is_htmx)
		r.ParseForm()

		components.Welcome(r.FormValue("name")).Render(r.Context(), w)
	})

	router.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	log.Printf("Server started on http://localhost:%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}
