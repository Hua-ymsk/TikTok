package mysql

import (
	"database/sql"
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
		return false, fmt.Errorf("like action error: %v", err)
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

// SelectLikeList 查询喜欢列表
func SelectLikeList(userId string) (*sql.Rows, error) {
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return nil, fmt.Errorf("string to int error:%v", err)
	}
	var likes = make([]*models.Like, 0)
	res := db.Where("user_id = ?", userIdInt).Find(&likes)
	//检查是否找到数据
	if len(likes) == 0 {
		return nil, fmt.Errorf("select likelist id null")
	}
	resRows, errRows := res.Rows()
	if errRows != nil {
		return nil, fmt.Errorf("select likelist row error:%v", errRows)
	}
	return resRows, nil
}
