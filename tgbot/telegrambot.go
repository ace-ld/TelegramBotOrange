package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	getInfoCovid "ace-h/tgbot/api/covidSummaryAPI"
	getJoke "ace-h/tgbot/api/jokesAPI"
	get "ace-h/tgbot/db"
)

var (
	buttonTD = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "Total Death COVID-19"}}
	buttonTC = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "Total Confirmed COVID-19"}}
	buttonNC = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "New confirmed COVID-19"}}
	buttonND = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "New Deaths COVID-19"}}
	buttonNR = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "New recovered COVID-19"}}
	buttonTR = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "Total recovered COVID-19"}}
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func main() {

	bot, err := tgbotapi.NewBotAPI("813814117:AAEy5T8hws-wU86USfOQOcHQ_kOZFu8-x68")
	check(err)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		var (
			message   tgbotapi.MessageConfig
			msgStiker tgbotapi.StickerConfig
			msgPhoto  tgbotapi.PhotoConfig

			chatID = update.Message.Chat.ID
			text   = update.Message.Text
		)

		log.Println("received text: ", update.Message.Text) // логируем сообщения

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		switch text {

		// DATABASE
		case "Hello", "hello", "Привет", "привет":
			message = tgbotapi.NewMessage(chatID, get.HelloWords())
		case "Иди нахуй", "иди нахуй", "Иди на хуй", "иди на хуй", "Пидарас", "пидарас":
			message = tgbotapi.NewMessage(chatID, get.DirtyWords())
		case "Расскажи о себе", "расскажи о себе":
			message = tgbotapi.NewMessage(chatID, get.AboutBot())
		case "Расскажи что-нибудь", "расскажи что-нибудь", "Расскажи что нибудь", "расскажи что нибудь", "Расскажи чтонибудь", "расскажи чтонибудь":
			message = tgbotapi.NewMessage(chatID, get.HistoryWords())
		// DATABASE

		// OTHER
		case "Бусинка", "бусинка", "Буся", "буся":
			f := "C:/Users/gl_ni/go/src/ace-h/tgbot/pics/busya.jpg"
			message = tgbotapi.NewMessage(chatID, `ЛЯ КАКАЯ`)
			msgPhoto = tgbotapi.NewPhotoUpload(chatID, f)
		// OTHER

		// REST API
		case "Хочу шутку", "хочу шутку":
			message = tgbotapi.NewMessage(chatID, getJoke.GetJoke())
		case "Total Death COVID-19":
			message = tgbotapi.NewMessage(chatID, "Всего смертей в мире: "+getInfoCovid.TotalDeath())
		case "Total Confirmed COVID-19":
			message = tgbotapi.NewMessage(chatID, "Всего подтвержденных заражений в мире: "+getInfoCovid.TotalConfirmed())
		case "New confirmed COVID-19":
			message = tgbotapi.NewMessage(chatID, "Новых заражений в мире: "+getInfoCovid.NewConfirmed())
		case "New Deaths COVID-19":
			message = tgbotapi.NewMessage(chatID, "Новых смертей в мире: "+getInfoCovid.NewDeaths())
		case "New recovered COVID-19":
			message = tgbotapi.NewMessage(chatID, "Вылечилось за сутки в мире: "+getInfoCovid.NewRecovered())
		case "Total recovered COVID-19":
			message = tgbotapi.NewMessage(chatID, "Всего вылечилось в мире: "+getInfoCovid.TotalRecovered())
		// REST API

		default:
			message = tgbotapi.NewMessage(chatID, `Я не понимаю что ты хочешь :(`+"\n"+`Напиши @boot_fail`)
			msgStiker = tgbotapi.NewStickerShare(chatID, "CAACAgIAAxkBAAEBXvNfbiYSR9q_YngPLLFUzPOwNHn6DwAC6gUAAtCG-wpnYVb-QhDZaBsE")
		}

		//message.ReplyToMessageID = update.Message.MessageID
		message.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttonTD, buttonTC, buttonTR, buttonND, buttonNC, buttonNR)

		bot.Send(message)
		bot.Send(msgStiker)
		bot.Send(msgPhoto)
	}
}
