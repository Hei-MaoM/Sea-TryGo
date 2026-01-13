package logic

import (
	"context"

	"sea-try-go/service/points/rpc/internal/svc"
	"sea-try-go/service/points/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginPointsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginPointsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginPointsLogic {
	return &LoginPointsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginPointsLogic) LoginPoints(in *__.LoginPointsReq) (*__.LoginPointsResp, error) {
	// todo: add your logic here and delete this line

	return &__.LoginPointsResp{}, nil
}
