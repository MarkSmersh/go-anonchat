package commands

import (
	"github.com/MarkSmersh/go-anonchat/consts"
	"github.com/MarkSmersh/go-anonchat/core"
	"github.com/MarkSmersh/go-anonchat/types/general"
	"github.com/MarkSmersh/go-anonchat/types/methods"
)

func Stop(t *core.Telegram, c *core.Chat, e *general.Message) {
	if *e.Text == "/stop" {
		if t.State.Get(e.Chat.ID) == consts.StateDefault {
			t.SendMessage(methods.SendMessage{
				Text:   "You have no companion. You are not even searching for him. What are trying to do?",
				ChatID: e.Chat.ID,
			})
			return
		}

		c.RemoveFromSearch(e.Chat.ID)
		b := c.Disconnect(e.Chat.ID)
		t.State.Set(e.Chat.ID, consts.StateDefault)

		if b != 0 {
			t.State.Set(b, consts.StateDefault)

			t.SendMessage(methods.SendMessage{
				ChatID: e.Chat.ID,
				Text:   "You've stopped the dialogue. Use /next to search a new one",
			})

			t.SendMessage(methods.SendMessage{
				ChatID: b,
				Text:   "You've been skipped by your companion. There is no need to worry. You can find another one with /next",
			})
		} else {
			t.SendMessage(methods.SendMessage{
				ChatID: e.Chat.ID,
				Text:   "You've stopped searching process",
			})
		}
	}
}
