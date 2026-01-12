// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package admin

import (
	"context"
	"errors"

	"sea-try-go/service/admin/api/internal/model"
	"sea-try-go/service/admin/api/internal/svc"
	"sea-try-go/service/admin/api/internal/types"
	"sea-try-go/service/common/cryptx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetuserpasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetuserpasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetuserpasswordLogic {
	return &ResetuserpasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetuserpasswordLogic) Resetuserpassword(req *types.ResetUserPasswordReq) (resp *types.ResetUserPasswordResp, err error) {
	id := req.Id
	user := model.User{}
	err = l.svcCtx.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	var password string
	password, err = cryptx.PasswordEncrypt(l.svcCtx.Config.System.DefaultPassword)
	if err != nil {
		return nil, err
	}
	user.Password = password
	err = l.svcCtx.DB.Model(&user).Update("password", password).Error
	return &types.ResetUserPasswordResp{
		Success: true,
	}, nil
}
