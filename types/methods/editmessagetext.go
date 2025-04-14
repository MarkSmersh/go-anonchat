package methods

type EditMessageText struct {
	BusinessConnectionID string      `json:"business_connection_id,omitempty"`
	ChatID               interface{} `json:"chat_id,omitempty"` // int64 or string (@channelusername)
	MessageID            int         `json:"message_id,omitempty"`
	InlineMessageID      string      `json:"inline_message_id,omitempty"`
	Text                 string      `json:"text"`                 // Required: 1-4096 chars
	ParseMode            string      `json:"parse_mode,omitempty"` // e.g. "MarkdownV2", "HTML"
	// Entities             []MessageEntity       `json:"entities,omitempty"`
	// LinkPreviewOptions   *LinkPreviewOptions   `json:"link_preview_options,omitempty"`
	ReplyMarkup string `json:"reply_markup,omitempty"`
}
