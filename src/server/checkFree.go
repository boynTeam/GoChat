package chat

import "time"

/**
检查空闲客户端
author:Boyn
date:2020/2/18
*/
const closeConnectionMinutes = 5

func checkFreeAndClose() {
	//对客户端进行遍历
	for {
		for _, cli := range clients {
			//过了5分钟还没有发消息
			if time.Now().Minute()-cli.lastSend.Minute() > closeConnectionMinutes {
				//当这个连接关闭的时候,会使input.Scan()也返回false,由客户端处理函数负责断开连接
				_ = cli.conn.Close()
			}
		}
	}
}
