package mysql_model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

// FileLink 用户拥有的文件
type FileLink struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	DeletedAt soft_delete.DeletedAt
	Username  string    `gorm:"index;size:16;not null;default:'';uniqueIndex:username_file,priority:1"` // todo not null
	Filename  string    `gorm:"index;size:255;not null"`
	IsPublic  bool      `gorm:"not null;default:true"`
	FileUUID  string    `gorm:"type:char(64);not null;uniqueIndex:username_file,priority:2"`
	FileMeta  *FileMeta `gorm:"foreignkey:UUID;references:FileUUID"`
}
