package main

import (
	provider "Microservices/providers"
	"Microservices/repositories"
	"time"
)

func main() {
	var body provider.Content

	ticker := time.NewTicker(time.Minute)

	db := database.New()
	for range ticker.C {
		bodyJSON := provider.GetBodyRequest()
		body = provider.UnMarshal(bodyJSON)
		for j := 0; j < len(body.Content); j++ {
			db.Add(body.Content[j])
		}
	}
}
