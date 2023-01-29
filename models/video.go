package models

type Video struct {
	ID        int64
	UserID    int64  `gorm:"column:user_id"`
	PlayURL  string `gorm:"column:video_url"`
	CoverURL string `gorm:"column:cover_url"`
	Title    string `gorm:"column:title"`
	TimeStamp int64  `gorm:"column:timestamp"`
}
