package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
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
	//establishing connection to database
	conn, err := pgx.Connect(context.Background(), cfg.DbUrl)
	if err != nil {
		err = es.Wrap(err, "Unable to connect to database:")
		fmt.Println(err)
	}
	defer conn.Close(context.Background())

	//id := 1
	//var username string
	//var user_id string
	//var language string
	//sqlQuery := "SELECT user_id, username, language FROM users WHERE id=1"
	//err = conn.QueryRow(context.Background(), sqlQuery).Scan(&user_id, &username, &language)
	//if err != nil {
	//	err = es.Wrap(err, "QueryRow failed:")
	//	fmt.Println(err)
	//}

	//fmt.Println(user_id, username, language)

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

	//var loc *tgbotapi.Location

	//handling messages from user
	for update := range updates {

		addUserIfNotExists(conn, update)
		//var req Request = Request{
		//	Config:    cfg,
		//	Update:    update,
		//	Location:  update.Message.Location,
		//	Keyboards: keyboards,
		//}
		//
		//switch {
		//case location != nil:
		//	{
		//		loc = userMsg.Location
		//		msg := tgbotapi.NewMessage(chatID, commentsEn["CoordsAccepted"])
		//		msg.ReplyMarkup = keyboards["period"]
		//		bot.Send(msg)
		//	}
		//}
		//switch text {
		//case "/start":
		//	{
		//		greeting := commentsEn["DefaultMessage"] + "\n" + pickASaying(sayingsEn)
		//		msg := tgbotapi.NewMessage(chatID, greeting)
		//		msg.ReplyMarkup = keyboards["main"]
		//		bot.Send(msg)
		//	}
		//case "< Back":
		//	{
		//		msg := tgbotapi.NewMessage(chatID, commentsEn["BackToMainMenu"])
		//		msg.ReplyMarkup = keyboards["main"]
		//		bot.Send(msg)
		//	}
		//case "<< Back":
		//	{
		//		msg := tgbotapi.NewMessage(chatID, "")
		//		msg.ReplyMarkup = keyboards["period"]
		//		bot.Send(msg)
		//	}
		//case "/stop":
		//	{
		//		msg := tgbotapi.NewMessage(chatID, commentsEn["OnStop"])
		//		msg.ReplyMarkup = tgbotapi.ReplyKeyboardHide{HideKeyboard: true}
		//		bot.Send(msg)
		//		bot.StopReceivingUpdates() //should turn the bot off
		//	}
		//case "Now":
		//	{
		//
		//		handleNow(req *Request) tgbotapi.MessageConfig
		//		msg := tgbotapi.NewMessage(chatID, repr)
		//		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(forecastMenu)
		//		bot.Send(msg)
		//	}
		//case "2 days":
		//	{
		//		forecast, err := getForecast(loc, text, cfg)
		//		if err != nil {
		//			err = es.Wrap(err, "failed to getForecast:")
		//			fmt.Println(err)
		//		}
		//		wr, err := parseWeather(forecast)
		//		if err != nil {
		//			err = es.Wrap(err, "failed to parseWeather:")
		//			fmt.Println(err)
		//		}
		//		repr := wr.formatHours(48)
		//		msg := tgbotapi.NewMessage(chatID, repr)
		//		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(forecastMenu)
		//		bot.Send(msg)
		//	}
		//case "5 days":
		//	{
		//		forecast, err := getForecast(loc, text, cfg)
		//		if err != nil {
		//			err = es.Wrap(err, "failed to getForecast:")
		//			fmt.Println(err)
		//		}
		//		wr, err := parseWeather(forecast)
		//		if err != nil {
		//			err = es.Wrap(err, "failed to parseWeather:")
		//			fmt.Println(err)
		//		}
		//		repr := wr.formatHours(120)
		//		msg := tgbotapi.NewMessage(chatID, repr)
		//		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(forecastMenu)
		//		bot.Send(msg)
		//	}
		//case "10 days":
		//	{
		//		forecast, err := getForecast(loc, text, cfg)
		//		if err != nil {
		//			err = es.Wrap(err, "failed to getForecast:")
		//			fmt.Println(err)
		//		}
		//		wr, err := parseWeather(forecast)
		//		if err != nil {
		//			err = es.Wrap(err, "failed to parseWeather:")
		//			fmt.Println(err)
		//		}
		//		repr := wr.formatDays(10)
		//		msg := tgbotapi.NewMessage(chatID, repr)
		//		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(forecastMenu)
		//		bot.Send(msg)
		//	}
		//case "A different place":
		//	{
		//		text := "Under development"
		//		msg := tgbotapi.NewMessage(chatID, text)
		//		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(mainMenu)
		//		bot.Send(msg)
		//	}
		//}
	}
}

func addUserIfNotExists(conn *pgx.Conn, update tgbotapi.Update) {
	username := update.Message.From.UserName
	userId := fmt.Sprintf("%v", update.Message.From.ID)
	language := update.Message.From.LanguageCode
	sqlQuery := fmt.Sprintf(
		`INSERT INTO users (user_id, username, language)
				VALUES (%v, '%v', '%v')
				ON CONFLICT (user_id) DO NOTHING
				RETURNING (id, username)
				`,
		userId, username, language,
	)

	var id int
	var user string
	err := conn.QueryRow(context.Background(), sqlQuery).Scan(&id, &user)
	if err != nil {
		err = es.Wrap(err, "QueryRow failed:")
		fmt.Println(err)
	}
	fmt.Println(id, user)
}

//func handleNow(req *Request) tgbotapi.MessageConfig {
//	text := req.Update.Message.Text
//	forecast, err := getForecast(req.Location, text, req.Config)
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
//	chatId := req.Update.Message.Chat.ID
//	msg := tgbotapi.NewMessage(chatId, repr)
//	keyboard := req.Keyboards["period"]
//	msg.ReplyMarkup = keyboard
//	return msg
//}
//
//func handleByHours(req *Request) tgbotapi.MessageConfig {
//	text := req.Update.Message.Text
//	forecast, err := getForecast(req.Location, text, req.Config)
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
//	chatId := req.Update.Message.Chat.ID
//	msg := tgbotapi.NewMessage(chatId, repr)
//	keyboard := req.Keyboards["hours"]
//	msg.ReplyMarkup = keyboard
//	return msg
//}
