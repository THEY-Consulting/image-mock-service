package api

import (
	"bytes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/they-consulting/gogogo/console"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"strconv"
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
		_, err := writeImage(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("I am okay."))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})
}

func writeImage(w http.ResponseWriter) (int, error) {
	file, err := os.Open("./data/image.jpg")
	if err != nil {
		http.Error(w, err. Error(), http.StatusUnprocessableEntity)
	}
	defer file.Close()
	image, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, err. Error(), http.StatusUnprocessableEntity)
	}
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, image, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	return w.Write(buffer.Bytes())
}
