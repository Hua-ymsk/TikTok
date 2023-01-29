package router

import (
	"tiktok/controller"
	"tiktok/middleware"

	"github.com/gin-gonic/gin"
)

var (
	videoAPI    = &controller.VideoAPI{}
	relationAPI = &controller.RelationAPI{}
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 全局应用logger和recover中间件
	r.Use(middleware.GinLogger, middleware.GinRecover(true))
	apiRouter := r.Group("/douyin")
	{
		// video apis
		apiRouter.GET("/feed", videoAPI.FeedHandler)
		video := apiRouter.Group("/publish")
		video.GET("/list", videoAPI.PublishListHandler)
		video.Use(middleware.JWTAuth())
		{
			video.POST("/action", videoAPI.PublishHandler)

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
