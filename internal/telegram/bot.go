package telegram

import (
	"fmt"
	"log"

	"github.com/Toolnado/tg-weather-bot/internal/openweathermap"
	"github.com/Toolnado/tg-weather-bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BotGetWeather interface {
	GetWeather(city string) (weatherData model.WeatherData, err error)
}

type Bot struct {
	Weather BotGetWeather
	token   string
}

func NewBot(weather BotGetWeather, token string) *Bot {
	return &Bot{
		Weather: weather,
		token:   token,
	}
}

func (b *Bot) Start() {
	bot, err := tgbotapi.NewBotAPI(b.token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Print(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		data, err := b.CheckWeather(update.Message.Text)

		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint(err))
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

		if data.Main.Kelvin == 0 {
			text := "Город не найден"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
			continue
		}
		temp := openweathermap.TransformTemp(data.Main.Kelvin)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint(temp))
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func (b *Bot) CheckWeather(city string) (weather model.WeatherData, err error) {
	data, err := b.Weather.GetWeather(city)
	if err != nil {
		return data, err
	}

	return data, nil
}
