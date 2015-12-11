package handlers

import (
	"errors"
	"github.com/DigitalInnovation/vitamns/global"
	"log"
	"net/http"
)

func CheckAPIKey(r *http.Request) error {

	key := r.Header.Get("mnshealthkey")
	if len(key) == 0 {
		return errors.New("No Key Provided")
	}

	if len(global.Config.APIKey) == 0 {
		return errors.New("No server side key")
	}

	if key == global.Config.APIKey {
		return nil
	}

	return errors.New("Invalid API Key")
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Home Handler Called")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
}
