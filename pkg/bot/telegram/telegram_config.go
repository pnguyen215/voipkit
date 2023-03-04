package telegram

import telegramBot "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	TelegramMessageMode map[string]string = map[string]string{
		"markdown_v1": telegramBot.ModeMarkdown,
		"markdown_v2": telegramBot.ModeMarkdownV2,
		"html":        telegramBot.ModeHTML,
	}
)
