package mysql

import (
	"fmt"
	"strconv"
	"tiktok/models"
)

// InsertCommentInfo 添加评论信息
func InsertCommentInfo(userId, videoId, commentText string, timestamp int64) error {
	userIdInt, errUser := strconv.Atoi(userId)
	if errUser != nil {
		return fmt.Errorf("string to int error:%v", errUser)
	}
	videoIdInt, errVideo := strconv.Atoi(videoId)
	if errVideo != nil {
		return fmt.Errorf("string to int error:%v", errUser)
	}
	commentInfo := &models.Comment{
		UserId:    int64(userIdInt),
		VideoId:   int64(videoIdInt),
		Timestamp: timestamp,
		Content:   commentText,
	}
	res := db.Create(commentInfo)
	if res.Error != nil {
		return fmt.Errorf("insert commentinfo error:%v", res.Error)
	}
	return nil
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
func SelectUserInfo(userId string) (userName string, followCount, followerCount int64, isFollow bool, err error) {
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return "", 0, 0, false, fmt.Errorf("string to int error:%v", err)
	}
	var user models.User
	res := db.Where("id = ?", userIdInt).Find(&user)
	//检查是否找到数据
	if res.RowsAffected == 0 {
		return "", 0, 0, false, fmt.Errorf("user no exist")
	}
	return user.UserName, user.Follows, user.Fans, user.IsFollow, nil

}
