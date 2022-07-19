package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Create a struct to conform to the JSON body of the send message request
// https://core.telegram.org/bots/api#sendmessage
type SendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type Config struct {
	ID    int64  `json:"ID"`
	Token string `json:"Token"`
}

func ReadConfig(fileName string) (Config, error) {
	cfg := Config{}
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return cfg, fmt.Errorf("readConfig: %w", err)
	}
	reader := strings.NewReader(string(b))

	if err := json.NewDecoder(reader).Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("readConfig: %w", err)
	}
	return cfg, nil
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
