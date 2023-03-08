package main

import (
	"github.com/reactivejson/flaypath/.idea/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	api.RegisterHandlers(router)

	log.Println("Listening on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
