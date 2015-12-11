package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DigitalInnovation/vitamns/entities"
	"github.com/DigitalInnovation/vitamns/global"
)

func GetItemsHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("GetMealsHandler called")

	err := CheckAPIKey(r)
	if err != nil {
		log.Printf("API key error: %v", err)
		http.Error(rw, "Authentication Failure", http.StatusUnauthorized)
		return
	}

	var itemsRetrieved []entities.Item
	itemsRetrieved, err = global.Dal.GetAllItems()
	log.Println("GetAllMessages")

	if err != nil {
		log.Printf("Error, failed to get message Data %v\n", err)
		http.Error(rw, "Error getting message data", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	err = json.NewEncoder(rw).Encode(itemsRetrieved)
	if err != nil {
		log.Printf("Error, failed to encode JSON %v\n", err)
		http.Error(rw, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
