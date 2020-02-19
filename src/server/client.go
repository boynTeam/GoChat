package chat

import (
	"net"
	"time"
)

/**
定义客户端的结构
author:Boyn
date:2020/2/18
*/

type client struct {
	ip       string      //客户端的ip
	name     string      //客户端在登录的时候指定的名i在
	channel  chan string //用于收发客户端的消息
	conn     net.Conn    // 客户端的TCP连接
	lastSend time.Time
}
