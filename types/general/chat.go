package general

type Chat struct {
	ID        int     `json:"id"`                   // 64-bit safe chat ID
	Type      string  `json:"type"`                 // "private", "group", "supergroup", "channel"
	Title     *string `json:"title,omitempty"`      // Title for groups, channels
	Username  *string `json:"username,omitempty"`   // Username if available
	FirstName *string `json:"first_name,omitempty"` // First name in private chat
	LastName  *string `json:"last_name,omitempty"`  // Last name in private chat
	IsForum   *bool   `json:"is_forum,omitempty"`   // True if it's a forum (supergroup with topics)
}
