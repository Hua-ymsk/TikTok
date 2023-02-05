package mysql

import (
	"fmt"
	"strconv"
	"tiktok/models"
	"time"
)

// InsertMessage 添加消息
func InsertMessage(userId int64, toUserId, content string) error {
	toUserIdInt, errConv := strconv.Atoi(toUserId)
	if errConv != nil {
		return fmt.Errorf("string to ine error:%v", errConv)
	}
	timestamp := time.Now().Unix()
	message := &models.Chat{
		SendUserId:    userId,
		ReceiveUserId: int64(toUserIdInt),
		Timestamp:     timestamp,
		Content:       content,
	}
	res := db.Create(message)
	if res.Error != nil {
		return fmt.Errorf("insert message error:%v", res.Error)
	}
	return nil
}

// SelectMessageChat 查询聊天记录
func SelectMessageChat(userId int64, toUserId string) ([]*models.Chat, error) {
	toUserIdInt, errConv := strconv.Atoi(toUserId)
	if errConv != nil {
		return nil, fmt.Errorf("string to int error:%v", errConv)
	}
	var chats = make([]*models.Chat, 0, 100)
	res := db.Where("send_user_id = ? AND receive_user_id = ? OR send_user_id = ? AND receive_user_id = ? ", userId, toUserIdInt, toUserIdInt, userId).Find(&chats)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("select message chat error")
	}
	return chats, nil
}
