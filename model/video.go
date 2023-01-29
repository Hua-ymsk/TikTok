package model

// 投稿请求
type PublishReq struct {
	Token string `json:"token"`
	Data  []byte `json:"data"`
	Title string `json:"title"`
}

// 投稿响应
type PunblishResp struct {
	StausCcode uint64 `json:"staus_code"`
	StausMsg   string `json:"status_msg,optional"`
}

// 视频流请求
type FeedReq struct {
	LastTime int64  `json:"last_time"`
	Token    string `json:"token,optional"`
}

type Video struct {
}
