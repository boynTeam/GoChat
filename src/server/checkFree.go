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
		clients.Range(func(key, value interface{}) bool {
			if time.Now().Minute()-value.(client).lastSend.Minute() > closeConnectionMinutes {
				leaving <- value.(client)
			}
			return true
		})
	}
}
