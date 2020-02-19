package chat

import "log"

/**
author:Boyn
date:2020/2/15
*/

var (
	entering = make(chan client)       // 监控客户端进入的消息
	leaving  = make(chan client)       // 监控客户端离开的消息
	messages = make(chan string)       // 掌握所有客户端发出的消息
	clients  = make(map[string]client) //掌握所有客户端消息
)

func broadcaster() {
	for {
		select {
		case msg := <-messages:
			log.Println(msg)
			for _, cli := range clients {
				cli.channel <- msg
			}
		case cli := <-entering:
			log.Printf("%s login. ip:%s", cli.name, cli.ip)
			//使用客户端的ip作为键
			clients[cli.ip] = cli
		case cli := <-leaving:
			delete(clients, cli.ip)
			close(cli.channel)
			_ = cli.conn.Close()
		}
	}
}
