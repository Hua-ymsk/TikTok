package router

import (
	"tiktok/controller"

	"github.com/gin-gonic/gin"
)

var (
	videoAPI = &controller.VideoAPI{}
	relationAPI = &controller.RelationAPI{}
)

func InitRouter() *gin.Engine {
	r := gin.New()
	apiRouter := r.Group("/douyin")
	{
		// video apis
		apiRouter.GET("/feed", videoAPI.FeedHandler)
		video := apiRouter.Group("/pulish")
		{
			video.POST("/action", videoAPI.PublishHandler)
			video.GET("/list", videoAPI.PublishListHandler)
		}
		// extra apis - II
		relation := apiRouter.Group("/relation")
		{
			relation.POST("/action", relationAPI.RelationAction)
			relation.GET("/follow/list", relationAPI.FollowList)
			relation.GET("/follower/list", relationAPI.FollowerList)
			relation.GET("/friend/list", relationAPI.FriendList)
		}
	}
	
	return r
}
