package protocol

import (
	"bytes"
	"encoding/binary"
	"net"
	. "server/internal"
)

// Author:Boyn
// Date:2020/2/21

// 控制协议内容

// 创建一个byte数组,其中已经放好包的部分头部(魔数部分)
func createPackage() (buf bytes.Buffer) {
	magicNum := make([]byte, 4)
	binary.BigEndian.PutUint32(magicNum, 0xABC123)
	buf.Write(magicNum)
	return
}

// 写入消息
func WriteMessage(content string, conn net.Conn) (err error) {
	msg := Message{Content: content}
	buf := createPackage()
	length := make([]byte, 2)
	msgJSON, _ := msg.ToJSON()
	binary.BigEndian.PutUint16(length, uint16(len(msgJSON)))
	buf.Write(length)
	buf.Write(msgJSON)
	_, err = conn.Write(buf.Bytes())
	return
}
