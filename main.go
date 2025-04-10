package main

import (
	"fmt"
	"time"

	"github.com/MarkSmersh/go-anonchat/types/general"
	"github.com/MarkSmersh/go-anonchat/types/methods"
)

var t = Telegram{token: "", updateId: 0}
var c = Chat{}

const (
	StateDefault   = 0
	StateSearch    = 1
	StateConnected = 2
)

func main() {
	t.eventer.messages.Add(messageHandler)
	t.eventer.commands.Add(commandHandler)
	t.Init()
}

func messageHandler(e general.Message) {
	if e.Text != nil {
		t.SendMessage(methods.SendMessageReq{
			ChatID: e.Chat.ID,
			Text:   *e.Text,
		})
	}
}

func commandHandler(e general.Message) {
	if *e.Text == "/start" {
		t.state.Set(int(e.Chat.ID), StateSearch)

		t.SendMessage(methods.SendMessageReq{
			Text:   "Searching...",
			ChatID: e.Chat.ID,
		})
	}

	if *e.Text == "/ping" {
		before := time.Now().UnixMilli()

		m, _ := t.SendMessage(methods.SendMessageReq{
			Text:   "Ping...",
			ChatID: e.Chat.ID,
		})

		after := time.Now().UnixMilli()

		t.EditMessageText(methods.EditMessageText{
			Text:      fmt.Sprintf("Pong:%dms", after-before),
			MessageID: m.MessageID,
			ChatID:    m.Chat.ID,
		})
	}

	if *e.Text == "/stop" {
		t.SendMessage(methods.SendMessageReq{
			Text:   "Goodbye!",
			ChatID: e.Chat.ID,
		})
	}
}
