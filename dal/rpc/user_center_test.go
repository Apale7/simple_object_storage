package rpc

import (
	"context"
	"reflect"
	"testing"

	user_center "Apale7/simple_object_storage/proto/user-center"
)

func TestLogin(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name     string
		args     args
		wantResp *user_center.LoginResponse
		wantErr  bool
	}{
		{
			name: "login",
			args: args{
				ctx:      context.Background(),
				username: "admin",
				password: "admin",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := Login(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Login() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
