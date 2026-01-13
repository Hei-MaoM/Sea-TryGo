package logic

import (
	"context"

	"sea-try-go/service/points/rpc/internal/svc"
	"sea-try-go/service/points/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleExposureLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddArticleExposureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleExposureLogic {
	return &AddArticleExposureLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddArticleExposureLogic) AddArticleExposure(in *__.AddArticleExposureReq) (*__.AddArticleExposureResp, error) {
	// todo: add your logic here and delete this line

	return &__.AddArticleExposureResp{}, nil
}
