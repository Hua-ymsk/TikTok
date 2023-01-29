package models

type Follow struct {
	ID          int64
	FollowingID int64 `gorm:"column:following_user_id"`
	FollowedID  int64 `gorm:"column:followed_user_id"`
	Relation    int64 `gorm:"relationship"`
}
