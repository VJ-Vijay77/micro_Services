package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// setting a reciever from main file struct i.e app * config
func(app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	//  specifying who is allowed to connect using a middleware in chi
	//  CSRF is Cross-Site request forgery ,for protection
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"Accept","Authorization","Content-Type","X-CSRF-Token"},
		AllowCredentials: true,
		ExposedHeaders: []string{"Link"},
		MaxAge: 300,
	}))

	//To checking the service is still alive
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/",app.Broker)

	return mux

}