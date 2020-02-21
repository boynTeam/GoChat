package chat

import (
	"fmt"
	"log"
	"net"
	"server/handlers"
	. "server/internal"
)

/**
author:Boyn
date:2020/2/15
*/

/*
主函数,用于打开TCP端口监听以及处理广播和连接
*/
func Serve(port int) {
	defer Stop()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	go handlers.HandleEvent()
	go checkValidAndClose()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// 在退出的时候清理资源
func Stop() {
	// 将所有已经连接的客户端关闭
	Clients.Range(func(key, value interface{}) bool {
		_ = value.(Client).Conn.Close()
		return true
	})
}
