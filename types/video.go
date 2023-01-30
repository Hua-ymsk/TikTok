package types

import "mime/multipart"

// 视频流
type FeedReq struct {
	LatestTime string `json:"latest_time"`
	Token      string `json:"token"`
}
type FeedResp struct {
	NextTime   int64   `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 视频列表
}

// 发布投稿
type PublishReq struct {
	Data  *multipart.FileHeader `form:"data"`
	Token string                `form:"token"`
	Title string                `form:"title"`
}

// 发布列表
type PublishListReq struct {
	Token  string `form:"token"`
	UserID string `form:"user_id"` //
}
type PublishListResp struct {
	StatusCode int64   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	VideoList  []Video `json:"video_list"`
}

type Video struct {
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

type User struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	UserID        int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}
