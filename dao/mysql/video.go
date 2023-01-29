package mysql

import (
	"tiktok/models"

	"gorm.io/gorm"
)

func SaveVideo(video *models.Video) (err error) {
	if res := db.Save(video); res.Error != nil {
		return res.Error
	}
	return
}

func PublishList(user_id int64) (list []models.Video, err error) {
	list = make([]models.Video, 0, 20)
	res := db.Where("user_id = ?", user_id).Find(&list)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return nil, res.Error
	}

	return
}
// db.Limit(p.PageSize).Offset((p.PageNum - 1) * p.PageSize).Order(p.Order + " " + p.Sort).Find(&tableData)

func ChekFollow(sender_id, user_id int64) (idfollow bool, err error) {
	var follow models.Follow
	// 避免回表
	condition := "following_user_id = ? and followed_user_id = ?"
	res := db.Select("following_user_id").Where(condition, sender_id, user_id).Take(&follow)
	// 数据库出错
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return false, res.Error
	}
	// 未关注
	if res.Error == gorm.ErrRecordNotFound {
		return false, nil
	}

	return true, nil
}
