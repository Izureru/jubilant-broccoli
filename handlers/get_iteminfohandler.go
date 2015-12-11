package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DigitalInnovation/vitamns/global"
)

func GetItemInfoHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Called GetItemInfoHandler called")

	err := CheckAPIKey(r)
	if err != nil {
		log.Printf("API key error: %v", err)
		http.Error(rw, "Authentication Failure", http.StatusUnauthorized)
		return
	}

	upc := r.URL.Query().Get(":upc")

	ItemRetrieved, err := global.Dal.GetItemInfo(upc)
	if err != nil {
		log.Printf("Error, failed to get Store Data %v\n", err)
		http.Error(rw, "Error, Failed to get Store Data", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	err = json.NewEncoder(rw).Encode(ItemRetrieved)
	if err != nil {
		log.Printf("Error, failed to get Store Data %v\n", err)
		http.Error(rw, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
