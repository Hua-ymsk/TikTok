package dao

import (
	"database/sql"
	"fmt"
	"strconv"
)

// UserExist 查询用户是否存在并返回其UserId,不存在或错误返回错误
func UserExist(Token string) (UserId string, err error) {
	return
}

// FollowExist 查询关注信息是否存在
func FollowExist(UserId string, ToUserId string) (Exist bool, err error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return false, fmt.Errorf("Parameter error :%s", err)
	}
	toUserId, err := strconv.Atoi(ToUserId)
	if err != nil {
		return false, fmt.Errorf("Parameter error :%s", err)
	}
	sqlStatement := "SELECT * FROM Follows WHERE following_user_id=? and followed_user_id=?;"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return false, fmt.Errorf("select init error :%s", err)
	}
	res, err := stmt.Query(userId, toUserId)
	if err != nil {
		return false, fmt.Errorf("select execute error :%s", err)
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
		return fmt.Errorf("Parameter error :%s", err)
	}
	toUserId, err := strconv.Atoi(ToUserId)
	if err != nil {
		return fmt.Errorf("Parameter error :%s", err)
	}
	sqlStatement := "INSERT INTO Follows (`following_user_id`, `followed_user_id`, `relationship`) VALUES (?, ?, ?);"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return fmt.Errorf("insert init error :%s", err)
	}
	_, err = stmt.Exec(userId, toUserId, Relationship)
	if err != nil {
		return fmt.Errorf("insert execute error :%s", err)
	}
	return nil
}

// ChangeRelation 修改关注信息
func ChangeRelation(UserId string, ToUserId string, Relationship int) (err error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return fmt.Errorf("Parameter error :%s", err)
	}
	toUserId, err := strconv.Atoi(ToUserId)
	if err != nil {
		return fmt.Errorf("Parameter error :%s", err)
	}
	sqlStatement := "UPDATE Follows set relationship=? WHERE following_user_id=? and followed_user_id=?;"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return fmt.Errorf("update init error :%s", err)
	}
	_, err = stmt.Exec(Relationship, userId, toUserId)
	if err != nil {
		return fmt.Errorf("update execute error :%s", err)
	}
	return nil
}

// DeleteFollow 删除关注信息
func DeleteFollow(UserId string, ToUserId string) (err error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return fmt.Errorf("Parameter error :%s", err)
	}
	toUserId, err := strconv.Atoi(ToUserId)
	if err != nil {
		return fmt.Errorf("Parameter error :%s", err)
	}
	sqlStatement := "DELETE FROM Follows WHERE following_user_id=? and followed_user_id=?;"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return fmt.Errorf("delete init error :%s", err)
	}
	_, err = stmt.Exec(userId, toUserId)
	if err != nil {
		return fmt.Errorf("delete execute error :%s", err)
	}
	return nil
}

// SelectFollowList 查询所有关注的信息列表
func SelectFollowList(UserId string) (*sql.Rows, error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return nil, fmt.Errorf("parameter error :%s", err)
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
		return nil, fmt.Errorf("select init error :%s", err)
	}
	res, err := stmt.Query(userId)
	if err != nil {
		return nil, fmt.Errorf("select execute error :%s", err)
	}
	return res, nil
}

// SelectFollowerList 查询所有粉丝的信息列表
func SelectFollowerList(UserId string) (*sql.Rows, error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return nil, fmt.Errorf("parameter error :%s", err)
	}
	sqlStatement :=
		`SELECT
			Users.id,
			Users.user_name,
			fans,
			follows,
		IF
			(f.relationship= ?, 1, 0) AS is_followed -- 修改占位符
		FROM
			Users
			LEFT JOIN Follows f ON Users.id = f.following_user_id 
			LEFT JOIN Follows d ON Users.id = f.followed_user_id
		WHERE
			f.followed_user_id = ?;`
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("select init error :%s", err)
	}
	res, err := stmt.Query(userId, userId)
	if err != nil {
		return nil, fmt.Errorf("select execute error :%s", err)
	}
	return res, nil
}

// SelectFriendList 查询所有互关的信息列表
func SelectFriendList(UserId string) (*sql.Rows, error) {
	userId, err := strconv.Atoi(UserId)
	if err != nil {
		return nil, fmt.Errorf("parameter error :%s", err)
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
		return nil, fmt.Errorf("select init error :%s", err)
	}
	res, err := stmt.Query(userId)
	if err != nil {
		return nil, fmt.Errorf("select execute error :%s", err)
	}
	return res, nil
}
