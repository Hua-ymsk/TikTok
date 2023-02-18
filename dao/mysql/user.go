package mysql

import (
	"fmt"
	"gorm.io/gorm"
	"tiktok/models"
)

// 查询操作
func QueryUserName(username string) (userN bool, user *models.User, err error) {

	res := db.Where("user_name", username).Take(&user)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return false, nil, nil
	}
	if res.Error == gorm.ErrRecordNotFound {
		return false, nil, err
	}
	return true, user, err
}

// 注册用户
func RegisterUser(user *models.User) (userid int64) {
	db.Select("user_name", "password", "nickname").Create(user)
	return user.ID
}

func QueryUserID(userId int64, usernowId any) (workCount int64, favoriteCount int64, responseUser *models.User, isFollow bool, err error) {
	var userNow models.User
	var follow models.Follow
	var Count int64
	var workCounts int64
	res := db.Where("id", userId).Take(&responseUser)
	res2 := db.Model(models.Like{}).Where("user_id = ?", userId).Count(&Count)       //查询喜欢数量
	res3 := db.Model(models.Video{}).Where("user_id = ?", userId).Count(&workCounts) //查询作品数量
	if res2.Error == gorm.ErrRecordNotFound {
		return 0, 0, nil, false, fmt.Errorf("string to int error:%v", res2)
	}
	if res2.Error != nil && res2.Error != gorm.ErrRecordNotFound {
		return 0, 0, nil, false, fmt.Errorf("string to int error:%v", res2)
	}
	if res3.Error == gorm.ErrRecordNotFound {
		return 0, 0, nil, false, fmt.Errorf("string to int error:%v", res2)
	}
	if res3.Error != nil && res3.Error != gorm.ErrRecordNotFound {
		return 0, 0, nil, false, fmt.Errorf("string to int error:%v", res3)
	}

	r := db.Where("id", usernowId).Take(&userNow)
	condition := "following_user_id = ? AND followed_user_id = ?"
	re := db.Where(condition, usernowId, userId).Find(&follow)
	//当前登录用户不存在
	if r.Error == gorm.ErrRecordNotFound {

		return workCounts, Count, nil, false, err
	}
	//当前登录用户查询有错误
	if r.Error != nil && r.Error != gorm.ErrRecordNotFound {

		return workCounts, Count, nil, false, err
	}
	//查询的用户不存在
	if res.Error == gorm.ErrRecordNotFound {

		return workCounts, Count, nil, false, err
	}
	//查询的用户数据库有错误
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {

		return workCounts, Count, nil, false, err
	}
	//查询关注列表出错误

	//当查询列表有我们查到的
	if follow.FollowingID == userNow.ID && follow.FollowedID == responseUser.ID {
		return workCounts, Count, responseUser, true, nil
		//return true, user.ID, user.PassWord, user.NickName, user.Fans, user.Follows, user.IsFollow, nil
	}
	//查不到关系
	if re.RowsAffected == 0 {
		return workCounts, Count, responseUser, false, nil
	}
	//查表出错
	if re.Error != nil && re.Error != gorm.ErrRecordNotFound {
		return workCounts, Count, nil, false, err
	}
	//默认返回
	return workCounts, Count, responseUser, false, nil
}
