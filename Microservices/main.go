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
		body             models.Content
		dataToInsertInDB []models.Content
		isDone           bool
		wg               sync.WaitGroup
	)

	ticker := time.NewTicker(100 * time.Millisecond)
	stopCh := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()

		db := database.New()
		for !isDone {
			select {
			case <-stopCh:
				isDone = true
			case <-ticker.C:
				bodyJSON := controller.GetBodyRequest()
				body = controller.UnMarshal(bodyJSON)
				dataToInsertInDB = append(dataToInsertInDB, body)
			}
		}

		for i := 0; i < len(dataToInsertInDB); i++ {
			for j := 0; j < len(dataToInsertInDB[i].Content); j++ {
				db.Add(dataToInsertInDB[i].Content[j])
			}
		}
	}()

	time.Sleep(601 * time.Millisecond)
	stopCh <- true
	ticker.Stop()
	defer wg.Wait()
}
