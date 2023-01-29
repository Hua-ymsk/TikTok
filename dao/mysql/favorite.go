package mysql

import (
	"fmt"
	"tiktok/model"
)

// LikeExist 查询是否已经点赞
func LikeExist(userId int, videoId int) (bool bool, err error) {
	var like = make([]*model.Like, 0)
	res := db.Where("user_id = ? AND video_id = ?", userId, videoId).Find(&like)
	if res.Error != nil {
		return false, fmt.Errorf("likeaction error: %v", err)
	}
	if res.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

// InsertLikeList 添加点赞信息
func InsertLikeList(userId int, videoId int) (bool bool, err error) {
	var likeInfo = &model.Like{
		UserId:  userId,
		VideoId: videoId,
	}
	res := db.Create(likeInfo)
	if res.Error != nil {
		return false, fmt.Errorf("insertlike error: %v", err)
	}
	return true, nil
}
