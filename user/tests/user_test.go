package tests

import (
	"context"
	"fmt"
	"proto/user"
	"testing"
	"user/config"
	"user/service"
	"utils"
)

func init() {
	if err := utils.ScanConfig(&config.Config); err != nil {
		panic(err)
	}
	config.Init()
}

func TestLogin(t *testing.T) {
	var userServer = new(service.UserServer)

	var req user.LoginReq
	req.UserName = "root"
	req.PassWord = utils.Md5("123456")

	var rsp user.LoginRsp
	err := userServer.Login(context.Background(), &req, &rsp)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Code)
	fmt.Println(rsp.Token)
}

func TestRegister(t *testing.T) {
	var userServer = new(service.UserServer)
	var req user.RegisterReq
	req.User = &user.User{
		Username: "root",
		Password: utils.Md5("123456"),
		Phone:    "177",
		Nickname: "鲍勃",
	}

	var rsp user.RegisterRsp

	_ = userServer.Register(context.Background(), &req, &rsp)
	fmt.Println(rsp.Code)
}
