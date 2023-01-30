package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/logic"
)

type CommentAPI struct {
}

func (a CommentAPI) CommentAction(c *gin.Context) {
	/*
			query:
				token<string>:用户鉴权token
				video_id<string>:视频id
				action_type<string>:1-发布评论，2-删除评论
				comment_text<string>:用户填写的评论内容，在action_type=1的时候使用
				comment_id<string>:要删除的评论id，在action_type=2的时候使用
			response:
				 status_code<int64>:状态码，0-成功，其他值-失败
		    	 status_msg<string|null>:返回状态描述
		         comment<object(Comment){4}|null>:返回评论信息
		         	id<int64>:评论id
					user<object>:用户信息
				                id<int>:用户id
				                name<string>:用户名称
				                follow_count<int>:关注总数
				                follower_count<int>:粉丝总数
				                is_follow<bool>:true-已关注，false-未关注
				content<string>:评论内容
		        create_date<string>:评论发布日期，格式 mm-dd
		    }
	*/
	//使用中间件将token转化成user_id
	userId := c.GetString("user_id")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	commentId := c.Query("comment_id")
	if actionType == "1" {
		commentText := c.Query("comment_text")
		response := logic.DoCommentAction(userId, videoId, commentText, commentId)
		c.JSON(http.StatusOK, response)
	} else if actionType == "2" {
		response := logic.DoUnCommentAction(commentId)
		c.JSON(http.StatusOK, response)
	}

}

func (a CommentAPI) CommentList(c *gin.Context) {

}
