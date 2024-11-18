package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"time"
)

func SendData(ip string, port string, apiKey string, cpu interface{}, memory interface{}, disk interface{}) error {
	payload := Payload{
		CPU:    cpu,
		Memory: memory,
		Disk:   disk,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://%s:%s/receive-data", ip, port)

	for i := 0; i < 5; i++ {
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
		if err != nil {
			log.Printf("Attempt %d: Error creating request: %v", i+1, err)
			return err
		}

		req.Header.Set("X-API-Key", apiKey)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("Attempt %d: Error sending data to API: %v", i+1, err)
		} else {
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				return nil
			}
			log.Printf("Attempt %d: Server returned status %s", i+1, resp.Status)
		}

		time.Sleep(5 * time.Second)
	}

	return fmt.Errorf("failed to send data after multiple attempts")
}

