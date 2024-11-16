package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReceiveDataHandler(w http.ResponseWriter, r *http.Request) {
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

