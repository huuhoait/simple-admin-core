package Merchant

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMerchantLogic {
	return &CreateMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Merchant management
func (l *CreateMerchantLogic) CreateMerchant(in *core.MerchantInfo) (*core.BaseIDResp, error) {
	// todo: add your logic here and delete this line

	return &core.BaseIDResp{}, nil
}
