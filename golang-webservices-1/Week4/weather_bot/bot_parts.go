package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	es "github.com/pkg/errors"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"strconv"
)

var locationKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonLocation(commentsEn["AtMyLocation"]),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(commentsEn["AtADiffPlace"]),
	),
)

var daysOrHoursKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(commentsEn["ByHours"]),
		tgbotapi.NewKeyboardButton(commentsEn["ByDays"]),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(commentsEn["Now"]),
		tgbotapi.NewKeyboardButton(commentsEn["Back0"]),
	),
)

var daysKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(commentsEn["3Days"]),
		tgbotapi.NewKeyboardButton(commentsEn["5Days"]),
		tgbotapi.NewKeyboardButton(commentsEn["7Days"]),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(commentsEn["10Days"]),
		tgbotapi.NewKeyboardButton(commentsEn["16Days"]),
		tgbotapi.NewKeyboardButton(commentsEn["Back1"]),
	),
)

var hoursKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(commentsEn["24Hours"]),
		tgbotapi.NewKeyboardButton(commentsEn["48Hours"]),
		tgbotapi.NewKeyboardButton(commentsEn["72Hours"]),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(commentsEn["96Hours"]),
		tgbotapi.NewKeyboardButton(commentsEn["120Hours"]),
		tgbotapi.NewKeyboardButton(commentsEn["Back1"]),
	),
)

var keyboards = map[string]tgbotapi.ReplyKeyboardMarkup{
	"main":   locationKeyboard,
	"period": daysOrHoursKeyboard,
	"days":   daysKeyboard,
	"hours":  hoursKeyboard,
}

//HANDLERS, FUTURE MAP[STRING]FUNC
func handleStop(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	fmt.Println("EXECUTING: handleStop")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, commentsEn["Stop"])
	msg.ReplyMarkup = tgbotapi.ReplyKeyboardHide{HideKeyboard: true}
	bot.Send(msg)
	bot.StopReceivingUpdates() //should turn the bot off
}
func handleStart(bot *tgbotapi.BotAPI, conn *pgx.Conn, update tgbotapi.Update) {
	fmt.Println("EXECUTING: handleStart")
	go addUserIfNotExists(conn, update)
	greeting := commentsEn["DefaultMessage"] + "\n" + pickASaying(sayingsEn)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, greeting)
	msg.ReplyMarkup = keyboards["main"]
	bot.Send(msg)
}
func handleBackToMainMenu(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	fmt.Println("EXECUTING: handleBackToMainMenu")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, commentsEn["ChooseLocation"])
	msg.ReplyMarkup = keyboards["main"]
	bot.Send(msg)
}
func handleBack(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	fmt.Println("EXECUTING: handleBack")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, commentsEn["ChoosePeriodType"])
	msg.ReplyMarkup = keyboards["period"]
	bot.Send(msg)
}
func handleByDays(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	fmt.Println("EXECUTING: handleByDays")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, commentsEn["ChoosePeriod"])
	msg.ReplyMarkup = keyboards["days"]
	bot.Send(msg)
}
func handleByHours(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	fmt.Println("EXECUTING: handleByHours")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, commentsEn["ChoosePeriod"])
	msg.ReplyMarkup = keyboards["hours"]
	bot.Send(msg)
}
func handleNow(bot *tgbotapi.BotAPI, conn *pgx.Conn, update tgbotapi.Update) {
	fmt.Println("EXECUTING: handleNow")
	text := update.Message.Text
	userId := update.Message.From.ID
	chatId := update.Message.Chat.ID

	loc, _ := getMostRecentLocation(conn, userId)
	forecast, err := getForecast(&loc, text)
	if err != nil {
		err = es.Wrap(err, "failed to getForecast")
		fmt.Println(err)
	}

	wr, err := parseWeather(forecast)
	if err != nil {
		err = es.Wrap(err, "failed to parseWeather")
		fmt.Println(err)
	}

	nameMostRecentLocation(wr.Data[0].CityName, conn, userId)
	repr := wr.formatNow()

	msg := tgbotapi.NewMessage(chatId, repr)
	msg.ReplyMarkup = keyboards["period"]
	bot.Send(msg)
}
func handleHours(bot *tgbotapi.BotAPI, conn *pgx.Conn, update tgbotapi.Update) {
	fmt.Println("EXECUTING: handleHours")
	text := update.Message.Text
	timePeriod := timePeriodsEn[text]
	userId := update.Message.From.ID
	chatId := update.Message.Chat.ID

	loc, _ := getMostRecentLocation(conn, userId)
	forecast, err := getForecast(&loc, text)
	if err != nil {
		err = es.Wrap(err, "failed to getForecast")
		fmt.Println(err)
	}

	wr, err := parseWeather(forecast)
	if err != nil {
		err = es.Wrap(err, "failed to parseWeather")
		fmt.Println(err)
	}

	nameMostRecentLocation(wr.CityName, conn, userId)
	repr := wr.formatHours(timePeriod)

	msg := tgbotapi.NewMessage(chatId, repr)
	msg.ReplyMarkup = keyboards["hours"]
	bot.Send(msg)
}
func handleDays(bot *tgbotapi.BotAPI, conn *pgx.Conn, update tgbotapi.Update) {
	fmt.Println("EXECUTING: handleDays")
	text := update.Message.Text
	timePeriod := timePeriodsEn[text]
	userId := update.Message.From.ID
	chatId := update.Message.Chat.ID

	loc, _ := getMostRecentLocation(conn, userId)
	forecast, err := getForecast(&loc, text)
	if err != nil {
		err = es.Wrap(err, "failed to getForecast")
		fmt.Println(err)
	}

	wr, err := parseWeather(forecast)
	if err != nil {
		err = es.Wrap(err, "failed to parseWeather")
		fmt.Println(err)
	}

	nameMostRecentLocation(wr.CityName, conn, userId)
	repr := wr.formatDays(timePeriod)

	msg := tgbotapi.NewMessage(chatId, repr)
	msg.ReplyMarkup = keyboards["days"]
	bot.Send(msg)
}
func handleLocationByCoords(bot *tgbotapi.BotAPI, conn *pgx.Conn, update tgbotapi.Update) {
	fmt.Println("EXECUTING: handleLocationByCoords")
	addLocationByCoords(conn, update.Message)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, commentsEn["CoordsAccepted"])
	msg.ReplyMarkup = keyboards["period"]
	bot.Send(msg)
}
func handleLocationByText(bot *tgbotapi.BotAPI, conn *pgx.Conn, update tgbotapi.Update) {
	fmt.Println("EXECUTING: handleLocationByText")
	addLocationByCoords(conn, update.Message)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, commentsEn["CoordsAccepted"])
	msg.ReplyMarkup = keyboards["period"]
	bot.Send(msg)
}
func handleUnknown(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	fmt.Println("EXECUTING: handleUnknown")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, commentsEn["Unknown"])
	msg.ReplyMarkup = keyboards["main"]
	bot.Send(msg)
}

//DATABASE FUNCTIONS
func addUserIfNotExists(conn *pgx.Conn, update tgbotapi.Update) {
	/* TESTED, IT WORKS AS EXPECTED
	 */
	fmt.Println("EXECUTING: addUserIfNotExists")
	username := update.Message.From.UserName
	userId := fmt.Sprintf("%v", update.Message.From.ID)
	language := update.Message.From.LanguageCode
	sqlQuery := fmt.Sprintf(
		`INSERT INTO users (user_id, username, language)
				VALUES (%v, '%v', '%v')
				ON CONFLICT (user_id) DO NOTHING
				`,
		userId, username, language,
	)
	err := conn.QueryRow(context.Background(), sqlQuery).Scan()
	if err.Error() == "no rows in result set" {
		fmt.Println("SUCCEEDED: QueryRow when adding UserIfNotExists")
	} else if err != nil {
		err = es.Wrap(err, "FAILED: QueryRow when adding UserIfNotExists")
		fmt.Println(err)
	}
}
func addLocationByCoords(conn *pgx.Conn, msg *tgbotapi.Message) {
	/* TESTED, IT WORKS AS EXPECTED
	 */
	fmt.Println("EXECUTING: addLocationByCoords")
	userId := msg.From.ID
	lat := fmt.Sprintf("%v", msg.Location.Latitude)
	long := fmt.Sprintf("%v", msg.Location.Longitude)

	sqlQuery := fmt.Sprintf(
		`INSERT INTO locations ("user", "latitude", "longitude")
				VALUES (%v, %v, %v)
				`,
		userId, lat, long,
	)
	err := conn.QueryRow(context.Background(), sqlQuery).Scan()
	if err.Error() == "no rows in result set" {
		fmt.Println("SUCCEEDED: QueryRow when adding LocationByCoords")
	} else if err != nil {
		err = es.Wrap(err, "FAILED: QueryRow when adding LocationByCoords")
		fmt.Println(err)
	}
}
func nameMostRecentLocation(name string, conn *pgx.Conn, userId int) {
	fmt.Println("EXECUTING: nameMostRecentLocation")
	loc, id := getMostRecentLocation(conn, userId)
	sqlQuery := fmt.Sprintf(
		`INSERT INTO locations (id, "user", latitude, longitude, location)
				VALUES (%v, %v, '%v', '%v', '%v')
				ON CONFLICT (id) DO UPDATE
				SET location = '%v'
				`,
		id, userId, loc.Latitude, loc.Longitude, name, name,
	)
	fmt.Println(sqlQuery)
	err := conn.QueryRow(context.Background(), sqlQuery).Scan()
	if err.Error() == "no rows in result set" {
		fmt.Println("SUCCEEDED: QueryRow when naming MostRecentLocation")
	} else if err != nil {
		err = es.Wrap(err, "FAILED: QueryRow when naming MostRecentLocation")
		fmt.Println(err)
	}
}
func getMostRecentLocation(conn *pgx.Conn, userId int) (tgbotapi.Location, int) {
	/* TESTED, IT WORKS AS EXPECTED
	 */
	fmt.Println("EXECUTING: getMostRecentLocation")
	loc := tgbotapi.Location{}
	sqlQuery := fmt.Sprintf(
		`SELECT id, latitude, longitude
				FROM locations
				WHERE "user" = '%v'
				ORDER BY id DESC
				LIMIT 1
				`,
		userId,
	)
	fmt.Println(sqlQuery)
	var lat, long string
	var id int
	err := conn.QueryRow(context.Background(), sqlQuery).Scan(&id, &lat, &long)
	if err != nil {
		err = es.Wrap(err, "FAILED: QueryRow when adding LocationByCoords")
		fmt.Println(err)
	}
	fmt.Println("LONG", long, "LAT", lat)
	loc.Latitude, err = strconv.ParseFloat(lat, 64)
	if err != nil {
		err = es.Wrap(err, "Latitude parsing failed")
		fmt.Println(err)
	}
	loc.Longitude, err = strconv.ParseFloat(long, 64)
	if err != nil {
		err = es.Wrap(err, "Longitude parsing failed")
		fmt.Println(err)
	}
	return loc, id
}
