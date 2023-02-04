package models

type Chat struct {
	ID            int64  `gorm:"primaryKey"`
	SendUserId    int64  `column:"send_user_id"`
	ReceiveUserId int64  `column:"receive_user_id"`
	Timestamp     int64  `column:"timestamp"`
	Content       string `column:"content"`
}
