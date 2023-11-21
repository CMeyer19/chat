package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2/log"
)

type Message struct {
	Content string `json:"content"`
	UserId  string `json:"userId"`
}

type ConversationResult struct {
	Results []Message `json:"results"`
}

func Fetch(path string) string {
	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func GetConversationMessages(conversationId string) []Message {
	body := Fetch(TickerPath + "?" + ApiKey + "&ticker=" + strings.ToUpper(conversationId))

	data := ConversationResult{}
	json.Unmarshal([]byte(string(body)), &data)

	return data.Results
}
