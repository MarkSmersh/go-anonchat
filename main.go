package main

import (
	"fmt"

	"github.com/MarkSmersh/go-anonchat/consts"
	"github.com/MarkSmersh/go-anonchat/core"
	"github.com/MarkSmersh/go-anonchat/functions/commands"
	"github.com/MarkSmersh/go-anonchat/functions/inline"
	"github.com/MarkSmersh/go-anonchat/helpers"
	"github.com/MarkSmersh/go-anonchat/types/general"
	"github.com/MarkSmersh/go-anonchat/types/methods"
)

var env, _ = helpers.GetEnv()

var t = core.Telegram{Token: env["BOT_TOKEN"], UpdateId: 0}
var c = core.Chat{}

func main() {
	t.Eventer.Messages.Add(messageHandler)
	t.Eventer.Commands.Add(commandHandler)
	t.Eventer.CallbackQuery.Add(callbackHandler)
	t.Init(onInit)
}

func onInit(me general.User) {
	c.Users = map[int]*core.User{}

	fmt.Printf("Bot %s is started", me.FirstName)
}

func messageHandler(e general.Message) {
	if t.State.Get(e.Chat.ID) == consts.StateConnected {
		req := methods.CopyMessage{
			ChatID:     c.Get(e.Chat.ID),
			FromChatID: e.Chat.ID,
			MessageID:  e.MessageID,
		}

		if e.ReplyToMessage != nil {

			tr := true
			messageId := c.GetMessageA(e.ReplyToMessage.MessageID)
			from := c.Get(e.Chat.ID)

			if e.ReplyToMessage.From.ID == e.From.ID {
				messageId = c.GetMessageB(e.ReplyToMessage.MessageID)
			}

			rp := general.ReplyParameters{
				MessageID:                messageId,
				ChatID:                   from,
				AllowSendingWithoutReply: &tr,
			}

			req.ReplyParameters = rp.ToJSON()
		}

		mes, _ := t.CopyMessage(req)

		c.AddMessage(mes.MessageID, e.MessageID)
	}
}

func commandHandler(e general.Message) {
	_, ok := c.Users[e.Chat.ID]

	if !ok {
		c.Users[e.Chat.ID] = &core.User{Id: e.Chat.ID, Interests: []int{}, Age: 0, Companion: 0, Sex: 0}
	}

	commands.Start(&t, &c, &e)

	commands.ChatSearch(&t, &c, &e)

	commands.Stop(&t, &c, &e)

	commands.Ping(&t, &c, &e)

	commands.Interests(&t, &c, &e)
}

func callbackHandler(e general.CallbackQuery) {
	_, ok := c.Users[e.From.ID]

	if !ok {
		c.Users[e.From.ID] = &core.User{Id: e.From.ID, Interests: []int{}, Age: 0, Companion: 0, Sex: 0}
	}

	inline.Interests(&t, &c, &e)
}
