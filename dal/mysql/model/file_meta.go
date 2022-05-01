package mysql_model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

// FileMeta 文件的元信息. 不会存储重复的文件, 文件名就是uuid
type FileMeta struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt soft_delete.DeletedAt
	UUID      string `gorm:"uniqueIndex;size:64;not null"`
	SHA256    string `gorm:"type:char(64);not null"`
	Size      int64  `gorm:"not null"`
}
