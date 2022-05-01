package handler

import (
	"fmt"

	"Apale7/simple_object_storage/dal/redis"
	"Apale7/simple_object_storage/model"
	"Apale7/simple_object_storage/service"

	"github.com/Apale7/common/utils"
	"github.com/gin-gonic/gin"
	goredis "github.com/go-redis/redis"
)

func DownloadBySharing(c *gin.Context) {
	// 解析参数
	var reqParams model.DownloadBySharingReq
	if err := c.ShouldBindQuery(&reqParams); err != nil {
		utils.RetErr(c, fmt.Errorf("parse params error: %w", err))
		return
	}
	// 检查分享码
	fmt.Println(reqParams)
	linkID, err := redis.CheckShare(reqParams.ShareID, reqParams.Password)
	if err != nil {
		if err == goredis.Nil {
			utils.RetErr(c, fmt.Errorf("分享码/密码错误或分享已过期"))
			return
		}
		utils.RetErr(c, fmt.Errorf("redis error: %w", err))
		return
	}
	downloadFile(c, linkID)
}

func downloadFile(c *gin.Context, linkID uint64) {
	link, err := service.GetFileLink(c, uint(linkID))
	if err != nil {
		utils.RetErr(c, err)
		return
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", link.Filename)) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	// c.String(200, "asd")
	c.File(base + link.FileUUID)
}
