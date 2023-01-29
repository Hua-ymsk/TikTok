package logic

import (
	"database/sql"
	"errors"
	"fmt"
	"tiktok/dao"
)

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
	User       []FollowUser `json:"user"`        //用户信息
}

type RelationResponse struct { //关注用户信息
	StatusCode int    `json:"status_code"` //0成功|1失败
	StatusMsg  string `json:"status_msg"`  //返回状态描述
}

// DoFollow 执行关注，先查询是否有对方关注信息存在，若存在修改标记并插入一条新信息；否则只插入一条新信息
func DoFollow(UserId string, ToUserId string) (relationResponse RelationResponse) {
	//查询关注信息是否存在
	exist, err := dao.FollowExist(UserId, ToUserId)
	if err != nil { //检验出现错误
		return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", err)}
	}
	if exist {
		return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", errors.New("follow exist"))}
	}
	//查询对方关注信息是否存在
	exist, err = dao.FollowExist(ToUserId, UserId)
	if err != nil { //检验出现错误
		return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", err)}
	}
	if exist {
		//有对方信息存在，修改标记并进行插入(0为未互关，1为已互关)
		e := dao.UpdInsRelation(UserId, ToUserId)
		if e != nil {
			return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", e)}
		}
	} else {
		//无对方信息存在，插入新信息(0为未互关，1为已互关)
		e := dao.InsertFollow(UserId, ToUserId, 0)
		if e != nil {
			return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", e)}
		}
	}
	relationResponse = RelationResponse{StatusCode: 0, StatusMsg: "True"}
	return
}

// DoUnFollow 执行取消关注，先查询是否有对方关注信息存在，若存在修改标记并删除自己的关注信息；否则只删除自己的关注信息
func DoUnFollow(UserId string, ToUserId string) (relationResponse RelationResponse) {
	//查询关注信息是否存在
	exist, err := dao.FollowExist(UserId, ToUserId)
	if err != nil { //检验出现错误
		return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", err)}
	}
	if !exist {
		return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", errors.New("follow not exist"))}
	}
	//查询对方关注信息是否存在
	exist, err = dao.FollowExist(ToUserId, UserId)
	if err != nil { //检验出现错误
		return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", err)}
	}
	if exist {
		//有对方信息存在，修改标记(0为未互关，1为已互关)
		e := dao.UpdDelRelation(ToUserId, UserId)
		if e != nil {
			return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", e)}
		}
	} else {
		//无对方信息存在，执行删除
		e := dao.DeleteFollow(UserId, ToUserId)
		if e != nil {
			return RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", e)}
		}
	}
	relationResponse = RelationResponse{StatusCode: 0, StatusMsg: "True"}
	return
}

// SelectFollowList 查询关注信息列表并格式化转换成JSON格式
func SelectFollowList(UserId string) (followListResponse RelationListResponse) {
	var followUserList = make([]FollowUser, 0, 100)
	//验证用户存在
	err := dao.UserExist(UserId)
	if err != nil {
		followListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: followUserList}
		return
	}
	//执行查询
	rows, err := dao.SelectFollowList(UserId)
	//错误处理
	if err != nil {
		followListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: followUserList}
		return
	}
	//格式化查询
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			followListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: followUserList}
			return
		}
	}(rows)
	//循环读取结果集中的数据
	for rows.Next() {
		var u FollowUser
		err := rows.Scan(&u.Id, &u.Name, &u.FollowCount, &u.FollowerCount, &u.IsFollow)
		if err != nil {
			followListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: followUserList}
			return
		}
		followUserList = append(followUserList, u)
	}
	followListResponse = RelationListResponse{StatusCode: "0", StatusMsg: "True", User: followUserList}
	return
}

// SelectFollowerList 查询粉丝信息列表并格式化转换成JSON格式
func SelectFollowerList(UserId string) (followerListResponse RelationListResponse) {
	var followerUserList = make([]FollowUser, 0, 100)
	//验证用户存在
	err := dao.UserExist(UserId)
	if err != nil {
		followerListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: followerUserList}
		return
	}
	//执行查询
	rows, err := dao.SelectFollowerList(UserId)
	//错误处理
	if err != nil {
		followerListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: followerUserList}
		return
	}
	//格式化查询
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			followerListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: followerUserList}
			return
		}
	}(rows)
	//循环读取结果集中的数据
	for rows.Next() {
		var u FollowUser
		err := rows.Scan(&u.Id, &u.Name, &u.FollowCount, &u.FollowerCount, &u.IsFollow)
		if err != nil {
			followerListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: followerUserList}
			return
		}
		followerUserList = append(followerUserList, u)
	}
	followerListResponse = RelationListResponse{StatusCode: "0", StatusMsg: "True", User: followerUserList}
	return
}

// SelectFriendList 查询互关信息列表并格式化转换成JSON格式
func SelectFriendList(UserId string) (friendListResponse RelationListResponse) {
	var friendUserList = make([]FollowUser, 0, 100)
	//验证用户存在
	err := dao.UserExist(UserId)
	if err != nil {
		friendListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: friendUserList}
		return
	}
	//执行查询
	rows, err := dao.SelectFriendList(UserId)
	//错误处理
	if err != nil {
		friendListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: friendUserList}
		return
	}
	//格式化查询
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			friendListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: friendUserList}
			return
		}
	}(rows)
	//循环读取结果集中的数据
	for rows.Next() {
		var u FollowUser
		err := rows.Scan(&u.Id, &u.Name, &u.FollowCount, &u.FollowerCount, &u.IsFollow)
		if err != nil {
			friendListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: friendUserList}
			return
		}
		friendUserList = append(friendUserList, u)
	}
	friendListResponse = RelationListResponse{StatusCode: "0", StatusMsg: "True", User: friendUserList}
	return
}
