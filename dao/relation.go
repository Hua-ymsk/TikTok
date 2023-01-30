package dao

import (
	"database/sql"
	"errors"
	"fmt"
)

//// TokenResolution 解析Token获得userId
//func TokenResolution(Token string) (userId string, err error) {
//	//测试用
//	userId, err = Token, nil
//	////正式使用
//	//claims, err := middleware.ParseToken(Token)
//	//userId = claims.UserID
//	return
//}

// UserExist 检验该UserId对应用户是否存在
func UserExist(UserId int64) (err error) {
	sqlStatement := "SELECT * FROM users WHERE id=?;"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return fmt.Errorf("select init error: %w", err)
	}
	res, err := stmt.Query(UserId)
	if err != nil {
		return fmt.Errorf("select execute error: %w", err)
	}
	defer res.Close()
	if res.Next() == false {
		return errors.New("user not exist error")
	}
	return nil
}

// FollowExist 查询关注信息是否存在
func FollowExist(UserId int64, ToUserId int64) (exist bool, err error) {
	//userId, err := strconv.Atoi(UserId)
	//if err != nil {
	//	return false, fmt.Errorf("parameter error: %w", err)
	//}
	//toUserId, err := strconv.Atoi(ToUserId)
	//if err != nil {
	//	return false, fmt.Errorf("parameter error: %w", err)
	//}
	sqlStatement := "SELECT * FROM follows WHERE following_user_id=? and followed_user_id=?;"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return false, fmt.Errorf("select init error: %w", err)
	}
	res, err := stmt.Query(UserId, ToUserId)
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
func InsertFollow(UserId int64, ToUserId int64, Relationship int) (err error) {
	//userId, err := strconv.Atoi(UserId)
	//if err != nil {
	//	return fmt.Errorf("parameter error: %w", err)
	//}
	//toUserId, err := strconv.Atoi(ToUserId)
	//if err != nil {
	//	return fmt.Errorf("parameter error: %w", err)
	//}
	sqlStatement := "INSERT INTO follows (`following_user_id`, `followed_user_id`, `relationship`) VALUES (?, ?, ?);"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return fmt.Errorf("insert init error: %w", err)
	}
	_, err = stmt.Exec(UserId, ToUserId, Relationship)
	if err != nil {
		return fmt.Errorf("insert execute error: %w", err)
	}
	return nil
}

// UpdateRelation 修改关注信息
func UpdateRelation(UserId int64, ToUserId int64, Relationship int) (err error) {
	//userId, err := strconv.Atoi(UserId)
	//if err != nil {
	//	return fmt.Errorf("parameter error: %w", err)
	//}
	//toUserId, err := strconv.Atoi(ToUserId)
	//if err != nil {
	//	return fmt.Errorf("parameter error: %w", err)
	//}
	sqlStatement := "UPDATE follows set relationship=? WHERE following_user_id=? and followed_user_id=?;"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return fmt.Errorf("update init error: %w", err)
	}
	_, err = stmt.Exec(Relationship, UserId, ToUserId)
	if err != nil {
		return fmt.Errorf("update execute error: %w", err)
	}
	return nil
}

// DeleteFollow 删除关注信息
func DeleteFollow(UserId int64, ToUserId int64) (err error) {
	//userId, err := strconv.Atoi(UserId)
	//if err != nil {
	//	return fmt.Errorf("parameter error: %w", err)
	//}
	//toUserId, err := strconv.Atoi(ToUserId)
	//if err != nil {
	//	return fmt.Errorf("parameter error: %w", err)
	//}
	sqlStatement := "DELETE FROM follows WHERE following_user_id=? and followed_user_id=?;"
	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		return fmt.Errorf("delete init error: %w", err)
	}
	_, err = stmt.Exec(UserId, ToUserId)
	if err != nil {
		return fmt.Errorf("delete execute error: %w", err)
	}
	return nil
}

// UpdInsRelation 关注时，有对方信息存在，执行修改标记并进行插入
func UpdInsRelation(UserId int64, ToUserId int64) (err error) {
	//修改标记
	begin, _ := db.Begin()
	err = UpdateRelation(ToUserId, UserId, 1)
	if err != nil {
		err := begin.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	//插入信息
	err = InsertFollow(UserId, ToUserId, 1)
	if err != nil {
		err := begin.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	err = begin.Commit()
	if err != nil {
		return err
	}
	return nil
}

// UpdDelRelation 取消关注时，有对方信息存在，执行修改标记并进行删除
func UpdDelRelation(UserId int64, ToUserId int64) (err error) {
	//修改标记
	begin, _ := db.Begin()
	err = UpdateRelation(ToUserId, UserId, 0)
	if err != nil {
		err := begin.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	//插入信息
	err = DeleteFollow(UserId, ToUserId)
	if err != nil {
		err := begin.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	err = begin.Commit()
	if err != nil {
		return err
	}
	return nil
}

// SelectFollowList 查询所有关注的信息列表
func SelectFollowList(UserId int64) (*sql.Rows, error) {
	//userId, err := strconv.Atoi(UserId)
	//if err != nil {
	//	return nil, fmt.Errorf("parameter error: %w", err)
	//}
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
	res, err := stmt.Query(UserId)
	if err != nil {
		return nil, fmt.Errorf("select execute error: %w", err)
	}
	return res, nil
}

// SelectFollowerList 查询所有粉丝的信息列表
func SelectFollowerList(UserId int64) (*sql.Rows, error) {
	//userId, err := strconv.Atoi(UserId)
	//if err != nil {
	//	return nil, fmt.Errorf("parameter error: %w", err)
	//}
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
	res, err := stmt.Query(UserId, UserId)
	if err != nil {
		return nil, fmt.Errorf("select execute error: %w", err)
	}
	return res, nil
}

// SelectFriendList 查询所有互关的信息列表
func SelectFriendList(UserId int64) (*sql.Rows, error) {
	//userId, err := strconv.Atoi(UserId)
	//if err != nil {
	//	return nil, fmt.Errorf("parameter error: %w", err)
	//}
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
	res, err := stmt.Query(UserId)
	if err != nil {
		return nil, fmt.Errorf("select execute error: %w", err)
	}
	return res, nil
}
