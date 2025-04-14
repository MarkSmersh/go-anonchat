package methods

type AnswerCallbackQuery struct {
	CallbackQueryID string `json:"callback_query_id"`    // REQUIRED: Unique ID for the query
	Text            string `json:"text,omitempty"`       // Optional: Notification text (0–200 chars)
	ShowAlert       bool   `json:"show_alert,omitempty"` // Optional: Show alert instead of toast
	URL             string `json:"url,omitempty"`        // Optional: URL to open
	CacheTime       int    `json:"cache_time,omitempty"` // Optional: Max cache duration in seconds (default 0)
}
