package methods

type CopyMessage struct {
	ChatID              int    `json:"chat_id"`                         // Required
	MessageThreadID     int    `json:"message_thread_id,omitempty"`     // Optional
	FromChatID          int    `json:"from_chat_id"`                    // Required
	MessageID           int    `json:"message_id"`                      // Required
	VideoStartTimestamp int    `json:"video_start_timestamp,omitempty"` // Optional
	Caption             string `json:"caption,omitempty"`               // Optional
	ParseMode           string `json:"parse_mode,omitempty"`            // Optional
	// CaptionEntities       []MessageEntity  `json:"caption_entities,omitempty"`         // Optional
	ShowCaptionAboveMedia bool        `json:"show_caption_above_media,omitempty"` // Optional
	DisableNotification   bool        `json:"disable_notification,omitempty"`     // Optional
	ProtectContent        bool        `json:"protect_content,omitempty"`          // Optional
	AllowPaidBroadcast    bool        `json:"allow_paid_broadcast,omitempty"`     // Optional
	ReplyParameters       string      `json:"reply_parameters,omitempty"`         // Optional
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`             // Optional (supports multiple types)
}
