package providers

import (
	dbRepo "Microservices/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Content struct {
	Content []dbRepo.Order `json:"content"`
}

func GetBodyRequest() []byte {
	url := "http://localhost:8081"
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error: ", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error: ", err)
		return nil
	}

	return body
}

func UnMarshal(body []byte) Content {
	var content Content
	err := json.Unmarshal(body, &content)

	if err != nil {
		log.Println("Error: ", err)
		return Content{}
	}
	return content
}
