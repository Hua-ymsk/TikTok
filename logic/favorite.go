package logic

import (
	"fmt"
	"strconv"
	"tiktok/dao/mysql"
	"tiktok/types"
)

// DoLike 执行赞操作
func DoLike(userId int64, videoId string) types.FavoriteActionResp {
	//查询赞是否存在
	exist, err := mysql.LikeExist(userId, videoId)
	if err != nil {
		return types.FavoriteActionResp{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("select like info error:%v", err),
		}
	}
	//不存在则点赞，存在则报错
	if !exist {
		err = mysql.InsertLikeInfo(userId, videoId)
		if err != nil {
			return types.FavoriteActionResp{
				StatusCode: 1,
				StatusMsg:  fmt.Sprintf("like action error:%v", err),
			}
		}
		return types.FavoriteActionResp{
			StatusCode: 0,
			StatusMsg:  "like action",
		}
	}
	return types.FavoriteActionResp{
		StatusCode: 1,
		StatusMsg:  "repeat like action",
	}
}

// DoUnlike 执行取消点赞操作
func DoUnlike(userId int64, videoId string) types.FavoriteActionResp {
	//查询赞是否存在
	exist, err := mysql.LikeExist(userId, videoId)
	if err != nil {
		return types.FavoriteActionResp{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("select like info error:%v", err),
		}
	}
	//存在则取消点赞，不存在则报错
	if exist {
		err := mysql.DeleteLikeInfo(userId, videoId)
		if err != nil {
			return types.FavoriteActionResp{
				StatusCode: 1,
				StatusMsg:  fmt.Sprintf("unlike action error:%v", err),
			}
		}
		return types.FavoriteActionResp{
			StatusCode: 0,
			StatusMsg:  "unlike action",
		}
	}
	return types.FavoriteActionResp{
		StatusCode: 1,
		StatusMsg:  "like no exist",
	}
}

// DoSelectLikeList 查询喜欢列表
func DoSelectLikeList(userId string, userIdNow int64) types.FavoriteListResp {
	userIdInt, errUser := strconv.Atoi(userId)
	if errUser != nil {
		return types.FavoriteListResp{
			StatusCode: "1",
			StatusMsg:  fmt.Sprintf("select likelist error:%v", errUser),
			VideoList:  nil,
		}
	}
	//查询喜欢列表
	res, errRead := mysql.SelectLikeList(int64(userIdInt))
	if errRead != nil {
		return types.FavoriteListResp{
			StatusCode: "1",
			StatusMsg:  fmt.Sprintf("select likelist error:%v", errRead),
			VideoList:  nil,
		}
	}
	//将查询到的结果集读取到返回值中
	var likeList = make([]types.Video, 0, 100)
	for _, videoInfo := range res {
		var like types.Video
		authorId := videoInfo.UserID
		_, _, _, authorNickName, followerCount, followCount, isFollow, err := mysql.QueryUserID(authorId, userIdNow)
		if err != nil {
			return types.FavoriteListResp{
				StatusCode: "1",
				StatusMsg:  fmt.Sprintf("select authorinfo error:%v", err),
				VideoList:  nil,
			}
		}
		videoIdStr := strconv.Itoa(int(videoInfo.ID))
		like.ID = videoInfo.ID
<<<<<<< HEAD
		like.Author.ID = videoInfo.UserID
		like.Author.NickName = authorName
		like.Author.Follows = followCount
		like.Author.Fans = followerCount
=======
		like.Author.UserID = authorId
		like.Author.Name = authorNickName
		like.Author.FollowCount = followCount
		like.Author.FollowerCount = followerCount
>>>>>>> 7711de7b60d61edcb2bb53c775ce518ad14b4a94
		like.Author.IsFollow = isFollow
		like.PlayURL = videoInfo.PlayURL
		like.CoverURL = videoInfo.CoverURL
		like.FavoriteCount = videoInfo.FavoriteCount
		like.CommentCount = videoInfo.CommentCount
		like.IsFavorite, err = mysql.LikeExist(int64(userIdInt), videoIdStr)
		if err != nil {
			return types.FavoriteListResp{
				StatusCode: "1",
				StatusMsg:  fmt.Sprintf("select like exist error:%v", err),
				VideoList:  nil,
			}
		}
		like.Title = videoInfo.Title
		likeList = append(likeList, like)
	}
	return types.FavoriteListResp{
		StatusCode: "0",
		StatusMsg:  "success",
		VideoList:  likeList,
	}
}
