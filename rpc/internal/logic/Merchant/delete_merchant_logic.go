package Merchant

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMerchantLogic {
	return &DeleteMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMerchantLogic) DeleteMerchant(in *core.IDsReq) (*core.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &core.BaseResp{}, nil
}
