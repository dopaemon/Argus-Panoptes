package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"log"
)

func SendData(ip string, port string, cpu interface{}, memory interface{}, disk interface{}) error {
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
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
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

