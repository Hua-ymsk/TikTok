package dao

import (
	"database/sql"
	"errors"
	"fmt"
)

// 查询用户是否存在并返回其UserId,不存在或错误返回错误
func UserExist(Token string) (UserId string, err error) {
	return
}

// 查询对方关注信息是否存在
func FollowExist(UserId string, ToUserId string) (Exist bool, err error) {
	return
}

// 插入关注信息(0为未互关，1为已互关)
func InsertFollow(UserId string, ToUserId string, Relationship int) (err error) {
	return
}

// 修改关注信息
func ChangeRelation(UserId string, ToUserId string, Relationship int) (err error) {
	return
}

// 删除关注信息
func DeleteFollow(UserId string, ToUserId string) (err error) {
	return
}

// 查询所有关注的信息列表
func SelectFollowList(UserId string, Token string) (*sql.Rows, error) {
	fmt.Println(db)
	stmt, err := db.Prepare("select * FROM Follows where `followed_user_id`=?")
	if err != nil {
		return nil, errors.New("error:select init")
	}
	res, err := stmt.Query(UserId)
	fmt.Println(res)
	if err != nil {
		return nil, errors.New("error:select execute")
	}
	return res, nil

}

// 查询所有粉丝的信息列表
func SelectFollowerList(UserId string, Token string) (*sql.Rows, error) {
	fmt.Println(db)
	stmt, err := db.Prepare("select * FROM Follows where `followed_user_id`=?")
	if err != nil {
		return nil, errors.New("error:select init")
	}
	res, err := stmt.Query(UserId)
	fmt.Println(res)
	if err != nil {
		return nil, errors.New("error:select execute")
	}
	return res, nil

}

// 查询所有互关的信息列表
func SelectFriendList(UserId string, Token string) (*sql.Rows, error) {
	fmt.Println(db)
	stmt, err := db.Prepare("select * FROM Follows where `followed_user_id`=?")
	if err != nil {
		return nil, errors.New("error:select init")
	}
	res, err := stmt.Query(UserId)
	fmt.Println(res)
	if err != nil {
		return nil, errors.New("error:select execute")
	}
	return res, nil

}
