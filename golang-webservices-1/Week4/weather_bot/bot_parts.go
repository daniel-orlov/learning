package main

import tgbotapi "gopkg.in/telegram-bot-api.v4"

type Request struct {
	Config    *config
	Update    tgbotapi.Update
	Location  *tgbotapi.Location
	Keyboards map[string]tgbotapi.ReplyKeyboardMarkup
}

//func NewRequest()

var locationKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonLocation(commentsEn["AtMyLocation"]),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(commentsEn["AtADifferentPlace"]),
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
