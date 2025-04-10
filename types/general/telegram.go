package general

type TelegramRes struct {
	Ok          bool    `json:"ok"`
	Result      *any    `json:"result,omitempty"`
	ErrorCode   *int    `json:"error_code,omitempty"`
	Description *string `json:"description,omitempty"`
}
