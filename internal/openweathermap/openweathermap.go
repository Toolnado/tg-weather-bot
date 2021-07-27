package openweathermap

import (
	"encoding/json"
	"net/http"

	"github.com/Toolnado/tg-weather-bot/model"
)

type OpenWeatherMapService struct {
	apiKey string
}

func NewOpenWeatherMapService(apiKey string) *OpenWeatherMapService {
	return &OpenWeatherMapService{
		apiKey: apiKey,
	}
}

func (o *OpenWeatherMapService) GetWeather(city string) (weatherData model.WeatherData, err error) {
	var data model.WeatherData
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&APPID=" + o.apiKey)

	if err != nil {
		return data, err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}

func TransformTemp(kelvin float64) int {

	kelvinConst := 273.15
	temp := kelvin - kelvinConst

	return int(temp)
}
