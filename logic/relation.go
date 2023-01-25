package logic

import (
	"database/sql"
	"errors"
	"fmt"
	"tiktok/dao"
)

type RelationLogic struct{}

type FollowUser struct { //关注用户信息
	Id            int    `json:"id"`             //用户id
	Name          string `json:"name"`           //用户名称
	FollowCount   int    `json:"follow_count"`   //关注数
	FollowerCount int    `json:"follower_count"` //粉丝数
	IsFollow      bool   `json:"is_follow"`      //是否已关注
}

type RelationListResp struct { //关注用户信息
	StatusCode string       `json:"status_code"` //0成功|1失败
	StatusMsg  string       `json:"status_msg"`  //返回状态描述
	User       []FollowUser `json:"user"`        //用户信息
}

type RelationResp struct { //关注用户信息
	StatusCode int    `json:"status_code"` //0成功|1失败
	StatusMsg  string `json:"status_msg"`  //返回状态描述
}

func NewRelationLogic() *RelationLogic {
	return &RelationLogic{}
}

// DoFollow 执行关注，先查询是否有对方关注信息存在，若存在修改标记并插入一条新信息；否则只插入一条新信息
func (logic *RelationLogic) DoFollow(Token string, ToUserId string) (relationResponse RelationResp) {
	//此处先转换Token并查询数据库观察用户是否存在
	if userId, err := dao.UserExist(Token); err == nil {
		//------此处测试使用，使Token=UserId即可使用------
		userId = Token
		//-----此处测试使用，使Token=UserId即可使用------
		//查询关注信息是否存在
		exist, err := dao.FollowExist(userId, ToUserId)
		if err != nil { //检验出现错误
			return RelationResp{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s ", err)}
		}
		if exist {
			return RelationResp{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s ", errors.New("follow exist"))}
		}
		//查询对方关注信息是否存在
		exist, err = dao.FollowExist(ToUserId, userId)
		if err != nil { //检验出现错误
			return RelationResp{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s ", err)}
		}
		if exist {
			//有对方信息存在，修改标记并进行插入(0为未互关，1为已互关)
			//修改标记
			e := dao.ChangeRelation(ToUserId, userId, 1)
			if e != nil {
				return RelationResp{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s ", err)}
			}
			//插入信息
			e = dao.InsertFollow(userId, ToUserId, 1)
			if e != nil {
				return RelationResp{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s ", err)}
			}
		} else {
			//无对方信息存在，插入新信息(0为未互关，1为已互关)
			e := dao.InsertFollow(userId, ToUserId, 0)
			if e != nil {
				return RelationResp{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s ", err)}
			}
		}
		relationResponse = RelationResp{StatusCode: 0, StatusMsg: "True"}
	} else {
		relationResponse = RelationResp{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s ", err)}
	}
	return
}

// DoUnFollow 执行取消关注，先查询是否有对方关注信息存在，若存在修改标记并删除自己的关注信息；否则只删除自己的关注信息
func (logic *RelationLogic) DoUnFollow(Token string, ToUserId string) (relationResponse RelationResp) {
	//此处先转换Token并查询数据库观察用户是否存在
	if userId, err := dao.UserExist(Token); err == nil {
		//------此处测试使用，使Token=UserId即可使用------
		userId = Token
		//-----此处测试使用，使Token=UserId即可使用------
		//查询关注信息是否存在
		exist, err := dao.FollowExist(userId, ToUserId)
		if err != nil { //检验出现错误
			return RelationResp{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s ", err)}
		}
		if !exist {
			return RelationResp{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s ", errors.New("follow not exist"))}
		}
		//查询对方关注信息是否存在
		exist, err = dao.FollowExist(ToUserId, userId)
		if err != nil { //检验出现错误
			return RelationResp{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s ", err)}
		}
		if exist {
			//有对方信息存在，修改标记(0为未互关，1为已互关)
			e := dao.ChangeRelation(ToUserId, userId, 0)
			if e != nil {
				return RelationResp{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s ", err)}
			}
		}
		//不管信息是否存在，执行删除并进行删除
		e := dao.DeleteFollow(userId, ToUserId)
		if e != nil {
			return RelationResp{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s ", err)}
		}
		relationResponse = RelationResp{StatusCode: 0, StatusMsg: "True"}
	} else {
		relationResponse = RelationResp{StatusCode: 1, StatusMsg: fmt.Sprintf("error, %s ", err)}
	}
	return
}

// SelectFollowList 查询关注信息列表并格式化转换成JSON格式
func (logic *RelationLogic) SelectFollowList(UserId string, Token string) (followListResponse RelationListResp) {
	var followUserList = make([]FollowUser, 0, 100)
	//验证Token
	//userId, err := dao.UserExist(Token)
	//if err != nil {
	//	followListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s ", err), User: followUserList}
	//	return
	//}
	//if userId != UserId {
	//	followListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s ", errors.New("Token authentication failed.")), User: followUserList}
	//	return
	//}
	//执行查询
	rows, err := dao.SelectFollowList(UserId)
	//错误处理
	if err != nil {
		followListResponse = RelationListResp{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s ", err), User: followUserList}
		return
	}
	//格式化查询
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			followListResponse = RelationListResp{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s ", err), User: followUserList}
			return
		}
	}(rows)
	//循环读取结果集中的数据
	for rows.Next() {
		var u FollowUser
		err := rows.Scan(&u.Id, &u.Name, &u.FollowCount, &u.FollowerCount, &u.IsFollow)
		if err != nil {
			followListResponse = RelationListResp{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s ", err), User: followUserList}
			return
		}
		followUserList = append(followUserList, u)
	}
	followListResponse = RelationListResp{StatusCode: "0", StatusMsg: "True", User: followUserList}
	return
}

// SelectFollowerList 查询粉丝信息列表并格式化转换成JSON格式
func (logic *RelationLogic) SelectFollowerList(UserId string, Token string) (followerListResponse RelationListResp) {
	var followerUserList = make([]FollowUser, 0, 100)
	//验证Token
	//userId, err := dao.UserExist(Token)
	//if err != nil {
	//	followerListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s ", err), User: followerUserList}
	//	return
	//}
	//if userId != UserId {
	//	followerListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s ", errors.New("Token authentication failed.")), User: followerUserList}
	//	return
	//}
	//执行查询
	rows, err := dao.SelectFollowerList(UserId)
	//错误处理
	if err != nil {
		followerListResponse = RelationListResp{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s ", err), User: followerUserList}
		return
	}
	//格式化查询
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			followerListResponse = RelationListResp{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s ", err), User: followerUserList}
			return
		}
	}(rows)
	//循环读取结果集中的数据
	for rows.Next() {
		var u FollowUser
		err := rows.Scan(&u.Id, &u.Name, &u.FollowCount, &u.FollowerCount, &u.IsFollow)
		if err != nil {
			followerListResponse = RelationListResp{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s ", err), User: followerUserList}
			return
		}
		followerUserList = append(followerUserList, u)
	}
	followerListResponse = RelationListResp{StatusCode: "0", StatusMsg: "True", User: followerUserList}
	return
}

// SelectFriendList 查询互关信息列表并格式化转换成JSON格式
func (logic *RelationLogic) SelectFriendList(UserId string, Token string) (friendListResponse RelationListResp) {
	var friendUserList = make([]FollowUser, 0, 100)
	//验证Token
	//userId, err := dao.UserExist(Token)
	//if err != nil {
	//	friendListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s ", err), User: friendUserList}
	//	return
	//}
	//if userId != UserId {
	//	friendListResponse = RelationListResponse{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s ", errors.New("Token authentication failed.")), User: friendUserList}
	//	return
	//}
	//执行查询
	rows, err := dao.SelectFriendList(UserId)
	//错误处理
	if err != nil {
		friendListResponse = RelationListResp{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s ", err), User: friendUserList}
		return
	}
	//格式化查询
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			friendListResponse = RelationListResp{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s ", err), User: friendUserList}
			return
		}
	}(rows)
	//循环读取结果集中的数据
	for rows.Next() {
		var u FollowUser
		err := rows.Scan(&u.Id, &u.Name, &u.FollowCount, &u.FollowerCount, &u.IsFollow)
		if err != nil {
			friendListResponse = RelationListResp{StatusCode: "1", StatusMsg: fmt.Sprintf("error, %s ", err), User: friendUserList}
			return
		}
		friendUserList = append(friendUserList, u)
	}
	friendListResponse = RelationListResp{StatusCode: "0", StatusMsg: "True", User: friendUserList}
	return
}
