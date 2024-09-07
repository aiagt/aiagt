package model

type User struct {
	Base

	Username      string `gorm:"column:username;index"`
	Password      string `gorm:"column:password"`
	Email         string `gorm:"column:email;index"`
	PhoneNumber   string `gorm:"column:phone_number"`
	Signature     string `gorm:"column:signature"`
	Homepage      string `gorm:"column:homepage"`
	DescriptionMd string `gorm:"column:description_md;type:text"`
	Github        string `gorm:"column:github"`
	Avatar        string `gorm:"column:avatar"`
}
