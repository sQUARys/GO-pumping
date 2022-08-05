package main

import (
	"RostPart4/controller"
	"RostPart4/database"
	"time"
)

func main() {
	for true {
		db := database.New()

		bodyJSON := controller.GetBodyRequest()
		body := controller.UnMarshal(bodyJSON)

		for i := 0; i < len(body.Content); i++ {
			db.Add(body.Content[i])
		}
		time.Sleep(time.Second)
	}

}
