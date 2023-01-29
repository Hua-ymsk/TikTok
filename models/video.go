package models

type Video struct {
	ID        int64
	UserID    int64  `gorm:"column:user_id"`
	Play_url  string `gorm:"column:video_url"`
	Cover_url string `gorm:"column:cover_url"`
	Titile    string `gorm:"column:title"`
	TimeStamp int64  `gorm:"column:timestamp"`
}
