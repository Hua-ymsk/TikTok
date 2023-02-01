package models

type Comment struct {
	ID        int64  `gorm:"primaryKey"`
	UserId    int64  `column:"user_id"`
	VideoId   int64  `column:"video_id"`
	Timestamp int64  `column:"timestamp"`
	Content   string `column:"content"`
}
