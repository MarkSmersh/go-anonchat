package general

type CallbackQuery struct {
	ID              string   `json:"id"`                          // REQUIRED: Unique identifier for the query
	From            User     `json:"from"`                        // REQUIRED: Sender
	Message         *Message `json:"message,omitempty"`           // OPTIONAL: Message sent by bot (can be inaccessible)
	InlineMessageID *string  `json:"inline_message_id,omitempty"` // OPTIONAL: Inline msg ID that originated query
	ChatInstance    string   `json:"chat_instance"`               // REQUIRED: Unique chat instance ID
	Data            *string  `json:"data,omitempty,omitempty"`    // OPTIONAL: Callback button data
	GameShortName   *string  `json:"game_short_name,omitempty"`   // OPTIONAL: Game short name
}
