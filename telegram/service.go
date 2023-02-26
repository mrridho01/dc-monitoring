package telegram

import (
	"fmt"
	"log"
	"time"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	BotToken string
}

type GH struct {
	Name                                       string
	AcInput, Dc110, Dc48, CurrentAC, CurrentDC float32
}

type Broadcast struct {
	TimePoll string
	GH       *GH
}

var Values = make(map[string]*Broadcast)

const StartMessage = "DC MONITORING PT. PLN (Persero) UP2D Kalimantan Barat\n\n Pilih lokasi GH"

func AnalogModbusText(data string) string {
	return fmt.Sprintf("GH %s\n=========\nPolling time : %s\n\nAC Input : %gVAC\nDC 110 : %gVDC\nDC 48 : %gVDC\nArus AC : %gA\nArus DC : %gA", Values[data].GH.Name, Values[data].TimePoll, Values[data].GH.AcInput, Values[data].GH.Dc110, Values[data].GH.Dc48, Values[data].GH.CurrentAC, Values[data].GH.CurrentDC)
}

func GenerateInlineKeyboard() tgbot.InlineKeyboardMarkup {
	var rows [][]tgbot.InlineKeyboardButton
	var buttons []tgbot.InlineKeyboardButton

	for _, v := range Values {
		buttons = append(buttons, tgbot.NewInlineKeyboardButtonData(v.GH.Name, v.GH.Name))

		if len(buttons)%5 == 0 {
			rows = append(rows, tgbot.NewInlineKeyboardRow(buttons...))
		}
	}

	return tgbot.NewInlineKeyboardMarkup(rows...)
}

func startTelegram() {
	var (
		UpdateId int
	)

	bot := &TelegramBot{
		BotToken: "5784852963:AAFyL566vdaCyuBrnqV7sb8VFNgecJWK_FU",
	}

	botAPI, err := tgbot.NewBotAPI(bot.BotToken)
	if err != nil {
		log.Panic(err)
	}

	botAPI.Debug = true
	log.Printf("Service start on bot : %s", botAPI.Self.UserName)

	updateInstance := tgbot.NewUpdate(UpdateId)
	updateInstance.Timeout = 60

	updates := botAPI.GetUpdatesChan(updateInstance)

	for update := range updates {
		if update.Message != nil {
			if !update.Message.IsCommand() {
				continue
			}

			chatId := update.Message.Chat.ID
			sendStartMessage := tgbot.NewMessage(chatId, "")

			if update.Message.Command() == "start" {
				sendStartMessage.Text = StartMessage
				sendStartMessage.ReplyMarkup = GenerateInlineKeyboard()

				_, err := botAPI.Send(sendStartMessage)
				if err != nil {
					log.Panic(err)
				}
			}
		} else if update.CallbackQuery != nil {
			callback := tgbot.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)

			if _, err := botAPI.Request(callback); err != nil {
				log.Panic(err)
			}

			chatId := update.CallbackQuery.Message.Chat.ID
			messageId := update.CallbackQuery.Message.MessageID

			data := update.CallbackQuery.Data
			value := Values[data]

			currentTime := time.Now().Format("2006/01/02 15:04:05")
			value.TimePoll = currentTime

			msg := tgbot.NewEditMessageText(chatId, messageId, AnalogModbusText(data))
			inlineKeyboard := GenerateInlineKeyboard()
			msg.ReplyMarkup = &inlineKeyboard

			_, err := botAPI.Send(msg)
			if err != nil {
				log.Panic(err)
			}
		}
	}
}
