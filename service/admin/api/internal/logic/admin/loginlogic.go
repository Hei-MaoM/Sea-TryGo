// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package admin

import (
	"context"
	"errors"
	"time"

	"sea-try-go/service/admin/api/internal/model"
	"sea-try-go/service/admin/api/internal/svc"
	"sea-try-go/service/admin/api/internal/types"
	"sea-try-go/service/common/cryptx"
	"sea-try-go/service/common/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	username := req.Username
	password := req.Password
	admin := model.Admin{}
	err = l.svcCtx.DB.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}
	correct := cryptx.CheckPassword(admin.Password, password)
	if !correct {
		return nil, errors.New("用户名或密码错误")
	}
	now := time.Now().Unix()
	accessSecret := l.svcCtx.Config.AdminAuth.AccessSecret
	accessExpire := l.svcCtx.Config.AdminAuth.AccessExpire
	token, er := jwt.GetToken(accessSecret, now, accessExpire, int64(admin.Id))
	if er != nil {
		return nil, er
	}
	return &types.LoginResp{
		Token: token,
	}, nil
}
