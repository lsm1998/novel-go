package service

import (
	"context"
	"proto/base"
	"utils"

	. "base/config"
)

func (b *BaseServer) BannerList(ctx context.Context, req *base.BannerListReq, rsp *base.BannerListRsp) error {
	rows, err := MysqlDB.Query("SELECT title,link,img FROM t_banner WHERE enable=? AND status=? LIMIT 4", utils.ENABLE, utils.ENABLE)
	if err != nil {
		rsp.Code = utils.ERR_DB
		return err
	}
	defer rows.Close()
	rsp.List = make([]*base.Banner, 0, 4)
	for rows.Next() {
		temp := base.Banner{}
		err = rows.Scan(&temp.Title, &temp.Link, &temp.Img)
		if err != nil {
			rsp.Code = utils.ERR_DB
			return err
		}
		rsp.List = append(rsp.List, &temp)
	}
	rsp.Code = 200
	return nil
}
