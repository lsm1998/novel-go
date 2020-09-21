package service

import (
	"context"
	"proto/im"
)

func (im *ImServer) SendMessage(ctx context.Context, req *im.Message, rsp *im.Message) error {
	rsp.Id = req.Id * 2
	return nil
}
