package commands

import (
	"fmt"

	"github.com/MarkSmersh/go-anonchat/core"
	"github.com/MarkSmersh/go-anonchat/types/general"
	"github.com/MarkSmersh/go-anonchat/types/methods"
)

func Start(t *core.Telegram, c *core.Chat, e *general.Message) {
	if *e.Text == "/start" {
		t.SendMessage(methods.SendMessage{
			ChatID: e.Chat.ID,
			Text:   fmt.Sprintf("Hello, %s. Welcome to the edging zone.", e.From.FirstName),
		})
	}
}
