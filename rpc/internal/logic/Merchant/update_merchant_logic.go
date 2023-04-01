package Merchant

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMerchantLogic {
	return &UpdateMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMerchantLogic) UpdateMerchant(in *core.MerchantInfo) (*core.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &core.BaseResp{}, nil
}
