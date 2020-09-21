package service

import "context"
import "proto/user"

func (t *UserServer) Login(ctx context.Context, req *user.LoginReq, rsp *user.LoginRsp) error {
	if req.UserName == "root" && req.PassWord == "123" {
		rsp.Code = 0
	} else {
		rsp.Code = 201
	}
	return nil
}
