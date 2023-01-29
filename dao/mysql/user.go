package mysql

import (
	"gorm.io/gorm"
	"tiktok/models"
)

// 查询用户名
func QueryUserName(username string) (userN bool, userid int64, password string, err error) {
	var user models.User
	res := db.Where("user_name", username).Take(&user)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return false, user.ID, user.PassWord, err
	}
	if res.Error == gorm.ErrRecordNotFound {
		return false, user.ID, user.PassWord, nil
	}
	return true, user.ID, user.PassWord, err
}

// 注册用户
func RegisterUser(user *models.User) (userid int64) {
	db.Create(user)
	return user.ID
}
