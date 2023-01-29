package logic

import (
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
