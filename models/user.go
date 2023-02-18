package models

type User struct {
	ID         int64
	UserName   string `gorm:"column:user_name"`
	PassWord   string `gorm:"column:password"`
	NickName   string `gorm:"column:nickname"`
	Fans       int64  `gorm:"column:fans"`
	Follows    int64  `gorm:"column:follows"`
	AvatarUrl  string `gorm:"column:avatar_url"`
	BackGround string `gorm:"column:background_image"`
	Signature  string `gorm:"column:signature"`
}
