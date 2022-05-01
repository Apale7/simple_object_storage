package mysql

import (
	"context"

	mysql_model "Apale7/simple_object_storage/dal/mysql/model"

	"gorm.io/gorm"
)

func GetFileMeta(ctx context.Context, queryFuncs ...Option) (fileMetas []mysql_model.FileMeta, err error) {
	db := getDB().WithContext(ctx).Model(&mysql_model.FileMeta{})
	for _, queryFunc := range queryFuncs {
		db = queryFunc(db)
	}
	err = db.Find(&fileMetas).Error
	return fileMetas, err
}

func CreateFileMeta(ctx context.Context, fileMeta *mysql_model.FileMeta) error {
	db := getDB().WithContext(ctx)
	return db.Create(fileMeta).Error
}

func UUID(uuid string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("uuid=?", uuid)
	}
}

func SHA256(sha256 string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("sha256=?", sha256)
	}
}
