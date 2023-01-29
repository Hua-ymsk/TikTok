package controller

import (
	"net/http"
	"strconv"
	"tiktok/common/result"
	"tiktok/logic"
	"tiktok/types"

	"github.com/gin-gonic/gin"
)

type VideoAPI struct{}

func (api *VideoAPI) FeedHandler(c *gin.Context) {

}

func (api *VideoAPI) PublishHandler(c *gin.Context) {
	// 接受参数
	var req types.PublishReq
	if err := c.ShouldBind(&req); err != nil {
		result.ResponseErr(c, "参数错误")
		return
	}
	// 保存文件
	l := logic.NewVideoLogic()
	if err := l.SaveVideo(c, req.Data, req.Title); err != nil {
		result.ResponseErr(c, "上传失败")
		return
	}
	result.ResponseSuccess(c)

	return
}

func (api *VideoAPI) PublishListHandler(c *gin.Context) {
	var req types.PublishListReq
	if err := c.ShouldBind(&req); err != nil {
		result.ResponseErr(c, "参数错误")
		return
	}
	l := logic.NewVideoLogic()
	user_id, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		result.ResponseErr(c, "参数错误")
		return
	}
	 videoList, err := l.VideoList(c, user_id)
	if err != nil {
		result.ResponseErr(c, "查询错误")
		return
	}
	c.JSON(http.StatusOK, types.PublishListResp{
		StatusCode: 0,
		StatusMsg: "请求成功",
		VideoList: videoList,
	})

	return
}
