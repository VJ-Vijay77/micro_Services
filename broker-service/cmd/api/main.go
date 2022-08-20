package main

import (
	"fmt"
	"log"
	"net/http"
)

const WebPort = "8080"

type Config struct{}

func main() {

	app := Config{}
	
	log.Printf("\nStarting broker service on Port %s\n", WebPort)

	//defining http server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s",WebPort),
		Handler: app.routes(), //routes from the routes.go
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
