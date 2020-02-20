package chat

import (
	. "server/internal"
	"time"
)

/**
检查空闲客户端
author:Boyn
date:2020/2/18
*/
const closeConnectionMinutes = 5

func checkFreeAndClose() {
	//对客户端进行遍历
	for {
		Clients.Range(func(key, value interface{}) bool {
			if time.Now().Minute()-value.(Client).LastSend.Minute() > closeConnectionMinutes {
				Leaving <- value.(Client)
			}
			return true
		})
	}
}
