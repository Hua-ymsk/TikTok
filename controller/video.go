package controller

import (
	"net/http"
	"strconv"
	"tiktok/common/result"
	"tiktok/logic"
	"tiktok/types"
	"time"

	"github.com/gin-gonic/gin"
)

type VideoAPI struct{}

func (api *VideoAPI) FeedHandler(c *gin.Context) {
	var req types.FeedReq
	if err := c.ShouldBind(&req); err != nil {
		result.ResponseErr(c, "参数错误")
		return
	}
	// 判断latest_time
	var latest int64
	if req.LatestTime == "" {
		latest = time.Now().Unix()
	} else {
		tmp, _ := strconv.ParseInt(req.LatestTime, 10, 64)
		latest = time.Unix(tmp, 0).Unix()
	}
	l := logic.NewVideoLogic()
	feed, next_time, err := l.Feed(latest)
	if err != nil {
		result.ResponseErr(c, "获取视频流失败")
	}
	c.JSON(http.StatusOK, types.FeedResp{
		NextTime:   next_time,
		StatusCode: 0,
		StatusMsg:  "请求成功",
		VideoList:  feed,
	})

	return
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
