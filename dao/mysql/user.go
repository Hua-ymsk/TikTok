package mysql

import (
	"gorm.io/gorm"
	"tiktok/models"
)

// 查询操作
func QueryUserName(username string) (userN bool, userid int64, password string, nickname string, fans int64, follows int64, err error) {
	var user models.User
	res := db.Where("user_name", username).Take(&user)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return false, user.ID, user.PassWord, user.NickName, user.Fans, user.Follows, nil
	}
	if res.Error == gorm.ErrRecordNotFound {
		return false, user.ID, user.PassWord, user.NickName, user.Fans, user.Follows, err
	}
	return true, user.ID, user.PassWord, user.NickName, user.Fans, user.Follows, err
}

// 注册用户
func RegisterUser(user *models.User) (userid int64) {
	db.Create(user)
	return user.ID
}
func QueryUserID(userId int64, usernowId int64) (userN bool, userid int64, password string, nickname string, fans int64, follows int64, isfollow bool, err error) {
	var user models.User
	var userNow models.User
	var follow models.Follow
	res := db.Where("id", userId).Take(&user)

	r := db.Where("id", usernowId).Take(&userNow)
	condition := "following_user_id = ? AND followed_user_id = ?"
	re := db.Where(condition, usernowId, userId).Take(&follow)
	//
	if r.Error == gorm.ErrRecordNotFound {

		return false, user.ID, user.PassWord, user.NickName, user.Fans, user.Follows, false, err
	}
	if r.Error != nil && r.Error != gorm.ErrRecordNotFound {

		return false, user.ID, user.PassWord, user.NickName, user.Fans, user.Follows, false, err
	}

	if res.Error == gorm.ErrRecordNotFound {

		return false, user.ID, user.PassWord, user.NickName, user.Fans, user.Follows, false, nil
	}
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {

		return false, user.ID, user.PassWord, user.NickName, user.Fans, user.Follows, false, nil

	}
	if re.Error != nil {

		return false, user.ID, user.PassWord, user.NickName, user.Fans, user.Follows, false, err
	}

	if follow.FollowingID == userNow.ID && follow.FollowedID == user.ID {
		user.IsFollow = true
		return true, user.ID, user.PassWord, user.NickName, user.Fans, user.Follows, user.IsFollow, nil
	}
	return false, user.ID, user.PassWord, user.NickName, user.Fans, user.Follows, false, nil
}
