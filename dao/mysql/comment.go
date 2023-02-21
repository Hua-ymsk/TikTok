package mysql

import (
	"fmt"
	"strconv"
	"tiktok/models"
)

// InsertCommentInfo 添加评论信息
func InsertCommentInfo(userId, timestamp int64, videoId, commentText string) (int64, error) {
	videoIdInt, errVideo := strconv.Atoi(videoId)
	if errVideo != nil {
		return 0, fmt.Errorf("string to int error:%v", errVideo)
	}
	commentInfo := &models.Comment{
		UserId:    userId,
		VideoId:   int64(videoIdInt),
		Timestamp: timestamp,
		Content:   commentText,
	}
	res := db.Create(commentInfo)
	if res.Error != nil {
		return 0, fmt.Errorf("insert commentinfo error:%v", res.Error)
	}
	return commentInfo.ID, nil
}

// DeleteCommentInfo 删除评论信息
func DeleteCommentInfo(commentId string) error {
	commentIdInt, err := strconv.Atoi(commentId)
	if err != nil {
		return fmt.Errorf("string to int error:%v", err)
	}
	resQuery := db.Where("id = ?", commentIdInt).Take(&models.Comment{})
	if resQuery.Error != nil {
		return fmt.Errorf("comment info no exist: %v", resQuery.Error)
	}
	res := db.Where("id = ?", commentIdInt).Delete(&models.Comment{})
	if res.Error != nil {
		return fmt.Errorf("delete like error: %v", res.Error)
	}
	return nil
}

// SelectUserInfo 通过用户id获取用户信息
func SelectUserInfo(userId int64) (user *models.User, favoriteCounts, workCounts, totalFavorite int64, err error) {
	res := db.Where("id = ?", userId).Take(&user)
	if res.Error != nil {
		err = res.Error
		return
	}
	//查询喜欢数量
	resFavoriteCount := db.Model(models.Like{}).Where("user_id = ?", userId).Count(&favoriteCounts)
	if resFavoriteCount.Error != nil {
		err = resFavoriteCount.Error
		return
	}
	//查询作品数量
	resWorkCounts := db.Model(models.Video{}).Where("user_id = ?", userId).Count(&workCounts)
	if resWorkCounts.Error != nil {
		err = resWorkCounts.Error
		return
	}
	//查询获赞数
	totalFavorite, err = TotalFavorite(userId)
	if err != nil {
		return
	}
	return
}

// SelectCommentList 查询评论列表
func SelectCommentList(videoId string) ([]*models.Comment, error) {
	videoIdInt, errVint := strconv.Atoi(videoId)
	if errVint != nil {
		return nil, fmt.Errorf("string to int error%v", errVint)
	}
	var comments = make([]*models.Comment, 0)
	res := db.Select("id", "user_id", "timestamp", "content").Where("video_id = ?", videoIdInt).Find(&comments)
	if res.Error != nil {
		return nil, fmt.Errorf("select comment list error:%v", res.Error)
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return comments, nil
}

// SelectVideoUserId 查询发布视频的用户id
func SelectVideoUserId(videoId int64) (int64, error) {
	var video models.Video
	res := db.Select("user_id").Where("id = ?", videoId).Take(&video)
	if res.Error != nil {
		return 0, fmt.Errorf("select video userinfo error:%v", res.Error)
	}
	return video.UserID, nil
}
