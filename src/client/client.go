package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
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
	}
	fmt.Printf("输入信息:")
	scanner := bufio.NewScanner(os.Stdin)
	go AcceptEnter(conn)
	for scanner.Scan() {
		text := scanner.Text()
		err := writeMessage(text, conn)
		if err != nil {
			fmt.Println("发送错误:", err)
			os.Exit(-1)
		}
		fmt.Printf("输入信息:")
	}
}

func AcceptEnter(conn net.Conn) {
	io.Copy(os.Stdout, conn)
}

func writeMessage(content string, conn net.Conn) (err error) {
	buf := createPackage()
	length := make([]byte, 2)
	binary.BigEndian.PutUint16(length, uint16(len(content)))
	buf.Write(length)
	buf.Write([]byte(content))
	_, err = conn.Write(buf.Bytes())
	return
}

func connect(ip string) (net.Conn, error) {
	// 设置超时时间为30秒
	dialTimeout, err := net.DialTimeout("tcp", ip, timeout*time.Second)
	return dialTimeout, err
}

// 创建一个byte数组,其中已经放好包的部分头部(魔数部分)
func createPackage() (buf bytes.Buffer) {
	magicNum := make([]byte, 4)
	binary.BigEndian.PutUint32(magicNum, 0xABC123)
	buf.Write(magicNum)
	return
}
