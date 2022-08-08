package main

import (
	"Microservices/database"
	"Microservices/models"
	database2 "Microservices/providers"
	"time"
)

func main() {
	var body models.Content

	ticker := time.NewTicker(time.Minute)

	db := database.New()
	for range ticker.C {
		bodyJSON := database2.GetBodyRequest()
		body = database2.UnMarshal(bodyJSON)
		for j := 0; j < len(body.Content); j++ {
			db.Add(body.Content[j])
		}
	}
}
