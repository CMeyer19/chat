package main

import (
	"io"
	"net/http"

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

func GetConversationMessages(conversationId string) ConversationResult {
	body := ConversationResult{
		Results: []Message{
			{
				Content: "Hi Cole",
				UserId:  "a",
			},
			{
				Content: "Hi Meghan",
				UserId:  "b",
			},
		},
	}

	return body
}
