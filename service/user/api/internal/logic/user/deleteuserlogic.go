// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"
	"encoding/json"

	"sea-try-go/service/user/api/internal/model"
	"sea-try-go/service/user/api/internal/svc"
	"sea-try-go/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteuserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteuserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteuserLogic {
	return &DeleteuserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteuserLogic) Deleteuser(req *types.DeleteUserReq) (resp *types.DeleteUserResp, err error) {

	//这里也是写的普通user对应的

	//"userId"和token的claims中的字段对应
	userId := l.ctx.Value("userId").(json.Number)
	id, _ := userId.Int64()

	err = l.svcCtx.DB.Where("id = ?", uint64(id)).Delete(&model.User{}).Error
	if err != nil {
		return &types.DeleteUserResp{
			Success: false,
		}, err
	}
	return &types.DeleteUserResp{
		Success: true,
	}, nil
}
