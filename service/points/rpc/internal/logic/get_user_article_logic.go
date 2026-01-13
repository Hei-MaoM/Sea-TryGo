package logic

import (
	"context"

	"sea-try-go/service/points/rpc/internal/svc"
	"sea-try-go/service/points/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserArticleLogic {
	return &GetUserArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserArticleLogic) GetUserArticle(in *__.GetUserArticleReq) (*__.GetUserArticleResp, error) {
	// todo: add your logic here and delete this line

	return &__.GetUserArticleResp{}, nil
}
