package main

import (
	"github.com/MarkSmersh/go-anonchat/types/general"
	"github.com/MarkSmersh/go-anonchat/types/methods"
)

var t = Telegram{token: "", updateId: 0}

func main() {
	type R struct {
		test  bool
		test1 bool
	}

	t.GetMe()
	t.SendMessage(methods.SendMessageReq{ChatID: 562140704, Text: "Hello!"})

	t.eventer.Add(general.UpdateMessage, messageHandler)
	t.Init()
}

func messageHandler(u general.Update) {
	if u.Message != nil && u.Message.Text != nil {
		println(u.Message.Text)

		req := methods.SendMessageReq{
			ChatID: u.Message.Chat.ID,
			Text:   *u.Message.Text,
		}

		t.SendMessage(req)
	}

	// http.HandleFunc()
}
