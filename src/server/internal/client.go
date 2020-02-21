package internal

import (
	"net"
	"time"
)

// Author:Boyn
// Date:2020/2/20

/**
定义客户端的结构
author:Boyn
date:2020/2/18
*/
const clientChannelBufferSize = 10

type Client struct {
	Ip       string      //客户端的ip
	Name     string      //客户端在登录的时候指定的名i在
	Channel  chan string //用于收发客户端的消息
	Conn     net.Conn    // 客户端的TCP连接
	LastSend time.Time
}

func NewClient(name string, conn net.Conn) (cli *Client) {
	return &Client{
		Ip:       conn.RemoteAddr().String(),
		Name:     name,
		Channel:  make(chan string, clientChannelBufferSize),
		Conn:     conn,
		LastSend: time.Now(),
	}
}
