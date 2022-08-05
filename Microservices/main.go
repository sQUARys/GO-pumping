package main

import (
	"Microservices/controller"
	"Microservices/database"
	"Microservices/models"
	"time"
)

func main() {
	var body models.Content

	ticker := time.NewTicker(time.Minute)

	db := database.New()
	for range ticker.C {
		bodyJSON := controller.GetBodyRequest()
		body = controller.UnMarshal(bodyJSON)
		for j := 0; j < len(body.Content); j++ {
			db.Add(body.Content[j])
		}
	}
}
