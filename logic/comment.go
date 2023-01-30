package logic

import (
	"fmt"
	"strconv"
	"tiktok/dao/mysql"
	"tiktok/types"
	"time"
)

// DoCommentAction 执行评论操作
func DoCommentAction(userId, videoId, commentText, commentId string) types.CommentActionResp {
	//评论操作
	timestamp := time.Now().Unix()
	errIn := mysql.InsertCommentInfo(userId, videoId, commentText, timestamp)
	if errIn != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("insert commentinfo error:%v", errIn),
		}
	}
	userIdInt, errUid := strconv.Atoi(userId)
	if errUid != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("string to int error:%v", errUid),
		}
	}
	commentIdInt, errCid := strconv.Atoi(commentId)
	if errCid != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("string to int error:%v", errCid),
		}
	}
	//查询评论用户的信息
	userName, followCount, followerCount, isFollow, errUserInfo := mysql.SelectUserInfo(userId)
	if errUserInfo != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("query userinfo error:%v", errUserInfo),
		}
	}
	//将时间戳转换为时间
	commentTime := time.Unix(timestamp, 0)
	//将时间格式化为mm:dd
	commentTimeMonthStr := commentTime.Month().String()
	//获取day的int类型再转换为字符串
	commentTimeDayInt := commentTime.Day()
	commentTimeDayStr := strconv.Itoa(commentTimeDayInt)
	//合并month和day
	commentTimeStr := commentTimeMonthStr + ":" + commentTimeDayStr
	return types.CommentActionResp{
		Comment: types.Comment{
			Content:    commentText,
			CreateDate: commentTimeStr,
			ID:         int64(commentIdInt),
			User: types.User{
				FollowCount:   followCount,
				FollowerCount: followerCount,
				UserID:        int64(userIdInt),
				IsFollow:      isFollow,
				Name:          userName,
			},
		},
		StatusCode: 0,
		StatusMsg:  "success",
	}
}

// DoUnCommentAction 执行删除评论操作
func DoUnCommentAction(commentId string) types.CommentActionResp {
	//首先查询需要删除评论的信息
	commentInfo, userInfo, errInfo := mysql.SelectDeleteCommentInfo(commentId)
	if errInfo != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("query conmmentinfo error:%v", errInfo),
		}
	}
	//执行删除操作
	errDelete := mysql.DeleteCommentInfo(commentId)
	if errDelete != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("query conmmentinfo error:%v", errDelete),
		}
	}
	//将时间戳转换为时间
	commentTime := time.Unix(commentInfo.Timestamp, 0)
	//将时间格式化为mm:dd
	commentTimeMonthStr := commentTime.Month().String()
	//获取day的int类型再转换为字符串
	commentTimeDayInt := commentTime.Day()
	commentTimeDayStr := strconv.Itoa(commentTimeDayInt)
	//合并month和day
	commentTimeStr := commentTimeMonthStr + ":" + commentTimeDayStr
	return types.CommentActionResp{
		Comment: types.Comment{
			Content:    commentInfo.Content,
			CreateDate: commentTimeStr,
			ID:         commentInfo.ID,
			User: types.User{
				FollowCount:   userInfo.Follows,
				FollowerCount: userInfo.Fans,
				UserID:        userInfo.ID,
				IsFollow:      userInfo.IsFollow,
				Name:          userInfo.UserName,
			},
		},
		StatusCode: 0,
		StatusMsg:  "success",
	}
}
