package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	. "server/internal"
	"time"
)

// Author:Boyn
// Date:2020/2/21

func packetSlitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 检查 atEOF 参数 和 数据包头部的四个字节是否 为 0x123456(我们定义的协议的魔数)
	if !atEOF && len(data) > 6 && binary.BigEndian.Uint32(data[:4]) == 0xABC123 {
		var l uint16
		// 读出 数据包中 实际数据 的长度(大小为 0 ~ 2^32)
		binary.Read(bytes.NewReader(data[4:6]), binary.BigEndian, &l)
		pl := int(l) + 6
		if pl <= len(data) {
			return pl, data[:pl], nil
		}
	}
	return
}

func AcceptEnter(conn net.Conn, Messages chan string, cli Client) {
	var buf [65542]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil && err != io.EOF {
			fmt.Println("传输数据错误:", err)
			break
		}
		scanner := bufio.NewScanner(result)
		scanner.Split(packetSlitFunc)
		for scanner.Scan() {
			cli.LastSend = time.Now()
			Messages <- cli.Ip + ":" + scanner.Text()[6:]
		}
	}
}
