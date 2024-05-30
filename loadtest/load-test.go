package main

import (
	"fmt"
	"net/http"
	"time"
)

func loadTestAPI(endpoint string, duration time.Duration) {
	end := time.Now().Add(duration)
	for time.Now().Before(end) {
		resp, err := http.Get(endpoint)
		if err != nil {
			fmt.Println("Request error:", err)
		} else {
			fmt.Println("Response status:", resp.StatusCode)
			resp.Body.Close()
		}
	}
}

func main() {
	endpoint := "http://training-app-service.dev.svc.cluster.local:8080/api/health"
	loadTestAPI(endpoint, 5*time.Minute)
}
