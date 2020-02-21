package internal

import "sync"

// Author:Boyn
// Date:2020/2/20

var (
	Entering    = make(chan Client)  // 监控客户端进入的消息
	Leaving     = make(chan Client)  // 监控客户端离开的消息
	BroadCaster = make(chan Message) // 广播器,将传入广播器通道的消息传到所有活跃的客户端中
	Clients     sync.Map             //掌握所有客户端消息
)
