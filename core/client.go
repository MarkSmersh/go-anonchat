package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/MarkSmersh/go-anonchat/types/general"
	"github.com/MarkSmersh/go-anonchat/types/methods"
)

type Telegram struct {
	Token    string
	UpdateId int
	State    State[int, int]
	Eventer  Updater
}

func (t *Telegram) Request(method string, params interface{}) ([]byte, error) {
	paramsValues := url.Values{}

	if params != nil {
		var paramsMap map[string]interface{}

		tmp, _ := json.Marshal(params)

		d := json.NewDecoder(strings.NewReader(string(tmp[:])))

		d.UseNumber()

		d.Decode(&paramsMap)

		for k, v := range paramsMap {
			paramsValues.Add(k, fmt.Sprintf("%v", v))
		}
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s?%s", t.Token, method, paramsValues.Encode())

	// println(url)

	res, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return []byte{}, err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
		return []byte{}, err
	}

	result := general.TelegramRes{}

	json.Unmarshal(body, &result)

	resultBytes, _ := json.Marshal(result.Result)

	// println(string(body[:]))

	if !result.Ok {
		log.Println(
			fmt.Sprintf("Telegram error. Code: %d. Description: %s", *result.ErrorCode, *result.Description),
		)
		return resultBytes, errors.New("Telegram Bad Response")
	}

	return resultBytes, nil
}

func (t *Telegram) Init(callback func(e general.User)) {
	if callback != nil {
		me, _ := t.GetMe()
		callback(me)
	}
	t.Polling()
}

func (t *Telegram) GetMe() (general.User, error) {
	result, _ := t.Request("getMe", nil)
	data := general.User{}
	json.Unmarshal(result, &data)
	return data, nil
}

func (t *Telegram) SendMessage(params methods.SendMessage) (general.Message, error) {
	result, _ := t.Request("sendMessage", params)
	data := general.Message{}
	json.Unmarshal(result, &data)
	return data, nil
}

func (t *Telegram) ForwardMessage(params methods.ForwardMessage) (general.Message, error) {
	result, _ := t.Request("forwardMessage", params)
	data := general.Message{}
	json.Unmarshal(result, &data)
	return data, nil
}

func (t *Telegram) CopyMessage(params methods.CopyMessage) (general.Message, error) {
	result, _ := t.Request("copyMessage", params)
	data := general.Message{}
	json.Unmarshal(result, &data)
	return data, nil
}

func (t *Telegram) GetUpdates(params methods.GetUpdates) ([]general.Update, error) {
	result, _ := t.Request("getUpdates", params)
	data := []general.Update{}
	json.Unmarshal(result, &data)
	return data, nil
}

func (t *Telegram) EditMessageText(params methods.EditMessageText) (general.Message, error) {
	result, _ := t.Request("editMessageText", params)
	data := general.Message{}
	json.Unmarshal(result, &data)
	return data, nil
}

func (t *Telegram) EditMessageReplyMarkup(params methods.EditMessageReplyMarkup) (general.Message, error) {
	result, _ := t.Request("editMessageReplyMarkup", params)
	data := general.Message{}
	json.Unmarshal(result, &data)
	return data, nil
}

func (t *Telegram) AnswerCallbackQuery(params methods.AnswerCallbackQuery) error {
	t.Request("answerCallbackQuery", params)
	return nil
}

func (t *Telegram) Polling() {
	for {
		req := methods.GetUpdates{
			Offset: t.UpdateId,
		}

		updates, _ := t.GetUpdates(req)

		for i := range updates {
			u := updates[i]

			if u.Message != nil {
				e := u.Message

				if e.Text != nil && (*e.Text)[0] == '/' {
					t.Eventer.Commands.Invoke(*u.Message)
					break
				}

				t.Eventer.Messages.Invoke(*u.Message)
				break
			}

			if u.InlineQuery != nil {
				t.Eventer.InlineQuery.Invoke(*u.InlineQuery)
			}

			if u.CallbackQuery != nil {
				t.Eventer.CallbackQuery.Invoke(*u.CallbackQuery)
			}
		}

		if len(updates) <= 0 {
			continue
		}

		t.UpdateId = updates[len(updates)-1].UpdateID + 1
	}
}
