package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/logic"
)

type MessageAPI struct {
}

func (api *MessageAPI) MessageAction(c *gin.Context) {
	/*
			query:
				token<string>:用户鉴权token
				to_user_id<string>:对方用户id
				action_type<string>:1-发送消息
				content<string>:消息内容
			response:
				 status_code<int64>:状态码，0-成功，其他值-失败
		    	 status_msg<string|null>:返回状态描述
		    }
	*/
	//使用中间件将token转化成user_id
	userId := c.GetInt64("user_id")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	content := c.Query("content")
	if actionType == "1" {
		response := logic.SendMessageAction(userId, toUserId, content)
		c.JSON(http.StatusOK, response)
	}
}

func (api *MessageAPI) MessageChat(c *gin.Context) {
	/*
				query:
					token<string>:用户鉴权token
					to_user_id<string>:对方用户id
				response:
					 status_code<int64>:状态码，0-成功，其他值-失败
			    	 status_msg<string|null>:返回状态描述
					 message_list<array[object (Message)|null>:用户列表
						id<int>:消息id
		            	content<string>:消息内容
		            	create_time<string>:消息发送时间 yyyy-MM-dd HH:MM:ss

	*/
	//使用中间件将token转化成user_id
	userId := c.GetInt64("user_id")
	toUserId := c.Query("to_user_id")
	response := logic.DoMessageChat(userId, toUserId)
	c.JSON(http.StatusOK, response)
}
