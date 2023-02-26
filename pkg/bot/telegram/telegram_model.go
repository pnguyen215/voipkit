package telegram

type TelegramBot struct {
	Enabled   bool   `json:"enabled"`
	Token     string `json:"telegram_bot_token"`
	ChatId    int64  `json:"telegram_chat_id"`
	DebugMode bool   `json:"debug_mode"`
}
