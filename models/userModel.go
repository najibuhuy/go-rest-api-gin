package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"type:varchar(100); unique" json:"email"`
	Password string `json:"password"`
}
