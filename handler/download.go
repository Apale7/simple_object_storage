package handler

import (
	"fmt"
	"net/url"
	"strconv"

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

func DownloadSelf(c *gin.Context) {
	fileID, ok := c.GetQuery("file_id")
	if !ok {
		utils.RetErr(c, fmt.Errorf("file_id is required"))
		return
	}
	fileIDInt, err := strconv.ParseUint(fileID, 10, 32)
	if err != nil {
		utils.RetErr(c, fmt.Errorf("file_id is invalid"))
		return
	}
	downloadFile(c, fileIDInt)
}

func downloadFile(c *gin.Context, linkID uint64) {
	link, err := service.GetFileLink(c, uint(linkID))
	if err != nil {
		utils.RetErr(c, err)
		return
	}
	
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(link.Filename))) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(base + link.FileUUID)
}

func FileExist(c *gin.Context) {
	fileID, ok := c.GetQuery("file_id")
	if !ok {
		utils.RetErr(c, fmt.Errorf("file_id is required"))
		c.Abort()
		return
	}
	fileIDInt, err := strconv.ParseUint(fileID, 10, 32)
	if err != nil {
		utils.RetErr(c, fmt.Errorf("file_id is invalid"))
		c.Abort()
		return
	}
	link, err := service.GetFileLink(c, uint(fileIDInt))
	if err != nil {
		utils.RetErr(c, err)
		c.Abort()
		return
	}
	c.Set("file_link", link)
}
