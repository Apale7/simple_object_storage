package mysql

import (
	"context"
	"fmt"

	mysql_model "Apale7/simple_object_storage/dal/mysql/model"

	"gorm.io/gorm"
)

func GetFileLink(ctx context.Context, queryFuncs ...Option) ([]mysql_model.FileLink, error) {
	db := getDB().WithContext(ctx).Model(&mysql_model.FileLink{})
	for _, queryFunc := range queryFuncs {
		db = queryFunc(db)
	}
	var fileLinks []mysql_model.FileLink
	err := db.Preload("FileMeta").Find(&fileLinks).Error
	return fileLinks, err
}

func CreateFileLink(ctx context.Context, fileLink mysql_model.FileLink) error {
	db := getDB().WithContext(ctx)
	return db.Create(&fileLink).Error
}

func ID(id uint) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id=?", id)
	}
}

func Filename(filename *string) Option {
	return func(db *gorm.DB) *gorm.DB {
		if filename == nil {
			return db
		}
		return db.Where("filename like ?", fmt.Sprintf("%%%s%%", *filename))
	}
}

func IsPublic(isPublic *bool) Option {
	return func(db *gorm.DB) *gorm.DB {
		if isPublic == nil {
			return db
		}
		return db.Where("is_public=?", *isPublic)
	}
}

func Username(username *string) Option {
	return func(db *gorm.DB) *gorm.DB {
		if username == nil {
			return db
		}
		return db.Where("username=?", *username)
	}
}
