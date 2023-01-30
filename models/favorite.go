package models

type Like struct {
	ID      int `gorm:"primaryKey"`
	UserId  int `gorm:"column:user_id"`
	VideoId int `gorm:"column:video_id"`
}
