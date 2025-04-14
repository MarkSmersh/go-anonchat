package commands

import (
	"strconv"

	"github.com/MarkSmersh/go-anonchat/consts"
	"github.com/MarkSmersh/go-anonchat/core"
	"github.com/MarkSmersh/go-anonchat/core/keyboard"
	"github.com/MarkSmersh/go-anonchat/helpers"
	"github.com/MarkSmersh/go-anonchat/types/general"
	"github.com/MarkSmersh/go-anonchat/types/methods"
)

func Interests(t *core.Telegram, c *core.Chat, e *general.Message) {
	if *e.Text == "/interests" {
		u := c.Users[e.From.ID].Interests

		keyboard := keyboard.ReplyMarkup{
			InlineButtons: [][]keyboard.InlineButton{
				{
					keyboard.InlineButton{
						Text:         helpers.InterestToStr(consts.InterestTalking) + " " + helpers.IsInterestIn(u, consts.InterestTalking),
						CallbackData: "i-" + strconv.Itoa(consts.InterestTalking),
					},
					keyboard.InlineButton{
						Text:         helpers.InterestToStr(consts.InterestSex) + " " + helpers.IsInterestIn(u, consts.InterestSex),
						CallbackData: "i-" + strconv.Itoa(consts.InterestSex),
					},
				},
			},
		}

		t.SendMessage(methods.SendMessage{
			ChatID:      e.Chat.ID,
			Text:        "Choose your interesets from below:",
			ReplyMarkup: keyboard.ToJSON(),
		})
	}
}
