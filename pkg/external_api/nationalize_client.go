package external_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Nationalize(name string) string {

	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)

	response, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("Error making HTTP request: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Sprintf("Error reading response body: %v", err)
	}

	var result NationalizeObject
	err = json.Unmarshal(body, &result)
	if err != nil {
		return fmt.Sprintf("Error decoding JSON: %v", err)
	}

	maxProbability := 0.0
	maxCountryID := ""

	for _, country := range result.Country {
		if country.Probability > maxProbability {
			maxProbability = country.Probability
			maxCountryID = country.CountryID
		}
	}

	return maxCountryID
}

type NationalizeObject struct {
	Count   int    `json:"count"`
	Name    string `json:"name"`
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}
