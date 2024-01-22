package external_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Agify(name string) int {
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error making HTTP request:", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}
	var result AgifyObject
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}
	return result.Age
}

type AgifyObject struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}
