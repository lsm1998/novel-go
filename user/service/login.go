package service

import (
	"context"
	"crypto/md5"
	"proto/user"
	"time"
	. "user/config"
	"user/model"
	"utils"
)

const (
	BaseUid = 1000
)

// Login 登录方法
func (t *UserServer) Login(ctx context.Context, req *user.LoginReq, rsp *user.LoginRsp) error {
	var userInfo model.UserInfo
	err := DB.Table((*model.UserInfo)(nil).TableName()).Select("uid,nickname,password,salt,status").Where("username=?", req.UserName).Find(&userInfo).Error
	if err != nil {
		rsp.Code = utils.ERR_USER_NOT_EXIS
		return nil
	}
	if userInfo.Status != utils.ENABLE {
		rsp.Code = utils.ERR_USER_DISABLE
		return nil
	}
	// 登录校验规则，前端MD5加密后请求，尾部拼接密码盐
	has := md5.Sum([]byte(userInfo.Password))
	bytes := has[:]
	bytes = append(bytes, []byte(userInfo.Salt)...)

	if req.PassWord == string(bytes) {
		rsp.Code = utils.OK
		token, _ := utils.GenerateToken(int64(userInfo.UID), []int64{1, 2, 3})
		rsp.Token = token
		_ = DB.Table((*model.UserDetails)(nil).TableName()).Select("head_img,phone,birthday,region,nation,autograph").Where("uid=?", userInfo.UID).Find(&rsp.User).Error
		rsp.User.Uid = int32(userInfo.UID)
		rsp.User.Nickname = userInfo.Nickname
		rsp.User.Username = userInfo.Username
	} else {
		rsp.Code = utils.ERR_LOGIN_FAIL
	}
	return nil
}

func (t *UserServer) Register(ctx context.Context, req *user.RegisterReq, rsp *user.RegisterRsp) error {
	key := "Register_lock"
	if err := utils.RedisLock(RedisClient.Get(), key, 10, 10); err != nil {
		rsp.Code = utils.ERR_LIMIT
		return nil
	}
	defer utils.UnRedisLock(RedisClient.Get(), key)
	tx := DB.Begin()
	var count int
	now := time.Now().Unix()
	_ = tx.Table((*model.UserInfo)(nil).TableName()).Count(&count).Error
	userInfo := &model.UserInfo{Username: req.User.Username, UID: BaseUid + count,
		Password: req.User.Password, Status: utils.ENABLE, CreateTime: int(now), UpdateTime: int(now)}
	userDetails := &model.UserDetails{UID: BaseUid + count, Phone: req.User.Phone, CreateTime: int(now), UpdateTime: int(now),
		HeadImg: req.User.HeadImg, Region: req.User.Region, Nation: req.User.Nation, Status: utils.ENABLE}
	if err := tx.Table((*model.UserInfo)(nil).TableName()).Save(userInfo).Error; err != nil {
		rsp.Code = utils.ERR_DB
		tx.Rollback()
		return nil
	}
	if err := tx.Table((*model.UserDetails)(nil).TableName()).Save(userDetails).Error; err != nil {
		rsp.Code = utils.ERR_DB
		tx.Rollback()
		return nil
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		rsp.Code = utils.ERR_DB
		return nil
	}
	return nil
}
