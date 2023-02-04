package logic

import (
	"fmt"
	"tiktok/dao/mysql"
	"tiktok/types"
	"time"
)

// SendMessageAction 执行发送消息
func SendMessageAction(userId int64, toUserId, content string) types.MessageActionResp {
	//发送消息
	err := mysql.InsertMessage(userId, toUserId, content)
	if err != nil {
		return types.MessageActionResp{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("send message error:%v", err),
		}
	}
	return types.MessageActionResp{
		StatusCode: 0,
		StatusMsg:  "success",
	}
}

// DoMessageChat 执行查看聊天记录
func DoMessageChat(userId int64, toUserId string) types.MessageChatResp {
	//查看聊天记录
	chats, err := mysql.SelectMessageChat(userId, toUserId)
	if err != nil {
		return types.MessageChatResp{
			MessageList: nil,
			StatusCode:  "1",
			StatusMsg:   fmt.Sprintf("select message chat error:%v", err),
		}
	}
	var res = make([]types.Message, 0, 100)
	for _, chat := range chats {
		//消息发送时间 yyyy-MM-dd HH:MM:ss
		messageTime := time.Unix(chat.Timestamp, 0)
		messageTimeStr := messageTime.Format("2006-01-02 15:04:05")
		temp := types.Message{
			Content:    chat.Content,
			CreateTime: messageTimeStr,
			ID:         chat.ID,
		}
		res = append(res, temp)
	}
	return types.MessageChatResp{
		MessageList: res,
		StatusCode:  "0",
		StatusMsg:   "success",
	}

}
