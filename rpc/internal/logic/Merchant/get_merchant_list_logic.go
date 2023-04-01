package Merchant

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMerchantListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantListLogic {
	return &GetMerchantListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMerchantListLogic) GetMerchantList(in *core.MerchantListReq) (*core.MerchantListResp, error) {
	// todo: add your logic here and delete this line

	return &core.MerchantListResp{}, nil
}
