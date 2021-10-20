package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/they-consulting/gogogo/console"
	"net/http"
)

func InitRouter(p string) {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "UserId"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	setupRoutes(r)

	console.Logger.Printf("Listening on port %s\n", p)
	console.Logger.Fatal(http.ListenAndServe(":"+p, r))
}

func setupRoutes(r chi.Router) {
	console.Logger.Println("Initializing Routes.")

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Here's your pic, Dick."))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})
}
