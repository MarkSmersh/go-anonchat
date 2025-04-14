package keyboard

type InlineButton struct {
	Text         string `json:"text"`                    // Required: Label text
	URL          string `json:"url,omitempty"`           // Optional: URL to open
	CallbackData string `json:"callback_data,omitempty"` // Optional: Callback data (1–64 bytes)
	// WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`                          // Optional: Web App descriptor
	// LoginURL                     *LoginURL                    `json:"login_url,omitempty"`                        // Optional: Login URL
	SwitchInlineQuery            string `json:"switch_inline_query,omitempty"`              // Optional: Insert inline query in ANY chat
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"` // Optional: Insert inline query in CURRENT chat
	// SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`  // Optional: Inline query + chat type selector
	// CopyText                     *CopyTextButton              `json:"copy_text,omitempty"`                        // Optional: Copy text to clipboard
	// CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`                    // Optional: Launch game (must be first in row)
	Pay bool `json:"pay,omitempty"` // Optional: Pay button (first in row, invoice only)
}
