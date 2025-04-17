package general

import "encoding/json"

type ReplyParameters struct {
	MessageID                int     `json:"message_id"`                            // REQUIRED: Message to be replied to
	ChatID                   int     `json:"chat_id,omitempty"`                     // OPTIONAL: Different chat ID (int or string)
	AllowSendingWithoutReply *bool   `json:"allow_sending_without_reply,omitempty"` // OPTIONAL: Send even if original message not found
	Quote                    *string `json:"quote,omitempty"`                       // OPTIONAL: Exact substring quote
	QuoteParseMode           *string `json:"quote_parse_mode,omitempty"`            // OPTIONAL: Parse mode for quote
	// QuoteEntities            []MessageEntity `json:"quote_entities,omitempty"`              // OPTIONAL: Special formatting entities in quote
	QuotePosition *int `json:"quote_position,omitempty"` // OPTIONAL: Position of the quote in UTF-16 units
}

func (r *ReplyParameters) ToJSON() string {
	v, _ := json.Marshal(r)

	return string(v[:])
}
