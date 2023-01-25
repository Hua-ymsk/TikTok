package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/logic"
)

type RelationAPI struct {}

func (api *RelationAPI) RelationAction(c *gin.Context) {
	/*
		query:
			token<string>:用户鉴权token
			to_user_id<string>:对方id
			action_type<string>:1关注|2取消
		response:
			status_code<int>:0成功|1失败
			status_msg<string>:信息
	*/
	Token := c.Query("token")
	ToUserId := c.Query("to_user_id")
	ActionType := c.Query("action_type")
	logic := logic.NewRelationLogic()
	if ActionType == "1" {
		response := logic.DoFollow(Token, ToUserId)
		c.JSON(http.StatusOK, response)
	} else if ActionType == "2" {
		response := logic.DoUnFollow(Token, ToUserId)
		c.JSON(http.StatusOK, response)
	}
}

func (api *RelationAPI) FollowList(c *gin.Context) { //
	/*
		query:
			user_id<string>:用户id
			token<string>:用户鉴权token
		response:
			status_code<string>:0成功|1失败
			status_msg<string|null>:返回状态描述
			user_list<array[object (User) {5}]|null>:用户信息列表
				id<int>:用户id
				name<string>:用户名称
				follow_count<int>:关注数
				follower_count<int>:粉丝数
				is_follow<bool>:是否已关注
	*/
	UserId := c.Query("user_id")
	Token := c.Query("token")
	logic := logic.NewRelationLogic()
	response := logic.SelectFollowList(UserId, Token)
	c.JSON(http.StatusOK, response)
}

func (api *RelationAPI) FollowerList(c *gin.Context) { //
	/*
		query:
			user_id<string>:用户id
			token<string>:用户鉴权token
		response:
			status_code<string>:0成功|1失败
			status_msg<string|null>:返回状态描述
			user_list<array[object (User) {5}]|null>:用户信息列表
				id<int>:用户id
				name<string>:用户名称
				follow_count<int>:关注数
				follower_count<int>:粉丝数
				is_follow<bool>:是否已关注
	*/
	UserId := c.Query("user_id")
	Token := c.Query("token")
	logic := logic.NewRelationLogic()
	response := logic.SelectFollowerList(UserId, Token)
	c.JSON(http.StatusOK, response)
}

func (api *RelationAPI) FriendList(c *gin.Context) { //
	/*
		query:
			user_id<string>:用户id
			token<string>:用户鉴权token
		response:
			status_code<string>:0成功|1失败
			status_msg<string|null>:返回状态描述
			user_list<array[object (User) {5}]|null>:用户信息列表
				id<int>:用户id
				name<string>:用户名称
				follow_count<int>:关注数
				follower_count<int>:粉丝数
				is_follow<bool>:是否已关注
	*/
	UserId := c.Query("user_id")
	Token := c.Query("token")
	logic := logic.NewRelationLogic()
	response := logic.SelectFriendList(UserId, Token)
	c.JSON(http.StatusOK, response)
}
