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
	isFriend, err := IsFriend(userId, int64(toUserIdInt))
	if err != nil {
		return nil, fmt.Errorf("select friend relationship error:%v", err)
	}
	if !isFriend {
		return nil, fmt.Errorf("frined no exist")
	}
	var chats = make([]*models.Chat, 0, 100)
	res := db.Where("send_user_id = ? AND receive_user_id = ? OR send_user_id = ? AND receive_user_id = ? ", userId, toUserIdInt, toUserIdInt, userId).Find(&chats)
	if res.Error != nil {
		return nil, fmt.Errorf("select messagechat error:%v", res.Error)
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return chats, nil
}

// IsFriend 是否为好友
func IsFriend(userId, toUserId int64) (bool, error) {
	var friend models.Follow
	res := db.Select("id").Where("following_user_id = ? AND followed_user_id = ? AND relationship = ?", userId, toUserId, 1).Find(&friend)
	if res.Error != nil {
		return false, fmt.Errorf("select freind relaotionship error:%v", res.Error)
	}
	if res.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}
