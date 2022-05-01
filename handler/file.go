package handler

import (
	"fmt"

	"Apale7/simple_object_storage/dal/mysql"
	"Apale7/simple_object_storage/model"

	"github.com/Apale7/common/utils"
	"github.com/gin-gonic/gin"
)

func FileList(c *gin.Context) {
	var reqParams model.FileListReq
	if err := c.ShouldBindQuery(&reqParams); err != nil {
		utils.RetErr(c, fmt.Errorf("parse params error: %w", err))
		return
	}

	files, err := mysql.GetFileLink(c, mysql.Filename(reqParams.Filename), mysql.IsPublic(reqParams.IsPublic))
	if err != nil {
		utils.RetErr(c, err)
		return
	}
	resp := model.FileListResp{}
	for _, file := range files {
		resp.Files = append(resp.Files, model.FileInfo{
			Filename: file.Filename,
			Username: file.Username,
			Size:     file.FileMeta.Size,
		})
	}
	utils.RetData(c, resp)
}
