package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Birthday  time.Time
	Age       int
	Name      string `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
	Num       int    `gorm:"AUTO_INCREMENT"` // 自增

	Emails []Email // One-To-Many (拥有多个 - Email表的UserID作外键)

}
