package mysql

import (
	"fmt"
	"strconv"
	"tiktok/models"
)

// LikeExist 查询是否已经点赞
func LikeExist(userId, videoId string) (bool, error) {
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return false, fmt.Errorf("string to int error:%v", err)
	}
	videoIdInt, err := strconv.Atoi(videoId)
	if err != nil {
		return false, fmt.Errorf("string to int error:%v", err)
	}
	var like = make([]*models.Like, 0)
	res := db.Where("user_id = ? AND video_id = ?", userIdInt, videoIdInt).Find(&like)
	if res.Error != nil {
		return false, fmt.Errorf("likeaction error: %v", err)
	}
	if res.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

// InsertLikeInfo 添加点赞信息
func InsertLikeInfo(userId, videoId string) error {
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return fmt.Errorf("string to int error:%v", err)
	}
	videoIdInt, err := strconv.Atoi(videoId)
	if err != nil {
		return fmt.Errorf("string to int error:%v", err)
	}
	var likeInfo = &models.Like{
		UserId:  userIdInt,
		VideoId: videoIdInt,
	}
	res := db.Create(likeInfo)
	if res.Error != nil {
		return fmt.Errorf("insert like error: %v", err)
	}
	return nil
}

// DeleteLikeInfo 删除点赞信息
func DeleteLikeInfo(userId, videoId string) error {
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return fmt.Errorf("string to int error:%v", err)
	}
	videoIdInt, err := strconv.Atoi(videoId)
	if err != nil {
		return fmt.Errorf("string to int error:%v", err)
	}
	var like = make([]*models.Like, 0)
	res := db.Where("user_id = ? AND video_id = ?", userIdInt, videoIdInt).Delete(&like)
	if res.Error != nil {
		return fmt.Errorf("delete like error: %v", err)
	}
	return nil
}
