package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// Create a struct to conform to the JSON body of the send message request
// https://core.telegram.org/bots/api#sendmessage
type SendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func SendMessage(chatID int64, token, message string) error {
	// Create the request body struct
	reqBody := &SendMessageReqBody{
		ChatID: chatID,
		Text:   message,
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	// Send a post request with your token
	res, err := http.Post("https://api.telegram.org/bot"+token+"/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status:" + res.Status)
	}

	return nil
}

func Teste() string {
	return "OK"
}
