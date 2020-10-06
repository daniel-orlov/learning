package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	es "github.com/pkg/errors"
	"gopkg.in/telegram-bot-api.v4"
	"net/http"
)

func main() {
	//parsing config from .env
	var cfg = parseConfig()

	//establishing connection to database
	conn, err := pgx.Connect(context.Background(), cfg.DbUrl)
	if err != nil {
		err = es.Wrap(err, "Unable to connect to database")
		fmt.Println(err)
	}
	defer conn.Close(context.Background())

	//creating a new BotAPI instance using token
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		err = es.Wrap(err, "failed to create new BotAPI with given token")
		fmt.Println(err)
	}
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	// bot.Debug = true

	//setting a webhook
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(cfg.WebhookURL))
	if err != nil {
		err = es.Wrap(err, "failed to set WebHook")
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
		fmt.Printf("MESSAGE: %+v\n", update.Message)
		fmt.Printf("TEXT: %+v\n", update.Message.Text)

		text := update.Message.Text
		switch text {
		case "/start":
			handleStart(bot, conn, update)
		case "/stop":
			handleStop(bot, update)
		case "":
			if update.Message.Location != nil {
				handleLocationByCoords(bot, conn, update)
			} else {
				handleUnknown(bot, update)
			}
		case commentsEn["AtADiffPlace"]:
			handleLocationByText(bot, conn, update)
		case commentsEn["ByDays"]:
			handleByDays(bot, update)
		case commentsEn["ByHours"]:
			handleByHours(bot, update)
		case commentsEn["Back0"]:
			handleBackToMainMenu(bot, update)
		case commentsEn["Back1"]:
			handleBack(bot, update)
		case commentsEn["Now"]:
			handleNow(bot, conn, update)
		case commentsEn["3Days"],
			commentsEn["5Days"],
			commentsEn["7Days"],
			commentsEn["10Days"],
			commentsEn["16Days"]:
			handleDays(bot, conn, update)
		case commentsEn["24Hours"],
			commentsEn["48Hours"],
			commentsEn["72Hours"],
			commentsEn["96Hours"],
			commentsEn["120Hours"]:
			handleHours(bot, conn, update)
		default:
			handleUnknown(bot, update)
		}
	}
}
