package weather

import (
	"encoding/json"
	"net/http"

	"github.com/Toolnado/tg-weather-bot/internal/model"
)

type WeatherService struct {
	apiKey string
}

func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{apiKey: apiKey}
}

func (w *WeatherService) GetWeather(city string) (weatherData model.WeatherData, err error) {
	var data model.WeatherData
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&APPID=" + w.apiKey)

	if err != nil {
		return data, err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}
