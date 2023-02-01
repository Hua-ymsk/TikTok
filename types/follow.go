package types

type FollowUser struct { //关注用户信息
	Id            int    `json:"id"`             //用户id
	Name          string `json:"name"`           //用户名称
	FollowCount   int    `json:"follow_count"`   //关注数
	FollowerCount int    `json:"follower_count"` //粉丝数
	IsFollow      bool   `json:"is_follow"`      //是否已关注
}

type RelationListResponse struct { //关注用户信息
	StatusCode string       `json:"status_code"` //0成功|1失败
	StatusMsg  string       `json:"status_msg"`  //返回状态描述
	User       []FollowUser `json:"user_list"`   //用户信息
}

type RelationResponse struct { //关注用户信息
	StatusCode int    `json:"status_code"` //0成功|1失败
	StatusMsg  string `json:"status_msg"`  //返回状态描述
}
