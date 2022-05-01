package handler

import (
	"fmt"

	"Apale7/simple_object_storage/dal/mysql"
	"Apale7/simple_object_storage/dal/redis"
	"Apale7/simple_object_storage/model"

	"github.com/Apale7/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
)

// 分享文件，可设置提取密码
func ShareFile(c *gin.Context) {
	var reqBody model.ShareFileReq
	var resp model.ShareFileRes
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		utils.RetErr(c, fmt.Errorf("parse params error: %w", err))
		return
	}
	// 检查文件是否存在
	fileLink, err := mysql.GetFileLink(c, mysql.ID(reqBody.FileID))
	if err != nil {
		utils.RetErr(c, fmt.Errorf("mysql error: %w", err))
		return
	}
	if len(fileLink) == 0 {
		utils.RetErr(c, fmt.Errorf("file not found"))
		return
	}
	// 分享文件
	shareID := uuid.New()
	pwd := genPassword(reqBody.NeedPwd, reqBody.Password)
	err = redis.StoreShare(shareID, pwd, reqBody.FileID, reqBody.Duration)
	if err != nil {
		utils.RetErr(c, fmt.Errorf("redis error: %w", err))
		return
	}
	// 返回分享码
	resp.ShareID = shareID
	if reqBody.NeedPwd && reqBody.Password == nil {
		resp.Password = &pwd
	}
	utils.RetData(c, resp)
}

func genPassword(needPwd bool, password *string) string {
	// 需要密码, 未指定密码, 返回随机密码
	if needPwd && password == nil {
		return utils.RandomAlphaString(8)
	}
	// 需要密码, 指定了密码, 返回指定的密码
	if needPwd && password != nil {
		return *password
	}
	// 不需要密码
	return ""
}
