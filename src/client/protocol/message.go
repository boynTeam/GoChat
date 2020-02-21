package protocol

import "encoding/json"

// Author:Boyn
// Date:2020/2/21

// 消息结构的设计

type Message struct {
	Content string `json:"content"`
}

func (m *Message) ToJSON() ([]byte, error) {
	marshal, err := json.Marshal(m)
	return marshal, err
}
