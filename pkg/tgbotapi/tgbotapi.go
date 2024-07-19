// Code generated DO NOT EDIT. (@generated)

package tgbotapi

import "net/url"

type Animation struct {
	FileID string
	Thumb PhotoSize
	FileName string
	MimeType string
	FileSize int
}
type Audio struct {
	FileID string
	Duration int
	Performer string
	Title string
	MimeType string
	FileSize int
}
type CallbackGame struct {
}
type Chat struct {
	ID int64
	Type string
	Title string
	UserName string
	FirstName string
	LastName string
	AllMembersAreAdmins bool
	Photo *ChatPhoto
	Description string
	InviteLink string
}
func (Chat) ChatConfig() ChatConfig { panic("not implemented") }
func (Chat) IsChannel() bool { panic("not implemented") }
func (Chat) IsGroup() bool { panic("not implemented") }
func (Chat) IsPrivate() bool { panic("not implemented") }
func (Chat) IsSuperGroup() bool { panic("not implemented") }
type ChatAnimation struct {
	FileID string
	Width int
	Height int
	Duration int
	Thumbnail *PhotoSize
	FileName string
	MimeType string
	FileSize int
}
type ChatConfig struct {
	ChatID int64
	SuperGroupUsername string
}
type ChatPhoto struct {
	SmallFileID string
	BigFileID string
}
type Contact struct {
	PhoneNumber string
	FirstName string
	LastName string
	UserID int
}
type Document struct {
	FileID string
	Thumbnail *PhotoSize
	FileName string
	MimeType string
	FileSize int
}
type EncryptedCredentials struct {
	Data string
	Hash string
	Secret string
}
type EncryptedPassportElement struct {
	Type string
	Data string
	PhoneNumber string
	Email string
	Files []PassportFile
	FrontSide *PassportFile
	ReverseSide *PassportFile
	Selfie *PassportFile
}
type Game struct {
	Title string
	Description string
	Photo []PhotoSize
	Text string
	TextEntities []MessageEntity
	Animation Animation
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
type Invoice struct {
	Title string
	Description string
	StartParameter string
	Currency string
	TotalAmount int
}
type KeyboardButton struct {
	Text string
	RequestContact bool
	RequestLocation bool
}
type Location struct {
	Longitude float64
	Latitude float64
}
type Message struct {
	MessageID int
	From *User
	Date int
	Chat *Chat
	ForwardFrom *User
	ForwardFromChat *Chat
	ForwardFromMessageID int
	ForwardDate int
	ReplyToMessage *Message
	EditDate int
	Text string
	Entities *[]MessageEntity
	Audio *Audio
	Document *Document
	Animation *ChatAnimation
	Game *Game
	Photo *[]PhotoSize
	Sticker *Sticker
	Video *Video
	VideoNote *VideoNote
	Voice *Voice
	Caption string
	Contact *Contact
	Location *Location
	Venue *Venue
	NewChatMembers *[]User
	LeftChatMember *User
	NewChatTitle string
	NewChatPhoto *[]PhotoSize
	DeleteChatPhoto bool
	GroupChatCreated bool
	SuperGroupChatCreated bool
	ChannelChatCreated bool
	MigrateToChatID int64
	MigrateFromChatID int64
	PinnedMessage *Message
	Invoice *Invoice
	SuccessfulPayment *SuccessfulPayment
	PassportData *PassportData
}
type MessageEntity struct {
	Type string
	Offset int
	Length int
	URL string
	User *User
}
func (MessageEntity) ParseURL() (*url.URL, error) { panic("not implemented") }
type OrderInfo struct {
	Name string
	PhoneNumber string
	Email string
	ShippingAddress *ShippingAddress
}
type PassportData struct {
	Data []EncryptedPassportElement
	Credentials *EncryptedCredentials
}
type PassportFile struct {
	FileID string
	FileSize int
	FileDate int64
}
type PhotoSize struct {
	FileID string
	Width int
	Height int
	FileSize int
}
type ReplyKeyboardMarkup struct {
	Keyboard [][]KeyboardButton
	ResizeKeyboard bool
	OneTimeKeyboard bool
	Selective bool
}
type ShippingAddress struct {
	CountryCode string
	State string
	City string
	StreetLine1 string
	StreetLine2 string
	PostCode string
}
type Sticker struct {
	FileID string
	Width int
	Height int
	Thumbnail *PhotoSize
	Emoji string
	FileSize int
	SetName string
}
type SuccessfulPayment struct {
	Currency string
	TotalAmount int
	InvoicePayload string
	ShippingOptionID string
	OrderInfo *OrderInfo
	TelegramPaymentChargeID string
	ProviderPaymentChargeID string
}
type User struct {
	ID int
	FirstName string
	LastName string
	UserName string
	LanguageCode string
	IsBot bool
}
type Venue struct {
	Location Location
	Title string
	Address string
	FoursquareID string
}
type Video struct {
	FileID string
	Width int
	Height int
	Duration int
	Thumbnail *PhotoSize
	MimeType string
	FileSize int
}
type VideoNote struct {
	FileID string
	Length int
	Duration int
	Thumbnail *PhotoSize
	FileSize int
}
type Voice struct {
	FileID string
	Duration int
	MimeType string
	FileSize int
}
