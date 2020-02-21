package main

import (
	"bufio"
	"client/protocol"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// Author:Boyn
// Date:2020/2/20

const timeout = 30

func main() {
	conn, err := connect("localhost:8009")
	if err != nil {
		fmt.Println("连接错误:", err)
		os.Exit(-1)
	}
	fmt.Printf("输入信息:")
	scanner := bufio.NewScanner(os.Stdin)
	go AcceptEnter(conn, os.Stdout)
	for scanner.Scan() {
		text := scanner.Text()
		err := protocol.WriteMessage(text, conn)
		if err != nil {
			fmt.Println("发送错误:", err)
			os.Exit(-1)
		}
		fmt.Printf("输入信息:")
	}
}

func AcceptEnter(conn net.Conn, output io.Writer) {
	message := make(chan string)
	go protocol.ResolveMessage(conn, message)
	for {
		select {
		case msg := <-message:
			_, _ = fmt.Fprintln(output, msg)
		}
	}
}

func connect(ip string) (net.Conn, error) {
	// 设置超时时间为30秒
	dialTimeout, err := net.DialTimeout("tcp", ip, timeout*time.Second)
	return dialTimeout, err
}
