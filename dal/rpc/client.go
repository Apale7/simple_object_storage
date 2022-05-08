package rpc

import (
	user_center "Apale7/simple_object_storage/proto/user-center"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var userCenterClient user_center.UserCenterClient

func init() {
	userCenterClient = user_center.NewUserCenterClient(getConn("apale7.cn:9999"))
}

func getConn(addr string) *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		logrus.Fatalf("%+v", err)
	}
	return conn
}
