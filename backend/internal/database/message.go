package database

import "encoding/json"

type Message struct {
	Text string `json:"text"`
	Id   uint   `json:"id"`
}

func (m Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

var nextId uint = 0

func NewMessage(text string) Message {
	msg := Message{
		Text: text,
		Id:   nextId,
	}
	nextId += 1

	return msg
}
