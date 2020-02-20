package chat

import (
	"log"
	. "server/internal"
)

/**
author:Boyn
date:2020/2/15
*/

func broadcaster() {
	for {
		select {
		case msg := <-Messages:
			log.Println(msg)
			Clients.Range(func(k, v interface{}) bool {
				v.(Client).Channel <- msg
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
