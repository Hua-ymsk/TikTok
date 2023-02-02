package mysql

import (
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
func QueryUserID(userId int64, usernowId any) (responseUser *models.User, isFollow bool, err error) {
	//var user models.User
	var userNow models.User
	var follow models.Follow
	res := db.Where("id", userId).Take(&responseUser)

	r := db.Where("id", usernowId).Take(&userNow)
	condition := "following_user_id = ? AND followed_user_id = ?"
	re := db.Where(condition, usernowId, userId).Take(&follow)
	//当前登录用户不存在
	if r.Error == gorm.ErrRecordNotFound {

		return nil, false, err
	}
	//当前登录用户查询有错误
	if r.Error != nil && r.Error != gorm.ErrRecordNotFound {

		return nil, false, err
	}
	//查询的用户不存在
	if res.Error == gorm.ErrRecordNotFound {

		return nil, false, err
	}
	//查询的用户数据库有错误
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {

		return nil, false, err
	}
	//查询关注列表出错误

	//当查询列表有我们查到的
	if follow.FollowingID == userNow.ID && follow.FollowedID == responseUser.ID {

		return responseUser, true, nil
		//return true, user.ID, user.PassWord, user.NickName, user.Fans, user.Follows, user.IsFollow, nil
	}
	//查不到关系
	if re.Error == gorm.ErrRecordNotFound {
		return responseUser, false, nil
	}
	//查表出错
	if re.Error != nil && re.Error != gorm.ErrRecordNotFound {
		return nil, false, err
	}
	//默认返回
	return responseUser, false, nil
}
