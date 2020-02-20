package chat

import (
	"log"
	"sync"
)

/**
author:Boyn
date:2020/2/15
*/

var (
	entering = make(chan client) // 监控客户端进入的消息
	leaving  = make(chan client) // 监控客户端离开的消息
	messages = make(chan string) // 掌握所有客户端发出的消息
	clients  sync.Map            //掌握所有客户端消息
)

func broadcaster() {
	for {
		select {
		case msg := <-messages:
			log.Println(msg)
			clients.Range(func(k, v interface{}) bool {
				v.(client).channel <- msg
				return true
			})
		case cli := <-entering:
			log.Printf("%s login. ip:%s", cli.name, cli.ip)
			//使用客户端的ip作为键
			clients.Store(cli.ip, cli)
		case cli := <-leaving:
			clients.Delete(cli.ip)
			close(cli.channel)
			_ = cli.conn.Close()
		}
	}
}
