package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}
type config struct {
	addr string
}

func (app *application) run(mux http.Handler) error {
	// app.routes() would typically set up your routes here
	// mux.Handle("/", app.routes())

	srver := &http.Server{
		Addr: app.config.addr,
		// Handler: app.routes(),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,                                   // Example of setting a read timeout
		WriteTimeout: 10 * time.Second,                                   // Example of setting a write timeout
		IdleTimeout:  120 * time.Second,                                  // Example of setting an idle timeout
		ErrorLog:     log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime), // Example of setting an error logger
	}
	return srver.ListenAndServe()
}
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)                 // Generate a unique request ID for each request
	r.Use(middleware.RealIP)                    // Use the real IP address of the client
	r.Use(middleware.Timeout(60 * time.Second)) // Set a timeout for requests
	r.Use(middleware.Recoverer)                 // Recover from panics and log them
	r.Use(middleware.Logger)
	r.Get("/", app.hello)
	r.Get("/health", app.healthCheck)

	// You can add more routes as needed
	// mux.HandleFunc("/api/v1/resource", app.handleResource)

	return r
}
