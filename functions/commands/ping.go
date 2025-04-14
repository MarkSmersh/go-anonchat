package commands

import (
	"fmt"
	"time"

	"github.com/MarkSmersh/go-anonchat/core"
	"github.com/MarkSmersh/go-anonchat/types/general"
	"github.com/MarkSmersh/go-anonchat/types/methods"
)

func Ping(t *core.Telegram, c *core.Chat, e *general.Message) {
	if *e.Text == "/ping" {
		before := time.Now().UnixMilli()

		m, _ := t.SendMessage(methods.SendMessage{
			Text:   "Ping...",
			ChatID: e.Chat.ID,
		})

		after := time.Now().UnixMilli()

		t.EditMessageText(methods.EditMessageText{
			Text:      fmt.Sprintf("Pong: %d ms", after-before),
			MessageID: m.MessageID,
			ChatID:    m.Chat.ID,
		})
	}
}
