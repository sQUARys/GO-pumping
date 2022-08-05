package controller

import (
	"RostPart4/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetBodyRequest() []byte {
	url := "http://localhost:8081"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return body
}

func UnMarshal(body []byte) models.Content {

	var content models.Content

	err := json.Unmarshal(body, &content)

	if err != nil {
		fmt.Println(err)
		return models.Content{}
	}
	return content
}
