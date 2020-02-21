package handlers

import (
	"fmt"
	. "server/internal"
	"server/protocol"
	"time"
)

// Author:Boyn
// Date:2020/2/21

// 负责管理客户端的登录流程

func HandleLogin(cli *Client) {
	// 其他状态的处理 如果不是未登录的状态,那么就不会进行处理
	if cli.State != NotLoggedIn {
		return
	}
	cli.Channel <- Message{Content: "请登录", User: "Server", Time: time.Now().String(), State: NotLoggedIn}
	message, err, valid := protocol.ReadOneMessage(cli.Conn)
	if err != nil || !valid {
		fmt.Println("传输数据错误:", err)
		cli.State = Exited
		return
	}
	cli.State = LoggedIn
	cli.Name = message.Content
	cli.Channel <- Message{Content: message.Content, User: "Server", Time: time.Now().String(), State: LoginSuccess}
	BroadCaster <- Message{Content: cli.Name + " Has Arrived", User: "Server", Time: time.Now().String(), State: LoggedIn}
	Entering <- *cli
	return
}
