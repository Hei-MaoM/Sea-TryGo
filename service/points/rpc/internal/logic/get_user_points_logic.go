package logic

import (
	"context"

	"sea-try-go/service/points/rpc/internal/svc"
	"sea-try-go/service/points/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPointsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPointsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPointsLogic {
	return &GetUserPointsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserPointsLogic) GetUserPoints(in *__.GetUserPointsReq) (*__.GetUserPointsResp, error) {
	// todo: add your logic here and delete this line

	return &__.GetUserPointsResp{}, nil
}
