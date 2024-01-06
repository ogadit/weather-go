package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	utils "github.com/ogadit/weather/utils"
)

func main() {
	godotenv.Load(".env")
	apiKey := os.Getenv("API_KEY")
	parsedData := utils.GetWeather(apiKey)

	fmt.Printf(
		"Weather for %s\n"+
			"Current Temperature:\t%.1f°C\n"+
			"Feels Like:\t\t%.1f°C\n",
		parsedData.Location.Name+", "+parsedData.Location.Country,
		parsedData.Current.Temp,
		parsedData.Current.FeelsLike)
}
