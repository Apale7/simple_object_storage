package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	config "Apale7/simple_object_storage/config_loader"
	"Apale7/simple_object_storage/dal/mysql"
	mysql_model "Apale7/simple_object_storage/dal/mysql/model"

	"github.com/Apale7/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/sirupsen/logrus"
)

var base string

func Init() {
	base = config.Get("file_root")
	fmt.Println("file_root: ", base)
	err := os.MkdirAll(base, 0o755)
	if err != nil {
		logrus.Fatalf("create dir failed, err: %v", err)
		panic(fmt.Errorf("create dir failed, err: %w", err))
	}
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	// 判空
	if err != nil {
		utils.RetErr(c, fmt.Errorf("upload file error: %w", err))
		return
	}
	if file == nil {
		utils.RetErr(c, fmt.Errorf("upload file error: %w", err))
		return
	}
	logrus.WithContext(c).Info("Get file: ", file.Filename)
	fileUUID := uuid.New()
	// 保存文件
	if err = c.SaveUploadedFile(file, base+fileUUID); err != nil {
		utils.RetErr(c, fmt.Errorf("upload file err: %s", err.Error()))
		return
	}

	// 计算sha256
	sha256, err := getSHA256(file)
	if err != nil {
		utils.RetErr(c, fmt.Errorf("calc sha256 err: %w", err))
		return
	}
	c.Set("file_sha256", sha256)
	logrus.Info(len(sha256))
	// 记录元数据
	fileMeta := mysql_model.FileMeta{
		UUID:   fileUUID,
		SHA256: c.GetString("file_sha256"),
		Size:   file.Size,
	}
	err = mysql.CreateFileMeta(c, &fileMeta)
	if err != nil {
		utils.RetErr(c, fmt.Errorf("mysql error: %w", err))
		return
	}
	fmt.Println(fileMeta.ID)

	// 用户关联到文件元数据
	err = mysql.CreateFileLink(c, mysql_model.FileLink{
		FileUUID: fileMeta.UUID,
		Filename: file.Filename,
		IsPublic: true,
	})
	if err != nil {
		utils.RetErr(c, fmt.Errorf("mysql error: %w", err))
		return
	}
	utils.RetSuccess(c)
}

func UniqueFile(c *gin.Context) {
	file, err := c.FormFile("file")
	// 判空
	if err != nil {
		utils.RetErr(c, fmt.Errorf("upload file error: %w", err))
		return
	}
	if file == nil {
		utils.RetErr(c, fmt.Errorf("upload file error: %w", err))
		return
	}
	logrus.WithContext(c).Info("Get file: ", file.Filename)
	sha256, err := getSHA256(file)
	fileMetas, err := mysql.GetFileMeta(c, mysql.SHA256(sha256))
	if err != nil {
		utils.RetErr(c, fmt.Errorf("mysql error: %w", err))
		return
	}
	if len(fileMetas) > 0 {
		c.Abort()
		link := mysql_model.FileLink{
			FileUUID: fileMetas[0].UUID,
			Filename: file.Filename,
			IsPublic: true,
		}

		err = mysql.CreateFileLink(c, link)
		if err != nil {
			utils.RetErr(c, fmt.Errorf("mysql error: %w", err))
			return
		}
		return
	}
}

func getSHA256(file *multipart.FileHeader) (string, error) {
	ha := sha256.New()
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	if _, err := io.Copy(ha, src); err != nil {
		return "", err
	}
	return hex.EncodeToString(ha.Sum(nil)), nil
}
