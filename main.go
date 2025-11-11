package main

import "net/http"
import "fmt"
import "github.com/go-chi/chi/v5"
import "github.com/go-chi/chi/v5/middleware"
import "github.com/MarceloPetrucio/go-scalar-api-reference"

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/fevunge", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello fevunge!")
	})
	r.Get("/reference", func(w http.ResponseWriter, r *http.Request) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json", 
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Simple API",
			},
			DarkMode: true,
		})

		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Fprintln(w, htmlContent)
	})
	http.ListenAndServe(":8080", r)
}