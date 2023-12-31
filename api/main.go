package main

import (
	"encoding/json"
	"log"
	"net/http"

	controllers "github.com/zoriya/tide/api/controllers"
	models "github.com/zoriya/tide/api/models"
	services "github.com/zoriya/tide/api/services"
)

func main() {
	db, err := services.NewDatabase()
	if err != nil {
		log.Fatal("Could not create database: ", err)
	}
	err = db.Migrate()
	if err != nil {
		log.Fatal("Could not migrate database: ", err)
	}

	qbittorent, err := services.NewQbittorent()
	if err != nil {
		log.Fatal("Could not connect to aria2", err)
	}
	controller := controllers.NewController(db, qbittorent)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var newItem models.NewItem
			err := json.NewDecoder(r.Body).Decode(&newItem)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			if newItem.Uri == "" {
				http.Error(w, "Uri is a required field", http.StatusBadRequest)
				return
			}

			item, err := controller.NewItem(newItem)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			json.NewEncoder(w).Encode(item)
			w.WriteHeader(http.StatusCreated)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	log.Print("Listening on :7890")
	http.ListenAndServe(":7890", nil)
}
