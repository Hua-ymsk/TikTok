package models

type Video struct {
	CommentCount  int64  `gorm:"column:comments_num"`
	CoverURL      string `gorm:"column:cover_url"`
	FavoriteCount int64  `gorm:"column:likes_num"`
	ID            int64  `gorm:"column:id"`
	PlayURL       string `gorm:"column:video_url"`
	Title         string `gorm:"column:title"`
	TimeStamp     int64  `gorm:"column:timestamp"`
	UserID        int64  `gorm:"column:user_id"`
}
