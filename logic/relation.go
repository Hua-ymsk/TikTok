package logic

import (
	"database/sql"
	"errors"
	"fmt"
	"tiktok/dao"
	"tiktok/types"
)

// DoFollow 执行关注，先查询是否有对方关注信息存在，若存在修改标记并插入一条新信息；否则只插入一条新信息
func DoFollow(UserId int64, ToUserId int64) (relationResponse types.RelationResponse) {
	//查询关注信息是否存在
	exist, err := dao.FollowExist(UserId, ToUserId)
	if err != nil { //检验出现错误
		return types.RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", err)}
	}
	if exist {
		return types.RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", errors.New("follow exist"))}
	}
	//查询对方关注信息是否存在
	exist, err = dao.FollowExist(ToUserId, UserId)
	if err != nil { //检验出现错误
		return types.RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", err)}
	}
	if exist {
		//有对方信息存在，修改标记并进行插入(0为未互关，1为已互关)
		e := dao.UpdInsRelation(UserId, ToUserId)
		if e != nil {
			return types.RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", e)}
		}
	} else {
		//无对方信息存在，插入新信息(0为未互关，1为已互关)
		e := dao.InsertFollow(UserId, ToUserId, 0)
		if e != nil {
			return types.RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", e)}
		}
	}
	relationResponse = types.RelationResponse{StatusCode: 0, StatusMsg: "True"}
	return
}

// DoUnFollow 执行取消关注，先查询是否有对方关注信息存在，若存在修改标记并删除自己的关注信息；否则只删除自己的关注信息
func DoUnFollow(UserId int64, ToUserId int64) (relationResponse types.RelationResponse) {
	//查询关注信息是否存在
	exist, err := dao.FollowExist(UserId, ToUserId)
	if err != nil { //检验出现错误
		return types.RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", err)}
	}
	if !exist {
		return types.RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", errors.New("follow not exist"))}
	}
	//查询对方关注信息是否存在
	exist, err = dao.FollowExist(ToUserId, UserId)
	if err != nil { //检验出现错误
		return types.RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", err)}
	}
	if exist {
		//有对方信息存在，修改标记(0为未互关，1为已互关)
		e := dao.UpdDelRelation(UserId, ToUserId)
		if e != nil {
			return types.RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", e)}
		}
	} else {
		//无对方信息存在，执行删除
		e := dao.DeleteFollow(UserId, ToUserId)
		if e != nil {
			return types.RelationResponse{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s", e)}
		}
	}
	relationResponse = types.RelationResponse{StatusCode: 0, StatusMsg: "True"}
	return
}

// SelectFollowList 查询关注信息列表并格式化转换成JSON格式
func SelectFollowList(UserId int64) (followListResponse types.RelationListResponse) {
	var followUserList = make([]types.FollowUser, 0, 100)
	//执行查询
	rows, err := dao.SelectFollowList(UserId)
	//错误处理
	if err != nil {
		followListResponse = types.RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: followUserList}
		return
	}
	//格式化查询
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			followListResponse = types.RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: followUserList}
			return
		}
	}(rows)
	//循环读取结果集中的数据
	for rows.Next() {
		var u types.FollowUser
		err := rows.Scan(&u.Id, &u.Name, &u.FollowCount, &u.FollowerCount, &u.IsFollow)
		if err != nil {
			followListResponse = types.RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: followUserList}
			return
		}
		followUserList = append(followUserList, u)
	}
	followListResponse = types.RelationListResponse{StatusCode: "0", StatusMsg: "True", User: followUserList}
	return
}

// SelectFollowerList 查询粉丝信息列表并格式化转换成JSON格式
func SelectFollowerList(UserId int64) (followerListResponse types.RelationListResponse) {
	var followerUserList = make([]types.FollowUser, 0, 100)
	//执行查询
	rows, err := dao.SelectFollowerList(UserId)
	//错误处理
	if err != nil {
		followerListResponse = types.RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: followerUserList}
		return
	}
	//格式化查询
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			followerListResponse = types.RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: followerUserList}
			return
		}
	}(rows)
	//循环读取结果集中的数据
	for rows.Next() {
		var u types.FollowUser
		err := rows.Scan(&u.Id, &u.Name, &u.FollowCount, &u.FollowerCount, &u.IsFollow)
		if err != nil {
			followerListResponse = types.RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: followerUserList}
			return
		}
		followerUserList = append(followerUserList, u)
	}
	followerListResponse = types.RelationListResponse{StatusCode: "0", StatusMsg: "True", User: followerUserList}
	return
}

// SelectFriendList 查询互关信息列表并格式化转换成JSON格式
func SelectFriendList(UserId int64) (friendListResponse types.FriendListResponse) {
	var friendUserList = make([]types.FriendUser, 0, 100)
	//执行查询
	rows, err := dao.SelectFriendList(UserId)
	//错误处理
	if err != nil {
		friendListResponse = types.FriendListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: friendUserList}
		return
	}
	//格式化查询
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			friendListResponse = types.FriendListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: friendUserList}
			return
		}
	}(rows)
	//循环读取结果集中的数据
	for rows.Next() {
		var u types.FriendUser
		err := rows.Scan(&u.Id, &u.Name, &u.FollowCount, &u.FollowerCount, &u.Avatar, &u.IsFollow)
		if err != nil {
			friendListResponse = types.FriendListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s", err), User: friendUserList}
			return
		}
		friendUserList = append(friendUserList, u)
	}
	friendListResponse = types.FriendListResponse{StatusCode: "0", StatusMsg: "True", User: friendUserList}
	return
}
