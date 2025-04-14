package methods

type EditMessageReplyMarkup struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"` // Optional: Biz connect ID
	ChatID               int    `json:"chat_id,omitempty"`                // Optional: Chat ID or @channelusername
	MessageID            int    `json:"message_id,omitempty"`             // Optional: Message ID
	InlineMessageID      string `json:"inline_message_id"`                // Optional: Inline msg ID
	ReplyMarkup          string `json:"reply_markup,omitempty"`           // Optional: Inline keyboard markup
}
