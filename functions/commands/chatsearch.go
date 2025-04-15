package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/MarkSmersh/go-anonchat/consts"
	"github.com/MarkSmersh/go-anonchat/core"
	"github.com/MarkSmersh/go-anonchat/helpers"
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

		text := "Searching...\n\nYou can stop a searching process with /stop"

		if len(c.Users[e.Chat.ID].Interests) > 0 {
			interests := []string{}

			for _, i := range c.Users[e.Chat.ID].Interests {
				interests = append(interests, helpers.InterestToStr(i))
			}

			text = fmt.Sprintf("Your interests: %s\n\n", strings.Join(interests, ", ")) + text
		}

		t.SendMessage(methods.SendMessage{
			Text:   text,
			ChatID: e.Chat.ID,
		})

		c.AddToSearch(e.Chat.ID)

		go func() {
			for {
				if t.State.Get(e.Chat.ID) != consts.StateSearch {
					return
				}

				userId, equalInterests := c.GetFirstCompanion(e.Chat.ID)

				interests := []string{}

				for _, i := range equalInterests {
					interests = append(interests, helpers.InterestToStr(i))
				}

				if userId != 0 {
					c.Connect(e.Chat.ID, userId)

					t.State.Set(e.Chat.ID, consts.StateConnected)
					t.State.Set(userId, consts.StateConnected)

					text := "New companion is found (id%d)"

					if len(equalInterests) > 0 {
						text = fmt.Sprintf("Equal interests: %s\n\n", strings.Join(interests, ", ")) + text
					}

					t.SendMessage(methods.SendMessage{
						ChatID: e.Chat.ID,
						Text:   fmt.Sprintf(text, userId),
					})

					t.SendMessage(methods.SendMessage{
						ChatID: userId,
						Text:   fmt.Sprintf(text, e.Chat.ID),
					})

					return
				} else {
					time.Sleep(1000 * time.Millisecond)
				}
			}
		}()
	}
}
