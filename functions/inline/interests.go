package inline

import (
	"strconv"
	"strings"

	"github.com/MarkSmersh/go-anonchat/consts"
	"github.com/MarkSmersh/go-anonchat/core"
	"github.com/MarkSmersh/go-anonchat/core/keyboard"
	"github.com/MarkSmersh/go-anonchat/helpers"
	"github.com/MarkSmersh/go-anonchat/types/general"
	"github.com/MarkSmersh/go-anonchat/types/methods"
)

func Interests(t *core.Telegram, c *core.Chat, e *general.CallbackQuery) {
	splittedQuery := strings.Split(*e.Data, "-")

	prefix := splittedQuery[0]
	value := splittedQuery[1]

	if prefix != "i" {
		return
	}

	v, err := strconv.Atoi(value)

	if err != nil {
		return
	}

	c.Users[e.From.ID].AddOrRemoveInterest(v)

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

	t.EditMessageReplyMarkup(methods.EditMessageReplyMarkup{
		ReplyMarkup: keyboard.ToJSON(),
		MessageID:   e.Message.MessageID,
		ChatID:      e.From.ID,
	})
}
