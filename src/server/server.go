package chat

import (
	"fmt"
	"log"
	"net"
)

/**
author:Boyn
date:2020/2/15
*/

/*
主函数,用于打开TCP端口监听以及处理广播和连接
*/
func Serve(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	go checkFreeAndClose()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
