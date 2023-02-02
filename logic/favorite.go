package logic

import (
	"fmt"
	"strconv"
	"tiktok/dao/mysql"
	"tiktok/types"
)

// DoLike 执行赞操作
func DoLike(userId, videoId string) types.FavoriteActionResp {
	//查询赞是否存在
	exist, err := mysql.LikeExist(userId, videoId)
	if err != nil {
		return types.FavoriteActionResp{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("select like info error:%v", err),
		}
	}
	//存在则取消赞，不存在则点赞
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
	} else {
		err := mysql.InsertLikeInfo(userId, videoId)
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
}

func DoSelectLikeList(userId string) types.FavoriteListResp {
	//查询喜欢列表
	res, errRead := mysql.SelectLikeList(userId)
	if errRead != nil {
		return types.FavoriteListResp{
			StatusCode: "1",
			StatusMsg:  fmt.Sprintf("select likelist error:%v", errRead),
			VideoList:  nil,
		}
	}
	////关闭结果集
	//defer func(res *sql.Rows) {
	//	err := res.Close()
	//	if err != nil {
	//		return
	//	}
	//}(res)
	//将查询到的结果集读取到返回值中
	//for res.Next() {
	//	var like types.Video
	//	err := res.Scan(&like.ID,
	//		&like.Author.UserID, &like.Author.Name, &like.Author.FollowCount, &like.Author.FollowerCount, &like.Author.IsFollow,
	//		&like.PlayURL, &like.CoverURL, &like.FavoriteCount, &like.CommentCount, &like.IsFavorite, &like.Title)
	//	if err != nil {
	//		return types.FavoriteListResp{
	//			StatusCode: "1",
	//			StatusMsg:  fmt.Sprintf("scan error:%v", err),
	//			VideoList:  nil,
	//		}
	//	}
	//	likeList = append(likeList, like)
	//}
	//将查询到的结果集读取到返回值中
	var likeList = make([]types.Video, 0, 100)
	for _, videoInfo := range res {
		var like types.Video
		authorId := strconv.Itoa(int(videoInfo.UserID))
		authorName, followCount, followerCount, isFollow, err := mysql.SelectUserInfo(authorId)
		if err != nil {
			return types.FavoriteListResp{
				StatusCode: "1",
				StatusMsg:  fmt.Sprintf("select authorinfo error:%v", err),
				VideoList:  nil,
			}
		}
		videoIdStr := strconv.Itoa(int(videoInfo.ID))
		like.ID = videoInfo.ID
		like.Author.ID = videoInfo.UserID
		like.Author.NickName = authorName
		like.Author.Follows = followCount
		like.Author.Fans = followerCount
		like.Author.IsFollow = isFollow
		like.PlayURL = videoInfo.PlayURL
		like.CoverURL = videoInfo.CoverURL
		like.FavoriteCount = videoInfo.FavoriteCount
		like.CommentCount = videoInfo.CommentCount
		like.IsFavorite, err = mysql.LikeExist(userId, videoIdStr)
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
