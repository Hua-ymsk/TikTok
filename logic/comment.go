package logic

import (
	"fmt"
	"github.com/jinzhu/copier"
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
	videoIdInt, errConv := strconv.Atoi(videoId)
	if errConv != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("string to int error:%v", errConv),
		}
	}
	//查询视频用户信息
	videoUserId, errVideoUser := mysql.SelectVideoUserId(int64(videoIdInt))
	if errVideoUser != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("select video userinfo error:%v", errIn),
		}
	}
	//查询评论用户的信息
	//是否关注视频作者
	_, _, _, isFollow, errIsFollow := mysql.QueryUserID(videoUserId, userId)
	if errIsFollow != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("query isfollow error:%v", errIsFollow),
		}
	}
	//其他信息
	commentUserInfo, favoriteCounts, workCounts, totalFavorite, errUserInfo := mysql.SelectUserInfo(userId)
	if errUserInfo != nil {
		return types.CommentActionResp{
			Comment:    types.Comment{},
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("query userinfo error:%v", errIsFollow),
		}
	}
	//将时间戳转换为时间
	commentTime := time.Unix(timestamp, 0)
	//将时间格式化为mm-dd
	commentTimeStr := commentTime.Format("01-02")
	return types.CommentActionResp{
		Comment: types.Comment{
			Content:    commentText,
			CreateDate: commentTimeStr,
			ID:         commentId,
			User: types.User{
				ID:             commentUserInfo.ID,
				NickName:       commentUserInfo.NickName,
				Follows:        commentUserInfo.Follows,
				Fans:           commentUserInfo.Fans,
				IsFollow:       isFollow,
				AvatarUrl:      commentUserInfo.AvatarUrl,
				BackGround:     commentUserInfo.BackGround,
				Signature:      commentUserInfo.Signature,
				TotalFavorited: totalFavorite,
				WorkCount:      workCounts,
				FavoriteCount:  favoriteCounts,
			},
		},
		StatusCode: 0,
		StatusMsg:  "success",
	}
}

// DoUnCommentAction 执行删除评论操作
func DoUnCommentAction(commentId string) types.CommentDeleteActionResp {
	//执行删除操作
	errDelete := mysql.DeleteCommentInfo(commentId)
	if errDelete != nil {
		return types.CommentDeleteActionResp{
			Comment:    nil,
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("query conmmentinfo error:%v", errDelete),
		}
	}
	return types.CommentDeleteActionResp{
		Comment:    nil,
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
		workCounts, favoriteCounts, commentUserInfo, isFollow, err := mysql.QueryUserID(commentUserId, userId)
		if err != nil {
			return types.CommentListResp{
				CommentList: nil,
				StatusCode:  1,
				StatusMsg:   fmt.Sprintf("select userinfo error:%v", err),
			}
		}
		totalFavorite, errTotal := mysql.TotalFavorite(commentUserId)
		if errTotal != nil {
			return types.CommentListResp{
				CommentList: nil,
				StatusCode:  1,
				StatusMsg:   fmt.Sprintf("select totalfavorite error:%v", errTotal),
			}
		}
		//将时间戳转换为时间
		commentTime := time.Unix(comment.Timestamp, 0)
		//将时间格式化为mm-dd
		commentTimeStr := commentTime.Format("01-02")
		errCopy := copier.Copy(&commentTemp, comment)
		if errCopy != nil {
			return types.CommentListResp{
				CommentList: nil,
				StatusCode:  1,
				StatusMsg:   fmt.Sprintf("copy commentinfo error:%v", err),
			}
		}
		errCopy = copier.Copy(&commentTemp.User, commentUserInfo)
		if errCopy != nil {
			return types.CommentListResp{
				CommentList: nil,
				StatusCode:  1,
				StatusMsg:   fmt.Sprintf("copy commentuserinfo error:%v", err),
			}
		}
		commentTemp.CreateDate = commentTimeStr
		commentTemp.User.IsFollow = isFollow
		commentTemp.User.TotalFavorited = totalFavorite
		commentTemp.User.WorkCount = workCounts
		commentTemp.User.FavoriteCount = favoriteCounts
		commentList = append(commentList, commentTemp)
	}
	return types.CommentListResp{
		CommentList: commentList,
		StatusCode:  0,
		StatusMsg:   "success",
	}
}

// DoNoLoginCommentList 未登录的用户查看评论列表
func DoNoLoginCommentList(videoId string) types.CommentListResp {
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
		commentUserInfo, favoriteCounts, workCounts, totalFavorite, errUserInfo := mysql.SelectUserInfo(commentUserId)
		if errUserInfo != nil {
			return types.CommentListResp{
				CommentList: nil,
				StatusCode:  1,
				StatusMsg:   fmt.Sprintf("select userinfo error:%v", errUserInfo),
			}
		}
		//将时间戳转换为时间
		commentTime := time.Unix(comment.Timestamp, 0)
		//将时间格式化为mm-dd
		commentTimeStr := commentTime.Format("01-02")
		errCopy := copier.Copy(&commentTemp, comment)
		if errCopy != nil {
			return types.CommentListResp{
				CommentList: nil,
				StatusCode:  1,
				StatusMsg:   fmt.Sprintf("copy commentinfo error:%v", err),
			}
		}
		errCopy = copier.Copy(&commentTemp.User, commentUserInfo)
		if errCopy != nil {
			return types.CommentListResp{
				CommentList: nil,
				StatusCode:  1,
				StatusMsg:   fmt.Sprintf("copy commentuserinfo error:%v", err),
			}
		}
		commentTemp.CreateDate = commentTimeStr
		commentTemp.User.IsFollow = false
		commentTemp.User.TotalFavorited = totalFavorite
		commentTemp.User.WorkCount = workCounts
		commentTemp.User.FavoriteCount = favoriteCounts
		commentList = append(commentList, commentTemp)
	}
	return types.CommentListResp{
		CommentList: commentList,
		StatusCode:  0,
		StatusMsg:   "success",
	}
}
