package methods

type SendMessage struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"` // Optional business connection ID
	ChatID               int    `json:"chat_id"`                          // Required: int64 or string (username)
	MessageThreadID      int    `json:"message_thread_id,omitempty"`      // Optional forum thread ID
	Text                 string `json:"text"`                             // Required message text (1–4096 chars)
	ParseMode            string `json:"parse_mode,omitempty"`             // Optional formatting mode
	// Entities             *[]MessageEntity    `json:"entities,omitempty"`               // Optional list of special entities
	// LinkPreviewOptions   *LinkPreviewOptions `json:"link_preview_options,omitempty"`   // Optional link preview controls
	DisableNotification bool   `json:"disable_notification,omitempty"` // Optional silent send
	ProtectContent      bool   `json:"protect_content,omitempty"`      // Optional forward protection
	AllowPaidBroadcast  bool   `json:"allow_paid_broadcast,omitempty"` // Optional override broadcast limits
	MessageEffectID     string `json:"message_effect_id,omitempty"`    // Optional special effect ID (private chats only)
	// ReplyParameters      *ReplyParameters    `json:"reply_parameters,omitempty"`       // Optional reply meta
	ReplyMarkup string `json:"reply_markup,omitempty"` // InlineKeyboardMarkup | ReplyKeyboardMarkup | ReplyKeyboardRemove | ForceReply
}
