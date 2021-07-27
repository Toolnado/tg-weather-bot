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
	bot     *tgbotapi.BotAPI
}

func NewBot(weather BotGetWeather, token string) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	return &Bot{
		Weather: weather,
		token:   token,
		bot:     bot,
	}
}

func (b *Bot) Start() {
	// bot.Debug = true
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)

	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		text := ""
		id := update.Message.Chat.ID

		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		data, err := b.CheckWeather(update.Message.Text)

		if err != nil {
			log.Fatal(err)
			continue
		}

		if update.Message.Text == "/start" {
			text = fmt.Sprintf("Здравствуйте, %s. Введите название места на английском языке.", update.Message.From.UserName)
			b.SendMessage(id, text)
			continue
		}

		if data.Main.Kelvin == 0 {
			text = "Город не найден"
			b.SendMessage(id, text)
			continue
		}

		temp := openweathermap.TransformTemp(data.Main.Kelvin)
		text = fmt.Sprintf("Cредняя температура в %s: %d ℃", data.Name, temp)

		b.SendMessage(id, text)
	}
}

func (b *Bot) CheckWeather(city string) (weather model.WeatherData, err error) {
	data, err := b.Weather.GetWeather(city)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (b *Bot) SendMessage(id int64, text string) {
	msg := tgbotapi.NewMessage(id, text)
	b.bot.Send(msg)
	log.Println("Send message")
}
