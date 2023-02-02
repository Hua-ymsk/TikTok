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
	var comment = make([]*models.Comment, 0)
	res := db.Where("id = ?", commentIdInt).Delete(&comment)
	if res.Error != nil {
		return fmt.Errorf("delete like error: %v", res.Error)
	}
	return nil
}

// SelectDeleteCommentInfo 查询删除评论的信息
func SelectDeleteCommentInfo(commentId string) (commentInfo models.Comment, userId int64, videoId int64, err error) {
	commentIdInt, errConv := strconv.Atoi(commentId)
	if errConv != nil {
		return models.Comment{}, 0, 0, fmt.Errorf("string to int error:%v", errConv)
	}
	resComment := db.Select("id", "user_id", "timestamp", "content").Where("id = ?", commentIdInt).Take(&commentInfo)
	if resComment.Error != nil {
		return models.Comment{}, 0, 0, fmt.Errorf("commentid error:%v", resComment.Error)
	}
	if resComment.RowsAffected == 0 {
		return models.Comment{}, 0, 0, fmt.Errorf("commentid no exist")
	}
	userId = commentInfo.UserId
	videoId = commentInfo.VideoId
	return
}

// SelectUserInfo 通过用户id获取用户信息
func SelectUserInfo(userId int64) (nickName string, followCount, followerCount int64, err error) {
	var user models.User
	res := db.Select("nickname", "follows", "fans").Where("id = ?", userId).Take(&user)
	if res.Error != nil {
		err = res.Error
		return
	}
	nickName = user.NickName
	followCount = user.Follows
	followerCount = user.Fans
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
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return comments, nil
}

// SelectVideoUserId 查询发布视频的用户id
func SelectVideoUserId(videoId string) (int64, error) {
	videoIdInt, err := strconv.Atoi(videoId)
	if err != nil {
		return 0, fmt.Errorf("string to int error:%v", err)
	}
	var video models.Video
	res := db.Select("user_id").Where("id = ?", videoIdInt).Take(&video)
	if res.Error != nil {
		return 0, fmt.Errorf("select video userinfo error:%v", res.Error)
	}
	return video.UserID, nil
}
