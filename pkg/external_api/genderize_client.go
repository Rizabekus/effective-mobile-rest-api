package external_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Genderize(name string) string {
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error making HTTP request:", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}
	var result GenderizeObject
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}
	return result.Gender
}

type GenderizeObject struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}
