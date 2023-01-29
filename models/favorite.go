package models

type Like struct {
	ID      int `gorm:"primaryKey"`
	UserId  int `column:"user_id"`
	VideoId int `column:"video_id"`
}
