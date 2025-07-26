package main

import (
	"fmt"
	"log"
	"my-go-project/internal"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	app := &application{
		config: config{
			addr: internal.GetStringEnv("ADDR", ":8080"),
		},
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	example := internal.NewExample("My Example")
	fmt.Println(example.DoSomething())

	fmt.Printf("Server is running on %s\n", app.config.addr)
	mux := app.mount()
	log.Fatal(app.run(mux))
}
