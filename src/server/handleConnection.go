package chat

import (
	"net"
	. "server/handlers"
	. "server/internal"
	"server/protocol"
)

/**
author:Boyn
date:2020/2/15
*/

func handleConn(conn net.Conn) {
	cli := NewClient(conn.RemoteAddr().String(), conn)
	go clientWriter(cli)
	HandleLogin(cli)
	protocol.AcceptEnter(conn, BroadCaster, *cli)
	Leaving <- *cli
	BroadCaster <- cli.Ip + " has left"
}

func clientWriter(cli *Client) {
	for msg := range cli.Channel {
		_ = protocol.WriteMessage(msg, cli.Conn)
	}
}
