package general

type InlineQuery struct {
	ID       string  `json:"id"`                  // Unique query identifier
	From     User    `json:"from"`                // Sender info
	Query    string  `json:"query"`               // The text of the query
	Offset   string  `json:"offset"`              // Offset for pagination
	ChatType *string `json:"chat_type,omitempty"` // Optional: type of the originating chat
	// Location *Location `json:"location,omitempty"`  // Optional: sender location
}
