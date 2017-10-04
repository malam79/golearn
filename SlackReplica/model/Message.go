package model

import (
	_ "fmt"
)

type Message struct {
	Message    string `json: "Message"`
	DestUserID int    `json: "To"`
	SrcUserID  int    `json: "From"`
	Time       string `json: "Time"`
}
