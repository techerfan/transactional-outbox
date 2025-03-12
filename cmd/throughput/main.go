package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	url := "http://localhost:8081/api/orders" // Replace with your server URL

	data := map[string]interface{}{
		"product": "sample", // Replace with your request payload
		"amount":  10,
		"price":   20000,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	ticker := time.NewTicker(50 * time.Millisecond)

	for range ticker.C {
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Fatalf("Failed to create request: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatalf("Failed to read response: %v", err)
		// }

		// fmt.Printf("Response status: %d\n", resp.StatusCode)
		// fmt.Printf("Response body: %s\n", string(body))
	}
}
