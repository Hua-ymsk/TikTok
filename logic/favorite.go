package logic

import (
	"database/sql"
	"fmt"
	"strconv"
	"tiktok/dao/mysql"
	"tiktok/types"
)

// DoLike 执行赞操作
func DoLike(token, videoId string) types.FavoriteActionResp {
	//转换token为user_id并查询用户是否存在
	if userExist, userId, _, err := mysql.QueryUserName(token); err == nil && userExist {
		userIdStr := strconv.Itoa(int(userId))
		//查询赞是否存在
		exist, err := mysql.LikeExist(userIdStr, videoId)
		if err != nil {
			return types.FavoriteActionResp{
				StatusCode: 1,
				StatusMsg:  fmt.Sprintf("select like info error:%v", err),
			}
		}
		//存在则取消赞，不存在则点赞
		if exist {
			err := mysql.DeleteLikeInfo(userIdStr, videoId)
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
			err := mysql.InsertLikeInfo(userIdStr, videoId)
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
	return types.FavoriteActionResp{
		StatusCode: 1,
		StatusMsg:  "User doesn't exist",
	}
}

func DoSelectLikeList(userId, token string) types.FavoriteListResp {
	//验证token
	userExist, _, _, err := mysql.QueryUserName(token)
	if err != nil {
		return types.FavoriteListResp{
			StatusCode: "1",
			StatusMsg:  fmt.Sprintf("select token error:%v", err),
			VideoList:  nil,
		}
	}
	if !userExist {
		fmt.Printf("token nonexistent")
		return types.FavoriteListResp{
			StatusCode: "1",
			StatusMsg:  fmt.Sprintf("token nonexistent"),
			VideoList:  nil,
		}
	}
	//查询喜欢列表
	res, errRead := mysql.SelectLikeList(userId)
	if errRead != nil {
		return types.FavoriteListResp{
			StatusCode: "1",
			StatusMsg:  fmt.Sprintf("select likelist error:%v", err),
			VideoList:  nil,
		}
	}
	//关闭结果集
	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {

		}
	}(res)
	//将查询到的结果集读取到返回值中
	var likeList = make([]types.Video, 0, 100)
	for res.Next() {
		var like types.Video
		err := res.Scan(&like.ID,
			&like.Author.ID, &like.Author.Name, &like.Author.FollowCount, &like.Author.FollowerCount, &like.Author.IsFollow,
			&like.PlayURL, &like.CoverURL, &like.FavoriteCount, &like.CommentCount, &like.IsFavorite, &like.Title)
		if err != nil {
			return types.FavoriteListResp{
				StatusCode: "1",
				StatusMsg:  fmt.Sprintf("scan error:%v", err),
				VideoList:  nil,
			}
		}
		likeList = append(likeList, like)
	}
	return types.FavoriteListResp{
		StatusCode: "0",
		StatusMsg:  "success",
		VideoList:  likeList,
	}
}
