package weather

import (
	"github.com/Toolnado/tg-weather-bot/model"
)

type WeatherApi interface {
	GetWeather(city string) (weatherData model.WeatherData, err error)
}

type WeatherService struct {
	Weather WeatherApi
}

func NewWeatherService(weather WeatherApi) *WeatherService {
	return &WeatherService{
		Weather: weather,
	}
}

func (o *WeatherService) GetWeather(city string) (weatherData model.WeatherData, err error) {
	data, err := o.Weather.GetWeather(city)

	if err != nil {
		return data, err
	}

	return data, nil
}
