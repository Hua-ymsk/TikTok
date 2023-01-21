package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/controller"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	//r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
}
