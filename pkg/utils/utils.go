package utils

type TelegramCallbackQuery struct {
	ID string
	From struct { ID int64 "json:\"id\""; IsBot bool "json:\"is_bot\""; FirstName string "json:\"first_name\""; LastName string "json:\"last_name\""; LanguageCode string "json:\"language_code\"" }
	Message struct { MessageID int "json:\"message_id\""; From struct { ID int64 "json:\"id\""; IsBot bool "json:\"is_bot\""; FirstName string "json:\"first_name\""; Username string "json:\"username\"" } "json:\"from\""; Chat struct { ID int64 "json:\"id\""; FirstName string "json:\"first_name\""; LastName string "json:\"last_name\""; Type string "json:\"type\"" } "json:\"chat\""; Date int "json:\"date\""; Text string "json:\"text\""; ReplyMarkup struct { InlineKeyboard [][]struct { Text string "json:\"text\""; CallbackData string "json:\"callback_data\"" } "json:\"inline_keyboard\"" } "json:\"reply_markup\"" }
	ChatInstance string
	Data string
}
type TelegramMessage struct {
	MessageID int64
	From struct { ID int64 "json:\"id\""; IsBot bool "json:\"is_bot\""; FirstName string "json:\"first_name\""; LastName string "json:\"last_name\""; Username string "json:\"username\""; LanguageCode string "json:\"language_code\"" }
	Chat struct { ID int64 "json:\"id\""; FirstName string "json:\"first_name\""; LastName string "json:\"last_name\""; Username string "json:\"username\""; Type string "json:\"type\"" }
	ReplyToMessage struct { MessageID int "json:\"message_id\""; From struct { ID int64 "json:\"id\""; IsBot bool "json:\"is_bot\""; FirstName string "json:\"first_name\""; LastName string "json:\"last_name\""; LanguageCode string "json:\"language_code\"" } "json:\"from\""; Chat struct { ID int64 "json:\"id\""; FirstName string "json:\"first_name\""; LastName string "json:\"last_name\""; Type string "json:\"type\"" } "json:\"chat\""; Date int "json:\"date\""; Text string "json:\"text\"" }
	Date int64
	Text string
}
type TelegramUpdate struct {
	UpdateID int64
	Message TelegramMessage
	CallbackQuery TelegramCallbackQuery
}
