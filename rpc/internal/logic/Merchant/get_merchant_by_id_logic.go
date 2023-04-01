package Merchant

import (
	"context"

	"github.com/huuhoait/zero-admin-core/rpc/internal/svc"
	"github.com/huuhoait/zero-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMerchantByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMerchantByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMerchantByIdLogic {
	return &GetMerchantByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMerchantByIdLogic) GetMerchantById(in *core.IDReq) (*core.MerchantInfo, error) {
	// todo: add your logic here and delete this line

	return &core.MerchantInfo{}, nil
}
