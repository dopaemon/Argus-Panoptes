package api

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/BurntSushi/toml"
	_ "strings"
)

type Config struct {
	Security struct {
		APIKey string `toml:"api_key"`
	} `toml:"security"`
}

func getConfig() (*Config, error) {
	var config Config
	_, err := toml.DecodeFile("./configs/server.toml", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func validateAPIKey(r *http.Request) bool {
	config, err := getConfig()
	if err != nil {
		log.Println("Error reading config:", err)
		return false
	}

	apiKey := r.Header.Get("X-API-Key")
	return apiKey == config.Security.APIKey
}

func ReceiveDataHandler(w http.ResponseWriter, r *http.Request) {
	if !validateAPIKey(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var payload Payload
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Received CPU Info: %+v\n", payload.CPU)
	log.Printf("Received Memory Info: %+v\n", payload.Memory)
	log.Printf("Received Disk Info: %+v\n", payload.Disk)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data received successfully"))
}

