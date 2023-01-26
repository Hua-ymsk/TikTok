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
		return "", fmt.Errorf("token resolution error: %w", err)
	}
	sqlStatement := "SELECT * FROM users WHERE id=?;"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return "", fmt.Errorf("select init error: %w", err)
	}
	res, err := stmt.Query(userId)
	if err != nil {
		return "", fmt.Errorf("select execute error: %w", err)
	}
	defer res.Close()
	if res.Next() == false {
		return "", errors.New("user not exist error")
	}
	return Token, nil
}

// FollowExist 查询关注信息是否存在
func FollowExist(UserId string, ToUserId string) (exist bool, err error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return false, fmt.Errorf("parameter error: %w", err)
	}
	toUserId, err := strconv.Atoi(ToUserId)
	if err != nil {
		return false, fmt.Errorf("parameter error: %w", err)
	}
	sqlStatement := "SELECT * FROM follows WHERE following_user_id=? and followed_user_id=?;"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return false, fmt.Errorf("select init error: %w", err)
	}
	res, err := stmt.Query(userId, toUserId)
	if err != nil {
		return false, fmt.Errorf("select execute error: %w", err)
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
		return fmt.Errorf("parameter error: %w", err)
	}
	toUserId, err := strconv.Atoi(ToUserId)
	if err != nil {
		return fmt.Errorf("parameter error: %w", err)
	}
	sqlStatement := "INSERT INTO follows (`following_user_id`, `followed_user_id`, `relationship`) VALUES (?, ?, ?);"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return fmt.Errorf("insert init error: %w", err)
	}
	_, err = stmt.Exec(userId, toUserId, Relationship)
	if err != nil {
		return fmt.Errorf("insert execute error: %w", err)
	}
	return nil
}

// ChangeRelation 修改关注信息
func ChangeRelation(UserId string, ToUserId string, Relationship int) (err error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return fmt.Errorf("parameter error: %w", err)
	}
	toUserId, err := strconv.Atoi(ToUserId)
	if err != nil {
		return fmt.Errorf("parameter error: %w", err)
	}
	sqlStatement := "UPDATE follows set relationship=? WHERE following_user_id=? and followed_user_id=?;"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return fmt.Errorf("update init error: %w", err)
	}
	_, err = stmt.Exec(Relationship, userId, toUserId)
	if err != nil {
		return fmt.Errorf("update execute error: %w", err)
	}
	return nil
}

// DeleteFollow 删除关注信息
func DeleteFollow(UserId string, ToUserId string) (err error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return fmt.Errorf("parameter error: %w", err)
	}
	toUserId, err := strconv.Atoi(ToUserId)
	if err != nil {
		return fmt.Errorf("parameter error: %w", err)
	}
	sqlStatement := "DELETE FROM follows WHERE following_user_id=? and followed_user_id=?;"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return fmt.Errorf("delete init error: %w", err)
	}
	_, err = stmt.Exec(userId, toUserId)
	if err != nil {
		return fmt.Errorf("delete execute error: %w", err)
	}
	return nil
}

// SelectFollowList 查询所有关注的信息列表
func SelectFollowList(UserId string) (*sql.Rows, error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return nil, fmt.Errorf("parameter error: %w", err)
	}
	sqlStatement :=
		`SELECT
			users.id,
			users.user_name,
			fans,
			follows,
			1 is_followed
		FROM
			users
			LEFT JOIN follows f ON users.id = f.followed_user_id
			LEFT JOIN follows d ON users.id = f.following_user_id 
		WHERE
			f.following_user_id = ?;`
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("select init error: %w", err)
	}
	res, err := stmt.Query(userId)
	if err != nil {
		return nil, fmt.Errorf("select execute error: %w", err)
	}
	return res, nil
}

// SelectFollowerList 查询所有粉丝的信息列表
func SelectFollowerList(UserId string) (*sql.Rows, error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return nil, fmt.Errorf("parameter error: %w", err)
	}
	sqlStatement :=
		`SELECT
			users.id,
			users.user_name,
			fans,
			follows,
		IF
			(f.relationship= ?, 1, 0) AS is_followed 
		FROM
			users
			LEFT JOIN follows f ON users.id = f.following_user_id 
			LEFT JOIN follows d ON users.id = f.followed_user_id
		WHERE
			f.followed_user_id = ?;`
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("select init error: %w", err)
	}
	res, err := stmt.Query(userId, userId)
	if err != nil {
		return nil, fmt.Errorf("select execute error: %w", err)
	}
	return res, nil
}

// SelectFriendList 查询所有互关的信息列表
func SelectFriendList(UserId string) (*sql.Rows, error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return nil, fmt.Errorf("parameter error: %w", err)
	}
	sqlStatement :=
		`SELECT
			users.id,
			users.user_name,
			fans,
			follows,
			1 is_followed
		FROM
			users
			LEFT JOIN follows f ON users.id = f.followed_user_id
			LEFT JOIN follows d ON users.id = f.following_user_id 
		WHERE
			f.following_user_id = ? and f.relationship = 1;`
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("select init error: %w", err)
	}
	res, err := stmt.Query(userId)
	if err != nil {
		return nil, fmt.Errorf("select execute error: %w", err)
	}
	return res, nil
}
