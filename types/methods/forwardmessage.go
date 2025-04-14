package methods

type ForwardMessage struct {
	ChatID              int  `json:"chat_id"`                         // Required: int64 or string (@channelusername)
	MessageThreadID     int  `json:"message_thread_id,omitempty"`     // Optional: for forum topics
	FromChatID          int  `json:"from_chat_id"`                    // Required: int64 or string
	VideoStartTimestamp int  `json:"video_start_timestamp,omitempty"` // Optional: for forwarding video with custom start
	DisableNotification bool `json:"disable_notification,omitempty"`  // Optional: silent mode ON
	ProtectContent      bool `json:"protect_content,omitempty"`       // Optional: protect against re-forward/save
	MessageID           int  `json:"message_id"`                      // Required
}
