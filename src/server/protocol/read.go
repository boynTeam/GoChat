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

const magicNumber = 0xABC123

// Author:Boyn
// Date:2020/2/21

func AcceptEnter(conn net.Conn, Messages chan string, cli Client) {
	for {
		message, err, isValid := ReadOneMessage(conn)
		if err != nil {
			fmt.Println("传输数据错误:", err)
			cli.State = Exited
			break
		}
		//读到无效包时直接将其丢弃
		if !isValid {
			continue
		}
		cli.LastSend = time.Now()
		Messages <- cli.Name + ":" + message
	}
}

// 只读取一条消息,并返回字符串与错误
func ReadOneMessage(conn net.Conn) (string, error, bool) {
	var buf [65542]byte
	n, err := conn.Read(buf[0:])
	if err != nil && err != io.EOF {
		return "", err, false
	}
	content, valid := readFromPackage(buf[0:n])
	return content, nil, valid
}

func packetSlitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 检查 atEOF 参数 和 数据包头部的四个字节是否 为 0x123456(我们定义的协议的魔数)
	if !atEOF && len(data) > 6 && binary.BigEndian.Uint32(data[:4]) == magicNumber {
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

func readFromPackage(buf []byte) (string, bool) {
	result := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(bytes.NewReader(buf))
	if !isPackageValid(buf) {
		return "", false
	}
	scanner.Split(packetSlitFunc)
	for scanner.Scan() {
		result.Write(scanner.Bytes())
	}
	return result.String()[6:], true
}

func isPackageValid(data []byte) bool {
	return len(data) > 6 && binary.BigEndian.Uint32(data[:4]) == magicNumber
}
