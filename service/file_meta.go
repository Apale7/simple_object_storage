package service

import (
	"context"
	"fmt"

	"Apale7/simple_object_storage/dal/mysql"
	mysql_model "Apale7/simple_object_storage/dal/mysql/model"
)

func GetFileMeta(ctx context.Context, uuid string) (*mysql_model.FileMeta, error) {
	metas, err := mysql.GetFileMeta(ctx, mysql.UUID(uuid))
	if err != nil {
		return nil, err
	}
	if len(metas) == 0 {
		return nil, fmt.Errorf("file not found")
	}
	return &metas[0], nil
}

func GetFileLink(ctx context.Context, linkID uint) (*mysql_model.FileLink, error) {
	links, err := mysql.GetFileLink(ctx, mysql.ID(linkID))
	if err != nil {
		return nil, err
	}
	if len(links) == 0 {
		return nil, fmt.Errorf("link not found")
	}
	return &links[0], nil
}
