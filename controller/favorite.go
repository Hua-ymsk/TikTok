package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/logic"
)

type FavoriteAPI struct {
}

func (api *FavoriteAPI) FavoriteAction(c *gin.Context) {
	/*
		query:
			token<string>:用户鉴权token
			video_id<string>:视频id
			action_type<string>:1-点赞，2-取消点赞
		response:
			status_code<int>:0-成功，其他值-失败
			status_msg<string>:返回状态描述
	*/
	//使用中间件将token转化成user_id
	userId := c.GetString("user_id")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	if actionType == "1" {
		response := logic.DoLike(userId, videoId)
		c.JSON(http.StatusOK, response)
	} else if actionType == "2" {
		response := logic.DoLike(userId, videoId)
		c.JSON(http.StatusOK, response)
	}
}

func (api *FavoriteAPI) FavoriteList(c *gin.Context) {
	/*
			query:
				user_id<string>:用户id
				token<string>:用户鉴权token
			response:
				status_code<string>:状态码，0-成功，其他值-失败
		    	status_msg<string|null>:返回状态描述
		    	video_list<array[object (Video) {8}|null>:用户点赞视频列表
		        	id<int>:视频唯一标识
		            author<object>:视频作者信息
		                id<int>:用户id
		                name<string>:用户名称
		                follow_count<int>:关注总数
		                follower_count<int>:粉丝总数
		                is_follow<bool>:true-已关注，false-未关注
		            play_url<string>:视频播放地址
		            cover_url<string>:视频封面地址
		            favorite_count<int>:视频的点赞总数
		            comment_count<int>:视频的评论总数
		            is_favorite<bool>:true-已点赞，false-未点赞
		            title<string>:视频标题
	*/
	userId := c.Query("user_id")
	response := logic.DoSelectLikeList(userId)
	c.JSON(http.StatusOK, response)
}
