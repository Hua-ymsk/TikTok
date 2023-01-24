package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

// TokenResolution 解析Token获得userId
func TokenResolution(Token string) (userId string, err error) {
	//暂不验证Token
	userId, err = Token, nil
	//
	return
}

// UserExist 解析Token获得userId,并检验该Token对应用户是否存在
func UserExist(Token string) (userId string, err error) {
	userId, err = TokenResolution(Token)
	if err != nil {
		return "", errors.New("token resolution error")
	}
	sqlStatement := "SELECT * FROM Users WHERE id=?;"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return "", errors.New("select init error")
	}
	res, err := stmt.Query(userId)
	if err != nil {
		return "", errors.New("select execute error")
	}
	defer res.Close()
	if res.Next() == false {
		fmt.Println("user not exist")
		return "", errors.New("user not exist error")
	}
	fmt.Println("user exist")
	return Token, nil
}

// FollowExist 查询关注信息是否存在
func FollowExist(UserId string, ToUserId string) (exist bool, err error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return false, errors.New("parameter error")
	}
	toUserId, err := strconv.Atoi(ToUserId)
	if err != nil {
		return false, errors.New("parameter error")
	}
	sqlStatement := "SELECT * FROM Follows WHERE following_user_id=? and followed_user_id=?;"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return false, errors.New("select init error")
	}
	res, err := stmt.Query(userId, toUserId)
	if err != nil {
		return false, errors.New("select execute error")
	}
	defer res.Close()
	if res.Next() == false {
		return false, nil
	}
	return true, nil
}

// InsertFollow 插入关注信息(0为未互关，1为已互关)
func InsertFollow(UserId string, ToUserId string, Relationship int) (err error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return errors.New("parameter error")
	}
	toUserId, err := strconv.Atoi(ToUserId)
	if err != nil {
		return errors.New("parameter error")
	}
	sqlStatement := "INSERT INTO Follows (`following_user_id`, `followed_user_id`, `relationship`) VALUES (?, ?, ?);"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return errors.New("insert init error")
	}
	_, err = stmt.Exec(userId, toUserId, Relationship)
	if err != nil {
		return errors.New("insert execute error")
	}
	return nil
}

// ChangeRelation 修改关注信息
func ChangeRelation(UserId string, ToUserId string, Relationship int) (err error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return errors.New("parameter error")
	}
	toUserId, err := strconv.Atoi(ToUserId)
	if err != nil {
		return errors.New("parameter error")
	}
	sqlStatement := "UPDATE Follows set relationship=? WHERE following_user_id=? and followed_user_id=?;"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return errors.New("update init error")
	}
	_, err = stmt.Exec(Relationship, userId, toUserId)
	if err != nil {
		return errors.New("update execute error")
	}
	return nil
}

// DeleteFollow 删除关注信息
func DeleteFollow(UserId string, ToUserId string) (err error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return errors.New("parameter error")
	}
	toUserId, err := strconv.Atoi(ToUserId)
	if err != nil {
		return errors.New("parameter error")
	}
	sqlStatement := "DELETE FROM Follows WHERE following_user_id=? and followed_user_id=?;"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return errors.New("delete init error")
	}
	_, err = stmt.Exec(userId, toUserId)
	if err != nil {
		return errors.New("delete execute error")
	}
	return nil
}

// SelectFollowList 查询所有关注的信息列表
func SelectFollowList(UserId string) (*sql.Rows, error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return nil, errors.New("parameter error")
	}
	sqlStatement :=
		`SELECT
			Users.id,
			Users.user_name,
			fans,
			follows,
			1 is_followed
		FROM
			Users
			LEFT JOIN Follows f ON Users.id = f.followed_user_id
			LEFT JOIN Follows d ON Users.id = f.following_user_id 
		WHERE
			f.following_user_id = ?;`
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return nil, errors.New("select init error")
	}
	res, err := stmt.Query(userId)
	if err != nil {
		return nil, errors.New("select execute error")
	}
	return res, nil
}

// SelectFollowerList 查询所有粉丝的信息列表
func SelectFollowerList(UserId string) (*sql.Rows, error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return nil, errors.New("parameter error")
	}
	sqlStatement :=
		`SELECT
			Users.id,
			Users.user_name,
			fans,
			follows,
		IF
			(f.relationship= ?, 1, 0) AS is_followed 
		FROM
			Users
			LEFT JOIN Follows f ON Users.id = f.following_user_id 
			LEFT JOIN Follows d ON Users.id = f.followed_user_id
		WHERE
			f.followed_user_id = ?;`
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return nil, errors.New("select init error")
	}
	res, err := stmt.Query(userId, userId)
	if err != nil {
		return nil, errors.New("select execute error")
	}
	return res, nil
}

// SelectFriendList 查询所有互关的信息列表
func SelectFriendList(UserId string) (*sql.Rows, error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return nil, errors.New("parameter error")
	}
	sqlStatement :=
		`SELECT
			Users.id,
			Users.user_name,
			fans,
			follows,
			1 is_followed
		FROM
			Users
			LEFT JOIN Follows f ON Users.id = f.followed_user_id
			LEFT JOIN Follows d ON Users.id = f.following_user_id 
		WHERE
			f.following_user_id = ? and f.relationship = 1;`
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return nil, errors.New("select init error")
	}
	res, err := stmt.Query(userId)
	if err != nil {
		return nil, errors.New("select execute error")
	}
	return res, nil
}
