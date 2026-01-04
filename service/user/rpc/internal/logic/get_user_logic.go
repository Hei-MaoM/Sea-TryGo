package logic

import (
	"context"

	"sea-try-go/service/user/rpc/internal/svc"
	"sea-try-go/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *__.GetUserReq) (*__.GetUserResp, error) {
	// todo: add your logic here and delete this line

	return &__.GetUserResp{}, nil
}
