package logic

import (
	"fmt"
	"strconv"
	"tiktok/dao/mysql"
	"tiktok/types"
	"time"
)

// DoCommentAction 执行评论操作
func DoCommentAction(userId int64, videoId, commentText string) types.CommentActionResp {
	//评论操作
	timestamp := time.Now().Unix()
	commentId, errIn := mysql.InsertCommentInfo(userId, timestamp, videoId, commentText)
	if errIn != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("insert commentinfo error:%v", errIn),
		}
	}
	//查询视频用户信息
	videoUserId, errVideoUser := mysql.SelectVideoUserId(videoId)
	if errVideoUser != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("select video userinfo error:%v", errIn),
		}
	}
	//查询评论用户的信息
	//是否关注视频作者
	_, _, _, _, _, _, isFollow, errIsFollow := mysql.QueryUserID(videoUserId, userId)
	if errIsFollow != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("query isfollow error:%v", errIsFollow),
		}
	}
	//其他信息
	userNickName, followCount, followerCount, errUserInfo := mysql.SelectUserInfo(userId)
	if errUserInfo != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("query userinfo error:%v", errIsFollow),
		}
	}
	//将时间戳转换为时间
	commentTime := time.Unix(timestamp, 0)
	//将时间格式化为mm:dd
	commentTimeStr := fmt.Sprintf("%02d:%02d", commentTime.Month(), commentTime.Day())
	return types.CommentActionResp{
		Comment: types.Comment{
			Content:    commentText,
			CreateDate: commentTimeStr,
			ID:         commentId,
			User: types.User{
				FollowCount:   followCount,
				FollowerCount: followerCount,
				UserID:        userId,
				IsFollow:      isFollow,
				Name:          userNickName,
			},
		},
		StatusCode: 0,
		StatusMsg:  "success",
	}
}

// DoUnCommentAction 执行删除评论操作
func DoUnCommentAction(commentId string) types.CommentActionResp {
	//首先查询需要删除评论的信息
	commentInfo, userId, videoId, errInfo := mysql.SelectDeleteCommentInfo(commentId)
	if errInfo != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("query conmmentinfo error:%v", errInfo),
		}
	}
	videoIdInt := strconv.Itoa(int(videoId))
	//查询视频用户信息
	videoUserId, errVideoUser := mysql.SelectVideoUserId(videoIdInt)
	if errVideoUser != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("select video userinfo error:%v", errVideoUser),
		}
	}
	//查删除评论用户的信息
	//是否关注视频作者
	_, _, _, _, _, _, isFollow, errIsFollow := mysql.QueryUserID(videoUserId, userId)
	if errIsFollow != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("query isfollow error:%v", errIsFollow),
		}
	}
	//其他信息
	userNickName, followCount, followerCount, errUserInfo := mysql.SelectUserInfo(userId)
	if errUserInfo != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("query userinfo error:%v", errIsFollow),
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
	commentTimeStr := fmt.Sprintf("%02d:%02d", commentTime.Month(), commentTime.Day())
	return types.CommentActionResp{
		Comment: types.Comment{
			Content:    commentInfo.Content,
			CreateDate: commentTimeStr,
			ID:         commentInfo.ID,
			User: types.User{
				FollowCount:   followCount,
				FollowerCount: followerCount,
				UserID:        userId,
				IsFollow:      isFollow,
				Name:          userNickName,
			},
		},
		StatusCode: 0,
		StatusMsg:  "success",
	}
}

// DoCommentList 查询评论列表
func DoCommentList(userId int64, videoId string) types.CommentListResp {
	comments, err := mysql.SelectCommentList(videoId)
	if err != nil {
		return types.CommentListResp{
			CommentList: nil,
			StatusCode:  1,
			StatusMsg:   fmt.Sprintf("select comment list error:%v", err),
		}
	}
	var commentList = make([]types.Comment, 0, 100)
	for _, comment := range comments {
		var commentTemp types.Comment
		commentUserId := comment.UserId
		_, _, _, commentName, followerCount, followCount, isFollow, err := mysql.QueryUserID(commentUserId, userId)
		if err != nil {
			return types.CommentListResp{
				CommentList: nil,
				StatusCode:  1,
				StatusMsg:   fmt.Sprintf("select userinfo error:%v", err),
			}
		}
		//将时间戳转换为时间
		commentTime := time.Unix(comment.Timestamp, 0)
		//将时间格式化为mm:dd
		commentTimeStr := fmt.Sprintf("%02d:%02d", commentTime.Month(), commentTime.Day())
		commentTemp.ID = comment.ID
		commentTemp.User.UserID = comment.UserId
		commentTemp.User.Name = commentName
		commentTemp.User.FollowCount = followCount
		commentTemp.User.FollowerCount = followerCount
		commentTemp.User.IsFollow = isFollow
		commentTemp.Content = comment.Content
		commentTemp.CreateDate = commentTimeStr
		commentList = append(commentList, commentTemp)
	}
	return types.CommentListResp{
		CommentList: commentList,
		StatusCode:  0,
		StatusMsg:   "success",
	}
}

//// DoNoLoginCommentList 未登录的用户查看评论列表
//func DoNoLoginCommentList(videoId string) types.CommentListResp {
//	comments, err := mysql.SelectCommentList(videoId)
//	if err != nil {
//		return types.CommentListResp{
//			CommentList: nil,
//			StatusCode:  1,
//			StatusMsg:   fmt.Sprintf("select comment list error:%v", err),
//		}
//	}
//	var commentList = make([]types.Comment, 0, 100)
//	for _, comment := range comments {
//		var commentTemp types.Comment
//		commentUserId := comment.UserId
//		commentName, followCount, followerCount, errUserInfo := mysql.SelectUserInfo(commentUserId)
//		if errUserInfo != nil {
//			return types.CommentListResp{
//				CommentList: nil,
//				StatusCode:  1,
//				StatusMsg:   fmt.Sprintf("select userinfo error:%v", errUserInfo),
//			}
//		}
//		//将时间戳转换为时间
//		commentTime := time.Unix(comment.Timestamp, 0)
//		//将时间格式化为mm:dd
//		commentTimeStr := fmt.Sprintf("%02d:%02d", commentTime.Month(), commentTime.Day())
//		commentTemp.ID = comment.ID
//		commentTemp.User.UserID = comment.UserId
//		commentTemp.User.Name = commentName
//		commentTemp.User.FollowCount = followCount
//		commentTemp.User.FollowerCount = followerCount
//		commentTemp.User.IsFollow = false
//		commentTemp.Content = comment.Content
//		commentTemp.CreateDate = commentTimeStr
//		commentList = append(commentList, commentTemp)
//	}
//	return types.CommentListResp{
//		CommentList: commentList,
//		StatusCode:  0,
//		StatusMsg:   "success",
//	}
//}
