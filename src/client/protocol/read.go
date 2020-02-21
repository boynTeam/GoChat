package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
)

// Author:Boyn
// Date:2020/2/21
const magicNumber = 0xABC123

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

func ResolveMessage(conn net.Conn, output chan Message) {
	var buf [65542]byte
	for {
		n, err := conn.Read(buf[0:])
		// 如果包格式不合格,就进行下次循环而不会继续进行处理
		if !isPackageValid(buf[0:n]) {
			continue
		}

		if err != nil && err != io.EOF {
			fmt.Println("传输数据错误:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(bytes.NewReader(buf[0:n]))
		scanner.Split(packetSlitFunc)
		result := bytes.NewBuffer(nil)
		for scanner.Scan() {
			result.Write(scanner.Bytes())
		}
		message, err := parseMessageFromJSON(result.Bytes()[6:])
		if err != nil {
			fmt.Println("解析错误:", err)
			continue
		}
		output <- message
	}
}

func parseMessageFromJSON(content []byte) (Message, error) {
	msg := Message{}
	err := json.Unmarshal(content, &msg)
	return msg, err
}

func isPackageValid(data []byte) bool {
	return len(data) > 6 && binary.BigEndian.Uint32(data[:4]) == magicNumber
}
