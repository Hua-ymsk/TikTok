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
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	if actionType == "1" {
		response := logic.DoLike(token, videoId)
		c.JSON(http.StatusOK, response)
	} else if actionType == "2" {
		response := logic.DoLike(token, videoId)
		c.JSON(http.StatusOK, response)
	}
}

func (api *FavoriteAPI) FavoriteList(c *gin.Context) {

}
