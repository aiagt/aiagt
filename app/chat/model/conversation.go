package model

type Conversation struct {
	Base

	Title string `gorm:"column:title;type:varchar(255)"`
}

type ConversationOptional struct {
}
