package protocol

import "net"

// Author:Boyn
// Date:2020/2/21

type Client struct {
	Number string // 唯一标识号
	Name   string // 姓名
	State  int
	Conn   net.Conn
}

const (
	NotLoggedIn  = iota // 未登录
	LoginSuccess        // 登录成功
	LoggedIn            // 已经登录
	Exited              // 已经退出,或者登录失败也会转移到这个状态
	Registering         // 未注册
)
