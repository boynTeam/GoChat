package internal

import "sync"

// Author:Boyn
// Date:2020/2/20

var (
	Entering = make(chan Client) // 监控客户端进入的消息
	Leaving  = make(chan Client) // 监控客户端离开的消息
	Messages = make(chan string) // 掌握所有客户端发出的消息
	Clients  sync.Map            //掌握所有客户端消息
)
