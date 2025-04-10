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
	if e.Text != nil && t.state.Get(int(e.Chat.ID)) == StateConnected {
		t.SendMessage(methods.SendMessageReq{
			ChatID: int64(c.Get(int(e.Chat.ID))),
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

		c.AddToSearch(int(e.Chat.ID))

		go func() {
			for {
				if t.state.Get(int(e.Chat.ID)) != StateSearch {
					return
				}

				userId := c.GetFirstCompanion(int(e.Chat.ID))

				if userId != 0 {
					c.Connect(int(e.Chat.ID), userId)

					t.state.Set(int(e.Chat.ID), StateConnected)
					t.state.Set(userId, StateConnected)

					t.SendMessage(methods.SendMessageReq{
						ChatID: e.Chat.ID,
						Text:   fmt.Sprintf("New companion is found (id%d)", userId),
					})

					t.SendMessage(methods.SendMessageReq{
						ChatID: int64(userId),
						Text:   fmt.Sprintf("New companion is found (id%d)", e.Chat.ID),
					})

					return
				} else {
					time.Sleep(1000 * time.Millisecond)
				}
			}
		}()
	}

	if *e.Text == "/stop" {
		t.SendMessage(methods.SendMessageReq{
			Text:   "Goodbye!",
			ChatID: e.Chat.ID,
		})

		c.RemoveFromSearch(int(e.Chat.ID))
	}

	if *e.Text == "/ping" {
		before := time.Now().UnixMilli()

		m, _ := t.SendMessage(methods.SendMessageReq{
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
