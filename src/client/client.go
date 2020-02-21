package main

import (
	"bufio"
	. "client/protocol"
	"fmt"
	"net"
	"os"
	"time"
)

// Author:Boyn
// Date:2020/2/20

// 本地客户端
var cli = Client{State: NotLoggedIn}

const timeout = 30

func main() {
	conn, err := connect("localhost:8009")
	if err != nil {
		fmt.Println("连接错误:", err)
		os.Exit(-1)
	}
	cli.Conn = conn
	scanner := bufio.NewScanner(os.Stdin)
	go AcceptEnter(conn)
	for scanner.Scan() {
		text := scanner.Text()
		err := WriteMessage(text, cli)
		if err != nil {
			fmt.Println("发送错误:", err)
			os.Exit(-1)
		}
	}
}

func AcceptEnter(conn net.Conn) {
	message := make(chan Message)
	go ResolveMessage(conn, message)
	for {
		select {
		case msg := <-message:
			handleMessage(msg)
		}
	}
}

func handleMessage(msg Message) {
	if cli.State == NotLoggedIn && msg.State == LoginSuccess {
		cli.State = LoggedIn
		cli.Name = msg.Content
		fmt.Println("登录成功,欢迎回来.", msg.Content)
	} else if cli.State == NotLoggedIn {
		fmt.Println("未登录:请输入账号")
	} else if cli.State == LoggedIn {
		fmt.Printf("%s %s\n%s\n", msg.User, msg.Time, msg.Content)
	}
}

func connect(ip string) (net.Conn, error) {
	// 设置超时时间为30秒
	dialTimeout, err := net.DialTimeout("tcp", ip, timeout*time.Second)
	return dialTimeout, err
}
