package telegram

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

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

func NewBotWithMultiChatId(token string, chatId []int64) (*telegramBot.BotAPI, error) {
	t := &TelegramBot{
		Token:       token,
		MultiChatId: chatId,
		DebugMode:   true,
		Enabled:     true,
	}
	return t.NewBot()
}

func (t *TelegramBot) NewBot() (*telegramBot.BotAPI, error) {

	if !t.Enabled {
		return &telegramBot.BotAPI{}, errors.New("Telegram Bot unavailable")
	}

	if strings.EqualFold(t.Token, "") {
		return &telegramBot.BotAPI{}, errors.New("Token telegram bot must be provided")
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
		log.Printf(err.Error())
		return ""
	}

	return string(result)
}

// Send message as simple
// message will be sent to only one channel or group via chat id
func (t *TelegramBot) SendMessage(message interface{}) (telegramBot.Message, error) {
	if t.ChatId == 0 {
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

// Send message to multi chat id (group or channel)
// refer `mode` from telegram_config.go
func (t *TelegramBot) SendMessages(mode string, message interface{}) ([]telegramBot.Message, error) {
	var response []telegramBot.Message

	if len(t.MultiChatId) == 0 {
		return response, errors.New("Multi chat Id must be provided")
	}

	bot, err := t.NewBot()
	if err != nil {
		log.Fatal(err)
		return response, err
	}

	key, ok := TelegramMessageMode[mode]

	if !ok {
		key = telegramBot.ModeHTML
	}

	content := t.ToJson(message)
	for _, v := range t.MultiChatId {
		_message := telegramBot.NewMessage(v, content)
		_message.ParseMode = key
		_response, err := bot.Send(_message)

		if err != nil {
			log.Fatal(err)
			return response, err
		}
		response = append(response, _response)
	}

	return response, nil
}

// Send message to multi chat id (group or channel)
// refer `mode` from telegram_config.go
func (t *TelegramBot) SendMessagesWith(message interface{}) ([]telegramBot.Message, error) {
	return t.SendMessages("html", message)
}
