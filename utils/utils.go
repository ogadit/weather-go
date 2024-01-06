package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ipResponse struct {
	City string `json:"city"`
}

type weatherData struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		Temp      float64 `json:"temp_c"`
		FeelsLike float64 `json:"feelslike_c"`
	} `json:"current"`
}

/*
# Gets the city name

Uses an http request to get the ip address and get the name
*/
func getCity() string {
	var response ipResponse

	data, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &response)

	return response.City
}

/*
# Gets the weather

Uses a `city` and `apiKey` input to get weatherData
*/
func GetWeather(apiKey string) weatherData {
	var response weatherData

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, getCity())

	data, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(body, &response)

	return response
}
