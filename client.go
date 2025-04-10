package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/MarkSmersh/go-anonchat/types/general"
	"github.com/MarkSmersh/go-anonchat/types/methods"
)

type Telegram struct {
	token    string
	updateId int
	eventer  Eventer[general.Update]
}

func (t *Telegram) Request(method string, params interface{}) ([]byte, error) {
	paramsString := ""

	if params != nil {
		var paramsMap map[string]interface{}

		tmp, _ := json.Marshal(params)

		d := json.NewDecoder(strings.NewReader(string(tmp[:])))

		d.UseNumber()

		d.Decode(&paramsMap)

		for k, v := range paramsMap {
			paramsString += fmt.Sprintf("%s=%v&", k, v)
		}
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s?%s", t.token, method, paramsString)

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

	if !result.Ok {
		log.Println("Telegram Bad Response")
		return resultBytes, errors.New("Telegram Bad Response")
	}

	return resultBytes, nil
}

func (t *Telegram) Init() {
	t.Polling()
}

func (t *Telegram) GetMe() (methods.GetMeRes, error) {
	result, _ := t.Request("getMe", nil)
	data := methods.GetMeRes{}
	json.Unmarshal(result, &data)
	return data, nil
}

func (t *Telegram) SendMessage(params methods.SendMessageReq) (general.Message, error) {
	result, _ := t.Request("sendMessage", params)
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

func (t *Telegram) Polling() {
	for {
		req := methods.GetUpdates{
			Offset: &t.updateId,
		}

		updates, _ := t.GetUpdates(req)

		for i := range updates {
			u := updates[i]

			t.eventer.Invoke(general.UpdateMessage, u)
		}

		if len(updates) <= 0 {
			continue
		}

		t.updateId = updates[len(updates)-1].UpdateID + 1
	}
}
