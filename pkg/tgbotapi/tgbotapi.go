package tgbotapi

type CallbackGame struct {
}
type InlineKeyboardButton struct {
	Text string
	URL *string
	CallbackData *string
	SwitchInlineQuery *string
	SwitchInlineQueryCurrentChat *string
	CallbackGame *CallbackGame
	Pay bool
}
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton
}
type KeyboardButton struct {
	Text string
	RequestContact bool
	RequestLocation bool
}
type ReplyKeyboardMarkup struct {
	Keyboard [][]KeyboardButton
	ResizeKeyboard bool
	OneTimeKeyboard bool
	Selective bool
}
