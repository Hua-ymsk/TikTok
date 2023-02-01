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
		UserId:  int64(userIdInt),
		VideoId: int64(videoIdInt),
	}
	res := db.Create(likeInfo)
	if res.Error != nil {
		return fmt.Errorf("insert like error: %v", err)
	}
	return nil
}

// DeleteLikeInfo 删除点赞信息
func DeleteLikeInfo(userId, videoId string) error {
	userIdInt, errUser := strconv.Atoi(userId)
	if errUser != nil {
		return fmt.Errorf("string to int error:%v", errUser)
	}
	videoIdInt, errVideo := strconv.Atoi(videoId)
	if errVideo != nil {
		return fmt.Errorf("string to int error:%v", errVideo)
	}
	var like = make([]*models.Like, 0)
	res := db.Where("user_id = ? AND video_id = ?", userIdInt, videoIdInt).Delete(&like)
	if res.Error != nil {
		return fmt.Errorf("delete like error: %v", res.Error)
	}
	return nil
}

// SelectLikeList 查询喜欢列表
func SelectLikeList(userId string) ([]*models.Video, error) {
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return nil, fmt.Errorf("string to int error:%v", err)
	}
	var likes = make([]*models.Like, 0)
	resLike := db.Where("user_id = ?", userIdInt).Find(&likes)
	//检查是否找到数据
	if resLike.RowsAffected == 0 {
		return nil, fmt.Errorf("query likelist is null")
	}
	//将video_id存进切片，为获取video信息
	var videoIds = make([]int64, 0, 100)
	for _, like := range likes {
		videoIds = append(videoIds, like.VideoId)
	}
	//获取video信息
	var videos = make([]*models.Video, 0)
	resVideo := db.Where("id IN ?", videoIds).Find(&videos)
	if resVideo.RowsAffected == 0 {
		return nil, fmt.Errorf("query videoinfo is null")
	}
	return videos, nil
}
