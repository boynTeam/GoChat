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
	Ip       string       //客户端的ip
	Number   string       //客户端的标识码
	Name     string       //客户端在登录的时候指定的名i在
	Channel  chan Message //用于收发客户端的消息
	Conn     net.Conn     // 客户端的TCP连接
	State    int          // 客户端状态
	LastSend time.Time
}

const (
	NotLoggedIn = iota // 未登录
	LoggedIn           // 已经登录
	Exited             // 已经退出,或者登录失败也会转移到这个状态
	Registering        // 未注册
)

func NewClient(name string, conn net.Conn) (cli *Client) {
	return &Client{
		Ip:       conn.RemoteAddr().String(),
		Number:   conn.RemoteAddr().String(),
		Name:     name,
		Channel:  make(chan Message, clientChannelBufferSize),
		State:    NotLoggedIn,
		Conn:     conn,
		LastSend: time.Now(),
	}
}
