package main

import (
	"log"
	"net/http"
	"encoding/json"

	services "tide/api/services"
	controllers "tide/api/controllers"
)

func main() {
	d, err := services.NewDatabase()
	if err != nil {
		log.Fatal("Could not create database: ", err)
	}
	err = d.Migrate()
	if err != nil {
		log.Fatal("Could not migrate database: ", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// r.Body
			var newItem NewItem
			err := json.NewDecoder(r.Body).Decode(newItem)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			item := controllers.NewItem()
			json.NewEncoder(w).Encode(item)
			w.WriteHeader(http.StatusCreated)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":7890", nil)
}
