package model

type Conversation struct {
	Base

	AppID  int64  `gorm:"column:app_id;index:idx_app_id_user_id"`
	UserID int64  `gorm:"column:user_id;index:idx_app_id_user_id"`
	Title  string `gorm:"column:title;type:varchar(255)"`
}

type ConversationOptional struct {
	Title *string `gorm:"column:title;type:varchar(255)"`
}
