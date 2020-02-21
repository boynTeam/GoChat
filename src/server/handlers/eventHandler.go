package handlers

import (
	"log"
	. "server/internal"
)

/**
事件循环处理器
使用select来循环事件
author:Boyn
date:2020/2/15
*/

func HandleEvent() {
	for {
		select {
		case msg := <-BroadCaster:
			log.Println(msg)
			Clients.Range(func(k, v interface{}) bool {
				cli := v.(Client)
				if cli.State == LoggedIn {
					cli.Channel <- msg
				}
				return true
			})
		case cli := <-Entering:
			log.Printf("%s login. ip:%s", cli.Name, cli.Ip)
			//使用客户端的ip作为键
			Clients.Store(cli.Ip, cli)
		case cli := <-Leaving:
			Clients.Delete(cli.Ip)
			close(cli.Channel)
			_ = cli.Conn.Close()
		}
	}
}
