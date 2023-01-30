package mysql

import (
	"tiktok/models"
	"tiktok/setting"

	"gorm.io/gorm"
)

func GetVideosByLatestTime(latest_time int64) (list []models.Video, err error) {
	conf := setting.Conf.VideoConfig
	list = make([]models.Video, 0, conf.PageSize)
	res := db.Where("timestamp < ?", latest_time).Order("timestamp desc").Limit(int(conf.PageSize)).Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}

	return
}

func SaveVideo(video *models.Video) (err error) {
	if res := db.Save(video); res.Error != nil {
		return res.Error
	}
	return
}

func GetPublishList(user_id int64) (list []models.Video, err error) {
	list = make([]models.Video, 0, 20)
	res := db.Where("user_id = ?", user_id).Find(&list)
	if res.Error != nil {
		return nil, res.Error
	}

	return
}

// 是否关注
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

// 是否点赞
func CheckFavorite(user_id int64, video_id int64) (isfavorite bool, err error) {
	var like models.Like
	res := db.Select("user_id").Where("user_id = ? and video_id = ?").Take(&like)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return false, res.Error
	}
	if res.Error == gorm.ErrRecordNotFound {
		return false, nil
	}
	
	return true, nil
}
