package keyboard

import "encoding/json"

type ReplyMarkup struct {
	InlineButtons [][]InlineButton `json:"inline_keyboard"`
}

func (r *ReplyMarkup) ToJSON() string {
	v, _ := json.Marshal(r)

	return string(v[:])
}
