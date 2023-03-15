package manage_socket_conn

import (
	"Notes-go-project/model/socketModel"
	"Notes-go-project/service/message"
	"fmt"
	"sync"
)

var cRoomThread = new(charRoomThread)

type charRoomThread struct {
	msgChannel chan socketModel.ConnMsg
	lock       sync.Mutex
	once       sync.Once
}

//	向通道发送数据
func (c *charRoomThread) SendMsg(msg socketModel.ConnMsg) {
	fmt.Println(msg)
	if msg.Msg.ChatMsgType == 1 {
		if msg.Msg.Data["room_id"] == nil {
			return
		}
	} else if msg.Msg.ChatMsgType == 2 {
		fmt.Println(msg.Msg.Data["to_user_id"])
		if msg.Msg.Data["to_user_id"] == nil {
			return
		}
	}

	c.msgChannel <- msg
}

//	单例
func GetCharRoomThread() *charRoomThread {
	cRoomThread.once.Do(func() {
		cRoomThread.msgChannel = make(chan socketModel.ConnMsg, 30)
		cRoomThread.lock = sync.Mutex{}
	})
	return cRoomThread
}

//	启动通道监听
//	ChatMsgType 1 群聊信息  2 一对一信息
func (c *charRoomThread) Start() {
	for {
		select {
		case msg := <-c.msgChannel:
			//发送
			if msg.Msg.ChatMsgType == 1 {
				//	标明发送方用户id
				msg.Msg.Data["form_user_id"] = msg.FormUserID
				//	在这里你可以将聊天信息入库等等操作
				// 	do something
				//msg.Msg.Data["cs"] = "SB"
				//	发送信息
				//	注意 msg.Msg.Data["room_id"].(int) 这种写法在data为nil时 运行时会 panic 导致整个系统停掉
				//	所以在上一层最好对数据内容进行判断，再把值发送到通道内
				GetRoomSet().SendMsgToUserList(int(msg.Msg.Data["room_id"].(float64)), msg.Msg)
			} else if msg.Msg.ChatMsgType == 2 {
				//	在这里你可以将聊天信息入库等等操作
				// 	do something
				msg.Msg.Data["form_user_id"] = msg.FormUserID
				if msg.FormUserID == int(msg.Msg.Data["to_user_id"].(float64)) {
					msg.Msg.Data["content"] = "pong"
					msg.Msg.Data["type"] = "oneself"
					msg.Msg.Data["messageCount"] = message.GetMessageMarks(int(msg.Msg.Data["to_user_id"].(float64)))
				}
				//	如果发送不成功 说明接收方不在线
				_ = GetUserSet().SendMsgByUid(int(msg.Msg.Data["to_user_id"].(float64)), msg.Msg)
			}
		}
	}
}
