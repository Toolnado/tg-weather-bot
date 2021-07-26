package main

import (
	"log"
	"os"

	"github.com/Toolnado/tg-weather-bot/internal/telegram"
	"github.com/Toolnado/tg-weather-bot/internal/weather-api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Print(err)
	}

	token, ok := os.LookupEnv("TOKEN")
	if !ok {
		log.Println("token not found")
	}

	apiKey, ok := os.LookupEnv("OPENWEATHERMAPAPIKEY")

	if !ok {
		log.Println("apikey not found")
	}

	weatherService := weather.NewWeatherService(apiKey)
	weatherBot := telegram.NewBot(weatherService, token)

	weatherBot.Start()
}
