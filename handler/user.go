package handler

import (
	"context"

	"Apale7/simple_object_storage/dal/rpc"
	"Apale7/simple_object_storage/model"
	user_center "Apale7/simple_object_storage/proto/user-center"

	"github.com/Apale7/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.LoginReq
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
		return
	}
	logrus.Infof("reqBody: %+v", reqBody)
	resp, err := rpc.Login(ctx, reqBody.Username, reqBody.Password)
	if err != nil {
		logrus.Warnf("login failed, err: %v", err)
		utils.RetErr(c, err)
		return
	}
	var loginRes model.LoginResp
	loginRes.LoginResponse = *resp
	loginRes.RefreshToken = ""
	loginRes.Auth = resp.AuthList
	utils.RetData(c, loginRes)
}

func Register(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.RegisterReq
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
		return
	}
	logrus.Infof("reqBody: %+v", reqBody)

	err := rpc.Register(ctx, &user_center.User{Username: reqBody.Username, Password: reqBody.Password}, &user_center.UserExtra{Nickname: reqBody.Nickname, PhoneNumber: reqBody.PhoneNumber})
	if err != nil {
		logrus.Warnf("register failed, err: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetSuccess(c)
}
