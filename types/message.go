package types

type MessageActionResp struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type MessageChatResp struct {
	MessageList []Message `json:"message_list"` // 用户列表
	StatusCode  string    `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string    `json:"status_msg"`   // 返回状态描述
}

type Message struct {
	Content    string `json:"content"`     // 消息内容
	CreateTime int64  `json:"create_time"` // 消息发送时间 yyyy-MM-dd HH:MM:ss
	//ID         int64  `json:"id"`           // 消息id
	ToUserId   int64 `json:"to_user_id"`   // 该消息接收者的id
	FormUserId int64 `json:"from_user_id"` // 该消息发送者的id
}
