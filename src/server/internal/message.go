package internal

import (
	"encoding/json"
	"time"
)

// Author:Boyn
// Date:2020/2/21

// 消息结构的设计

type Message struct {
	Content string `json:"content"`
	User    string `json:"user"`
	Time    string `json:"time"`
	State   int    `json:"type"`
}

func (m *Message) ToJSON() ([]byte, error) {
	m.Time = time.Now().Format("15:04:05")
	marshal, err := json.Marshal(m)
	return marshal, err
}
