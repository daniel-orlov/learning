package main

import (
	"fmt"
	es "github.com/pkg/errors"
	"gopkg.in/telegram-bot-api.v4"
	"net/http"
)

// func (wr *FullWeatherReport) formatForecast(period, step int) string {
// 	formatted := fmt.Sprintf("%v-h forecast for %v:\n ", period, wr.CityName)
// 	for i := 0; i < period; i += step {
// 		 formatted += fmt.Sprintf(
// 		 	" %c%c %v m/sec\n", emojis["Wind"], wr.formatWindDirection(i),
// 		 math.Round(wr.Data[i].WindSpeedMs),
// 		 )

// 	}
// }

func main() {
	//parsing config from .env
	var cfg = parseConfig()
	//creating a new BotAPI instance using token
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		err = es.Wrap(err, "failed to create new BotAPI with given token:")
		fmt.Println(err)
	}

	// bot.Debug = true

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)
	//setting a webhook
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(cfg.WebhookURL))
	if err != nil {
		err = es.Wrap(err, "failed to set WebHook:")
		fmt.Println(err)
	}

	//registers an http handler for a webhook, returns UpdatesChannel (<-chan Update)
	updates := bot.ListenForWebhook("/") //why exactly just root?
	//launching a server on a local host :8080
	go http.ListenAndServe(cfg.Port, nil)
	fmt.Println("start listen", cfg.Port)

	mainMenu := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonLocation(`Weather at my location`),
		tgbotapi.NewKeyboardButton(`A different place`),
		tgbotapi.NewKeyboardButton(`/stop`),
	)

	forecastMenu := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(`Now`),
		tgbotapi.NewKeyboardButton(`2 days`),
		tgbotapi.NewKeyboardButton(`5 days`),
		tgbotapi.NewKeyboardButton(`10 days`),
		tgbotapi.NewKeyboardButton(`< Back`),
	)

	var loc *tgbotapi.Location

	//handling messages from user
	for update := range updates {
		chatID := update.Message.Chat.ID
		userMsg := update.Message
		text := userMsg.Text
		location := userMsg.Location

		switch {
		case location != nil:
			{
				loc = userMsg.Location
				msg := tgbotapi.NewMessage(chatID, commentsEn["CoordsAccepted"])
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(forecastMenu)
				bot.Send(msg)
			}
		}
		switch text {
		case "/start":
			{
				pick := commentsEn["DefaultMessage"] + "\n" + pickASaying(sayingsEn)
				msg := tgbotapi.NewMessage(chatID, pick)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(mainMenu)
				bot.Send(msg)
			}
		case "< Back":
			{
				msg := tgbotapi.NewMessage(chatID, commentsEn["BackToMainMenu"])
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(mainMenu)
				bot.Send(msg)
			}
		case "/stop":
			{
				msg := tgbotapi.NewMessage(chatID, `Bye!`)
				msg.ReplyMarkup = tgbotapi.ReplyKeyboardHide{HideKeyboard: true}
				bot.Send(msg)
			}
		case "Now":
			{
				forecast, err := getForecast(loc, text, cfg)
				if err != nil {
					err = es.Wrap(err, "failed to getForecast:")
					fmt.Println(err)
				}
				wr, err := parseWeather(forecast)
				if err != nil {
					err = es.Wrap(err, "failed to parseWeather:")
					fmt.Println(err)
				}
				repr := wr.formatNow()
				msg := tgbotapi.NewMessage(chatID, repr)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(forecastMenu)
				bot.Send(msg)
			}
		case "2 days":
			{
				forecast, err := getForecast(loc, text, cfg)
				if err != nil {
					err = es.Wrap(err, "failed to getForecast:")
					fmt.Println(err)
				}
				wr, err := parseWeather(forecast)
				if err != nil {
					err = es.Wrap(err, "failed to parseWeather:")
					fmt.Println(err)
				}
				repr := wr.formatHours(48)
				msg := tgbotapi.NewMessage(chatID, repr)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(forecastMenu)
				bot.Send(msg)
			}
		case "5 days":
			{
				forecast, err := getForecast(loc, text, cfg)
				if err != nil {
					err = es.Wrap(err, "failed to getForecast:")
					fmt.Println(err)
				}
				wr, err := parseWeather(forecast)
				if err != nil {
					err = es.Wrap(err, "failed to parseWeather:")
					fmt.Println(err)
				}
				repr := wr.formatHours(120)
				msg := tgbotapi.NewMessage(chatID, repr)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(forecastMenu)
				bot.Send(msg)
			}
		case "10 days":
			{
				forecast, err := getForecast(loc, text, cfg)
				if err != nil {
					err = es.Wrap(err, "failed to getForecast:")
					fmt.Println(err)
				}
				wr, err := parseWeather(forecast)
				if err != nil {
					err = es.Wrap(err, "failed to parseWeather:")
					fmt.Println(err)
				}
				repr := wr.formatDays(10)
				msg := tgbotapi.NewMessage(chatID, repr)
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(forecastMenu)
				bot.Send(msg)
			}
			//case "10 days":
			//		{
			//			text := "Under development"
			//			msg := tgbotapi.NewMessage(chatID, text)
			//			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(forecastMenu)
			//			bot.Send(msg)
			//		}
		}
	}
}

//func handleNow() tgbotapi.MessageConfig {
//	forecast, err := getForecast(loc, text, cfg)
//	if err != nil {
//		err = es.Wrap(err, "failed to getForecast:")
//		fmt.Println(err)
//	}
//	wr, err := parseWeather(forecast)
//	if err != nil {
//		err = es.Wrap(err, "failed to parseWeather:")
//		fmt.Println(err)
//	}
//	repr := wr.formatNow()
//	msg := tgbotapi.NewMessage(chatID, repr)
//	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(forecastMenu)
//	return msg
//}
