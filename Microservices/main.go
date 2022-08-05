package main

import (
	"Microservices/controller"
	"Microservices/database"
	"Microservices/models"
	"sync"
	"time"
)

func main() {
	var (
		body models.Content
		wg   sync.WaitGroup
	)

	ticker := time.NewTicker(time.Minute)

	wg.Add(1)
	go func() {
		defer wg.Done()

		db := database.New()
		for {
			select {
			case <-ticker.C:
				bodyJSON := controller.GetBodyRequest()
				body = controller.UnMarshal(bodyJSON)
				for j := 0; j < len(body.Content); j++ {
					db.Add(body.Content[j])
				}
			}
		}
	}()
	defer wg.Wait()
}
