package chat

import (
	. "server/internal"
	"time"
)

/**
检查客户端的状态是否正常
author:Boyn
date:2020/2/18
*/
const closeConnectionMinutes = 5

func checkValidAndClose() {
	// 对客户端进行遍历
	// 会关闭状态不正确的客户端
	// 状态不正确包括:处于Exited等待关闭的客户端,超过5分钟没有发消息的客户端
	for {
		Clients.Range(func(key, value interface{}) bool {
			if time.Now().Minute()-value.(Client).LastSend.Minute() > closeConnectionMinutes {
				Leaving <- value.(Client)
			}
			if value.(Client).State == Exited {
				Leaving <- value.(Client)
			}
			return true
		})
	}
}
