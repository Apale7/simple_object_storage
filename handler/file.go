package handler

import (
	"fmt"
	"net/http"

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
			ID:       file.ID,
			Filename: file.Filename,
			Username: file.Username,
			Size:     file.FileMeta.Size,
		})
	}
	utils.RetData(c, resp)
}

func ListPage(c *gin.Context) {
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
	resp := make([]gin.H, 0)
	for _, file := range files {
		resp = append(resp, map[string]interface{}{
			"id":       file.ID,
			"filename": file.Filename,
			"username": file.Username,
			"size":     file.FileMeta.Size >> 10,
		})
	}
	c.HTML(http.StatusOK, "list.tmpl", gin.H{
		"files": resp,
	})
}
