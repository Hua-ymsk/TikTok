package logic

import (
	"fmt"
	"gin-demo/basic/dao"
)

type FollowUser struct { //关注用户信息
	Id            int    `json:"id,omitempty"`             //用户id
	Name          string `json:"name,omitempty"`           //用户名称
	FollowCount   int    `json:"follow_count,omitempty"`   //关注数
	FollowerCount int    `json:"follower_count,omitempty"` //粉丝数
	IsFollow      bool   `json:"is_follow,omitempty"`      //是否已关注
}

type RelationListResponse struct { //关注用户信息
	StatusCode string     `json:"status_code,omitempty"` //0成功|1失败
	StatusMsg  string     `json:"status_msg,omitempty"`  //返回状态描述
	User       FollowUser `json:"user"`                  //用户信息
}

type RelationResponse struct { //关注用户信息
	StatusCode int    `json:"status_code,omitempty"` //0成功|1失败
	StatusMsg  string `json:"status_msg,omitempty"`  //返回状态描述
}

// 执行关注，先查询是否有对方关注信息存在，若存在修改标记并插入一条新信息；否则只插入一条新信息
func DoFollow(Token string, ToUserId string) (relationResponse RelationResponse) {
	//此处先转换Token并查询数据库观察用户是否存在
	if user_id, err := dao.UserExist(Token); err == nil {
		//查询关注信息
		exist, err := dao.FollowExist(user_id, ToUserId)
		if err != nil { //检验出现错误
			return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error,%s", err)}
		}
		if exist {
			//有对方信息存在，修改标记并进行插入(0为未互关，1为已互关)
			//修改标记
			e := dao.ChangeRelation(ToUserId, user_id, 1)
			if e != nil {
				return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error,%s", err)}
			}
			//插入信息
			e = dao.InsertFollow(user_id, ToUserId, 1)
			if e != nil {
				return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error,%s", err)}
			}
		} else {
			//无对方信息存在，插入新信息(0为未互关，1为已互关)
			e := dao.InsertFollow(user_id, ToUserId, 0)
			if e != nil {
				return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error,%s", err)}
			}
		}
		relationResponse = RelationResponse{StatusCode: 0, StatusMsg: "True"}
	} else {
		relationResponse = RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error,%s", err)}
	}
	return
}

// 执行取消关注，先查询是否有对方关注信息存在，若存在修改标记并删除自己的关注信息；否则只删除自己的关注信息
func DoUnFollow(Token string, ToUserId string) (relationResponse RelationResponse) {
	//此处先转换Token并查询数据库观察用户是否存在
	if user_id, err := dao.UserExist(Token); err == nil {
		//查询关注信息
		exist, err := dao.FollowExist(user_id, ToUserId)
		if err != nil { //检验出现错误
			return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error,%s", err)}
		}
		if exist {
			//有对方信息存在，修改标记(0为未互关，1为已互关)
			e := dao.ChangeRelation(ToUserId, user_id, 0)
			if e != nil {
				return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error,%s", err)}
			}
		}
		//不管信息是否存在，执行删除并进行删除
		e := dao.DeleteFollow(user_id, ToUserId)
		if e != nil {
			return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error,%s", err)}
		}
		relationResponse = RelationResponse{StatusCode: 0, StatusMsg: "True"}
	} else {
		relationResponse = RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error,%s", err)}
	}
	return

}

// 查询关注信息列表并格式化转换成JSON格式
func SelectFollowList(UserId string, Token string) (relationListResponse RelationListResponse) {
	result, err := dao.SelectFollowList(UserId, Token)
	//错误处理
	if err == nil {
		followUser := FollowUser{}
		relationListResponse = RelationListResponse{StatusCode: "1", StatusMsg: "请求失败", User: followUser}
		return
	}
	//格式化查询
	fmt.Println(result)
	followUser := FollowUser{}
	relationListResponse = RelationListResponse{StatusCode: "0", StatusMsg: "请求成功", User: followUser}
	return
}

// 查询粉丝信息列表并格式化转换成JSON格式
func SelectFollowerList(UserId string, Token string) (relationListResponse RelationListResponse) {
	result, err := dao.SelectFollowerList(UserId, Token)
	//错误处理
	if err == nil {
		followUser := FollowUser{}
		relationListResponse = RelationListResponse{StatusCode: "1", StatusMsg: "请求失败", User: followUser}
		return
	}
	//格式化查询
	fmt.Println(result)
	followUser := FollowUser{}
	relationListResponse = RelationListResponse{StatusCode: "0", StatusMsg: "请求成功", User: followUser}
	return
}

// 查询互关信息列表并格式化转换成JSON格式
func SelectFriendList(UserId string, Token string) (relationListResponse RelationListResponse) {
	result, err := dao.SelectFriendList(UserId, Token)
	//错误处理
	if err == nil {
		followUser := FollowUser{}
		relationListResponse = RelationListResponse{StatusCode: "1", StatusMsg: "请求失败", User: followUser}
		return
	}
	//格式化查询
	fmt.Println(result)
	followUser := FollowUser{}
	relationListResponse = RelationListResponse{StatusCode: "0", StatusMsg: "请求成功", User: followUser}
	return
}
