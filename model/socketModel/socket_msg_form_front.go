package socketModel

type ConnMsg struct {
	Msg        ChatMsg `json:"msg,omitempty"`
	FormUserID int     `json:"form_user_id,omitempty"`
}

// ChatMsgType = 1 群聊信息 ChatMsgType = 2 一对一信息 ...
type ChatMsg struct {
	ChatMsgType int                    `json:"chat_msg_type,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
}
