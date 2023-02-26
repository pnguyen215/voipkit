package telegram

import (
	"encoding/json"
	"errors"
	"log"

	telegramBot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var preTelegram *telegramBot.BotAPI

func NewBotWith(token string, chatId int64) (*telegramBot.BotAPI, error) {
	t := &TelegramBot{
		Token:     token,
		ChatId:    chatId,
		DebugMode: true,
		Enabled:   true,
	}
	return t.NewBot()
}

func (t *TelegramBot) NewBot() (*telegramBot.BotAPI, error) {

	if !t.Enabled {
		return &telegramBot.BotAPI{}, errors.New("Telegram Bot unavailable")
	}

	if preTelegram != nil {
		return preTelegram, nil
	}

	bot, err := telegramBot.NewBotAPI(t.Token)
	bot.Debug = t.DebugMode
	preTelegram = bot
	return bot, err
}

func (t *TelegramBot) ResetBot() {
	preTelegram = nil
}

func (t *TelegramBot) ToJson(data interface{}) string {
	s, ok := data.(string)

	if ok {
		return s
	}

	result, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err.Error())
		return ""
	}

	return string(result)
}

func (t *TelegramBot) SendMessage(message interface{}) (telegramBot.Message, error) {
	if t.ChatId <= 0 {
		return telegramBot.Message{}, errors.New("Chat Id must be provided")
	}

	bot, err := t.NewBot()
	if err != nil {
		return telegramBot.Message{}, err
	}

	content := t.ToJson(message)
	_message := telegramBot.NewMessage(t.ChatId, content)
	_message.ParseMode = telegramBot.ModeHTML
	response, err := bot.Send(_message)
	return response, err
}
