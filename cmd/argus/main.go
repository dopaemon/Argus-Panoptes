package main

import (
	_ "encoding/json"
	"log"
	"net/http"
	"argus/internal/api"
)

func main() {
	http.HandleFunc("/receive-data", api.ReceiveDataHandler)

	log.Println("Server started at http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

