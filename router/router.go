package router

import (
	"tiktok/controller"
	"tiktok/middleware"

	"github.com/gin-gonic/gin"
)

var (
	videoAPI    = &controller.VideoAPI{}
	relationAPI = &controller.RelationAPI{}
	favoriteAPI = &controller.FavoriteAPI{}
	commentAPI  = &controller.CommentAPI{}
	userAPI     = &controller.UserAPI{}
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
		video.Use(middleware.JWTAuth())
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
		// favorite apis
		favorite := apiRouter.Group("/favorite")
		favorite.Use(middleware.JWTAuth())
		{
			favorite.POST("/action", favoriteAPI.FavoriteAction)
			favorite.GET("/list", favoriteAPI.FavoriteList)
		}
		// comment apis
		comment := apiRouter.Group("/comment")
		comment.Use(middleware.JWTAuth())
		{
			comment.POST("/action", commentAPI.CommentAction)
			comment.GET("/list", commentAPI.CommentList)
		}
		//user
		user := apiRouter.Group("/user")
		{
			user.GET("", middleware.JWTAuth(), userAPI.UserInfo)
			user.POST("/register", userAPI.Register)
			user.POST("/login", userAPI.Login)
		}

	}

	return r
}
