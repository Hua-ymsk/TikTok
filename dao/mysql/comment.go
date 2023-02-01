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
func SelectDeleteCommentInfo(commentId string) (commentInfo models.Comment, userInfo models.User, err error) {
	commentIdInt, errConv := strconv.Atoi(commentId)
	if errConv != nil {
		return models.Comment{}, models.User{}, fmt.Errorf("string to int error:%v", errConv)
	}
	resComment := db.Where("id = ?", commentIdInt).Find(&commentInfo)
	if resComment.RowsAffected == 0 {
		return models.Comment{}, models.User{}, fmt.Errorf("commentid no exist")
	}
	resUser := db.Where("id = ?", commentInfo.UserId).Find(&userInfo)
	if resUser.RowsAffected == 0 {
		return models.Comment{}, models.User{}, fmt.Errorf("user no exist")
	}
	return
}

// SelectUserInfo 通过用户id获取用户信息
func SelectUserInfo(userId int64) (userName string, followCount, followerCount int64, isFollow bool, err error) {
	var user models.User
	res := db.Where("id = ?", userId).Find(&user)
	//检查是否找到数据
	if res.RowsAffected == 0 {
		return "", 0, 0, false, fmt.Errorf("user no exist")
	}
	return user.UserName, user.Follows, user.Fans, user.IsFollow, nil
}

// SelectCommentList 查询评论列表
func SelectCommentList(videoId string) ([]*models.Comment, error) {
	videoIdInt, errVint := strconv.Atoi(videoId)
	if errVint != nil {
		return nil, fmt.Errorf("string to int error%v", errVint)
	}
	var comments = make([]*models.Comment, 0)
	res := db.Where("video_id = ?", videoIdInt).Find(&comments)
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("select comment error")
	}
	return comments, nil
}
