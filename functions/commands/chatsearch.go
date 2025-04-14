package commands

import (
	"fmt"
	"time"

	"github.com/MarkSmersh/go-anonchat/consts"
	"github.com/MarkSmersh/go-anonchat/core"
	"github.com/MarkSmersh/go-anonchat/types/general"
	"github.com/MarkSmersh/go-anonchat/types/methods"
)

func ChatSearch(t *core.Telegram, c *core.Chat, e *general.Message) {
	if *e.Text == "/next" {
		if t.State.Get(e.Chat.ID) == consts.StateSearch {
			t.SendMessage(methods.SendMessage{
				Text:   "You are already searching for a companion. There is no need to spam it.",
				ChatID: e.Chat.ID,
			})
			return
		}

		if t.State.Get(e.Chat.ID) == consts.StateConnected {
			t.SendMessage(methods.SendMessage{
				Text:   "You already have a companion. If you don't like use /stop and then /next. (it will be changed quite soon)",
				ChatID: e.Chat.ID,
			})
			return
		}

		t.State.Set(e.Chat.ID, consts.StateSearch)

		t.SendMessage(methods.SendMessage{
			Text:   "Searching...\n\nUse /stop to stop searching process",
			ChatID: e.Chat.ID,
		})

		c.AddToSearch(e.Chat.ID)

		go func() {
			for {
				if t.State.Get(e.Chat.ID) != consts.StateSearch {
					return
				}

				userId, _ := c.GetFirstCompanion(e.Chat.ID)

				if userId != 0 {
					c.Connect(e.Chat.ID, userId)

					t.State.Set(e.Chat.ID, consts.StateConnected)
					t.State.Set(userId, consts.StateConnected)

					t.SendMessage(methods.SendMessage{
						ChatID: e.Chat.ID,
						Text:   fmt.Sprintf("New companion is found (id%d)", userId),
					})

					t.SendMessage(methods.SendMessage{
						ChatID: userId,
						Text:   fmt.Sprintf("New companion is found (id%d)", e.Chat.ID),
					})

					return
				} else {
					time.Sleep(1000 * time.Millisecond)
				}
			}
		}()
	}
}
